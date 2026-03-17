package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("12345678901234567890123456789012")
var db *sql.DB

type User struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
}

type Credentials struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	TokenType string `json:"token_type"`
	jwt.StandardClaims
	HasuraClaims map[string]interface{} `json:"https://hasura.io/jwt/claims,omitempty"`
}

type HasuraActionRequest struct {
	Input Credentials `json:"input"`
}

func main() {
	var err error
	db, err = sql.Open("postgres", "postgres://myuser:mysecretpassword@postgres:5432/foodrecipes?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to Postgres at postgres:5432 (Docker network)")
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Go Auth API 🚀"))
	}).Methods("GET")
	r.HandleFunc("/user/register", SignupHandler).Methods("POST")
	r.HandleFunc("/user/login", LoginHandler).Methods("POST")
	r.HandleFunc("/user/refresh", RefreshHandler).Methods("GET")
	r.HandleFunc("/user/logout", LogoutHandler).Methods("POST")

	log.Println("Auth service running on :3001")
	log.Fatal(http.ListenAndServe(":3001", corsMiddleware(r)))
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("SignupHandler called")
	var req HasuraActionRequest
	body, _ := ioutil.ReadAll(r.Body)
	log.Println("Raw body:", string(body))
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // Reset body for json decoder
	json.NewDecoder(r.Body).Decode(&req)
	creds := req.Input
	log.Printf("Registering: name=%s, email=%s, password=%s", creds.Name, creds.Email, creds.Password)

	if creds.Email == "" || creds.Name == "" || creds.Password == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	id := uuid.New()
	_, err = db.Exec("INSERT INTO users (id, name, email, password_hash) VALUES ($1, $2, $3, $4)", id, creds.Name, creds.Email, string(hash))
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			http.Error(w, "Email already exists", http.StatusBadRequest)
		} else {
			http.Error(w, "User already exists or DB error", http.StatusBadRequest)
		}
		return
	}

	accessToken, refreshToken, err := generateTokens(id.String(), creds.Name, creds.Email)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Set refresh token as HTTP-only cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		// Secure: true, // Uncomment if using HTTPS
		Expires: time.Now().Add(7 * 24 * time.Hour),
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"accessToken": accessToken,
		"name":        creds.Name,
		"email":       creds.Email,
		"id":          id.String(),
	})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("LoginHandler called")
	var req HasuraActionRequest
	body, _ := ioutil.ReadAll(r.Body)
	log.Println("Raw body (login):", string(body))
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // Reset body for json decoder
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("JSON decode error (login):", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	creds := req.Input
	log.Printf("Logging in: email=%s, password=%s", creds.Email, creds.Password)

	var user User
	err = db.QueryRow("SELECT id, name, email, password_hash FROM users WHERE email=$1", creds.Email).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(creds.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	accessToken, refreshToken, err := generateTokens(user.ID.String(), user.Name, user.Email)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Set refresh token as HTTP-only cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		// Secure: true, // Uncomment if using HTTPS
		Expires: time.Now().Add(7 * 24 * time.Hour),
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"accessToken": accessToken,
		"name":        user.Name,
		"email":       user.Email,
		"id":          user.ID.String(),
	})
}

func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, "No refresh token", http.StatusUnauthorized)
		return
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || claims.TokenType != "refresh" {
		http.Error(w, "Invalid token type", http.StatusUnauthorized)
		return
	}

	// Generate new access token
	accessToken, _, err := generateTokens(claims.UserID, claims.Name, claims.Email)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"accessToken": accessToken,
		"name":        claims.Name,
		"email":       claims.Email,
		"id":          claims.UserID,
	})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
	})
	w.WriteHeader(http.StatusOK)
}

func generateTokens(userID, name, email string) (accessToken, refreshToken string, err error) {
	// Access Token (short-lived)
	accessClaims := Claims{
		UserID:    userID,
		Name:      name,
		Email:     email,
		TokenType: "access",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
		HasuraClaims: map[string]interface{}{
			"x-hasura-default-role":  "user",
			"x-hasura-allowed-roles": []string{"user", "anonymous"},
			"x-hasura-user-id":       userID,
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err = at.SignedString(jwtSecret)
	if err != nil {
		return
	}

	// Refresh Token (long-lived)
	refreshClaims := Claims{
		UserID:    userID,
		Name:      name,
		Email:     email,
		TokenType: "refresh",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
		},
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err = rt.SignedString(jwtSecret)
	return
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
