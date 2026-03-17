import axios from "axios";
const baseURL = "http://localhost:3001/";

const app = axios.create({
  baseURL,
  withCredentials: true,
});



app.interceptors.response.use(
    (response) => {
      return response;
    },
    (error) => {
        console.log(error.message);
      if (!error.response) {
        console.log("Please check your internet connection.");
      }
  
      return Promise.reject(error);
    }
  );
  
  export default app;
  
