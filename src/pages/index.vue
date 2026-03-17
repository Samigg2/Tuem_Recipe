<template>
  <div class="flex flex-col min-h-screen bg-gradient-to-b from-white to-slate-50">
    <!-- Hero Section -->
    <div class="relative">
      <!-- Hero Section -->
      <div class="relative flex flex-col items-center justify-center bg-center h-[32rem] bg-[url('/src/assets/hero-bg.jpg')] bg-cover bg-no-repeat before:content-[''] before:absolute before:inset-0 before:bg-gradient-to-b before:from-black/80 before:to-black/60">
        <div class="relative z-10 max-w-5xl mx-auto text-center px-4">
          <h1 class="font-bold text-white mb-8 font-['Dancing_Script'] flex flex-col items-center gap-8">
            <span class="animate-fade-in-down text-white/90 text-4xl md:text-6xl lg:text-7xl">Welcome to</span>
            <div class="relative">
              <span class="animate-typewriter-infinite text-5xl md:text-7xl lg:text-8xl text-emerald-400 tracking-wide whitespace-nowrap">Tuem Recipe</span>
              <span class="absolute right-0 top-0 h-full w-1 bg-emerald-400 animate-cursor-infinite"></span>
            </div>
          </h1>
          <p class="text-2xl md:text-3xl lg:text-4xl text-white/90 mb-8 max-w-3xl mx-auto font-['Dancing_Script'] leading-relaxed animate-fade-in">
            Discover and share authentic recipes from around the world. Join our community of food lovers!
          </p>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <main class="flex-grow">
      <!-- Sticky Navigation Bar -->
      <div class="sticky top-16 z-30 bg-white/95 backdrop-blur-sm shadow-md transition-all duration-300">
        <div class="container mx-auto px-4 py-3 flex flex-col sm:flex-row items-center gap-4">
          <!-- Categories -->
          <div class="flex items-center gap-2 flex-1">
            <button
              @click="() => { scrollBar.scrollLeft -= 200; showAllCategories = !showAllCategories }"
              class="p-2 rounded-full hover:bg-emerald-50 transition-all duration-300 text-emerald-600"
            >
              <BackArrow class="w-5 h-5" />
            </button>

            <div ref="scrollBar" class="flex gap-2 overflow-x-auto scrollbar-hide">
              <button
                v-for="item in currentCategories"
                @click="change_catagory(item.id)"
                :key="item.id"
                :class="{
                  'bg-emerald-600 text-white hover:bg-emerald-700': selectedCategoryId === item.id,
                  'bg-transparent text-emerald-800 hover:bg-emerald-50': selectedCategoryId !== item.id,
                }"
                class="px-4 py-2 rounded-full border-2 border-emerald-600 transition-all duration-300 font-['Quicksand'] whitespace-nowrap text-sm font-medium hover:scale-105"
              >
                {{ item.title }}
              </button>
            </div>

            <button
              @click="() => { scrollBar.scrollLeft += 200; showAllCategories = !showAllCategories }"
              class="p-2 rounded-full hover:bg-emerald-50 transition-all duration-300 text-emerald-600"
            >
              <ForwardArrow class="w-5 h-5" />
            </button>
          </div>

          <!-- Search Bar with Filter -->
          <div class="flex gap-2 flex-1 justify-end max-w-md">
            <div class="relative flex-1">
              <input
                @input="handleSearch"
                type="search"
                v-model="search"
                placeholder="Search recipes..."
                class="w-full px-4 py-2 text-sm rounded-lg border border-emerald-200 focus:border-emerald-500 outline-none transition-colors font-['Quicksand'] placeholder:text-gray-400"
              />
            </div>
            <div class="relative">
              <select
                v-model="searchby"
                @change="handleFilterChange"
                class="appearance-none bg-emerald-600 text-white px-8 py-2 rounded-lg text-sm hover:bg-emerald-700 transition-colors cursor-pointer outline-none font-['Quicksand'] pr-10"
              >
                <option disabled value="">Filter</option>
                <option v-for="item in option" :value="item" :key="item">{{ item }}</option>
              </select>
              <div class="absolute right-3 top-1/2 transform -translate-y-1/2 pointer-events-none">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path d="M19 9l-7 7-7-7" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2"/>
                </svg>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="flex w-full h-[200px] justify-center items-center">
        <img class="w-20 h-20" src="../assets/loading.svg" alt="Loading..." />
      </div>

      <!-- No Results State -->
      <div
        v-if="result && result.search_recipe.length < 1"
        class="flex relative w-full h-[200px] justify-center items-center"
      >
        <span class="text-2xl text-red-600 font-medium font-['Quicksand']">No recipes found</span>
      </div>

      <!-- Results Grid -->
      <div
        v-if="result"
        class="container mx-auto grid gap-6 px-4 mt-8 xl:grid-cols-3 lg:grid-cols-3 md:grid-cols-2 sm:grid-cols-2 grid-cols-1"
      >
        <div
          v-for="item in result.search_recipe"
          :key="item.id"
          class="group w-full max-w-sm bg-gradient-to-b from-white to-gray-50 rounded-xl shadow-lg overflow-hidden hover:shadow-2xl transition-all duration-300 transform hover:-translate-y-1 cursor-pointer border border-gray-100"
          @click="showDetail(item.id)"
        >
          <div class="relative overflow-hidden">
            <img 
              :src="item.thum_image" 
              :alt="item.name"
              class="w-full h-48 object-cover transform group-hover:scale-110 transition-transform duration-500"
            />
            <div class="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
          </div>
          
          <div class="p-4">
            <h3 class="text-2xl md:text-3xl font-bold text-emerald-900 mb-3 font-['Dancing_Script'] line-clamp-1 group-hover:text-emerald-600 transition-colors">{{ item.name }}</h3>
            <p class="text-gray-700 text-base mb-4 font-['Quicksand'] font-medium line-clamp-2">{{ item.description }}</p>
            
            <div class="flex items-center justify-between text-sm text-gray-600">
              <div class="flex items-center gap-3">
                <!-- Star Rating -->
                <div class="flex items-center">
                  <div class="flex">
                    <template v-for="index in 5" :key="index">
                      <svg 
                        class="w-5 h-5" 
                        :class="index <= item.rate ? 'text-yellow-400' : 'text-gray-300'"
                        xmlns="http://www.w3.org/2000/svg" 
                        viewBox="0 0 20 20" 
                        fill="currentColor"
                      >
                        <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                  </svg>
                    </template>
                  </div>
                  <span class="text-gray-500 ml-1">{{ item.numberofrate }}</span>
                </div>

                <!-- Comments -->
                <div class="flex gap-1">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-emerald-600" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd" d="M18 13V5a2 2 0 00-2-2H4a2 2 0 00-2 2v8a2 2 0 002 2h3l3 3 3-3h3a2 2 0 002-2zM5 7a1 1 0 011-1h8a1 1 0 110 2H6a1 1 0 01-1-1zm1 3a1 1 0 100 2h3a1 1 0 100-2H6z" clip-rule="evenodd" />
                  </svg>
                  <span class="text-gray-500">{{ item.numberofcomment }}</span>
                </div>

                <!-- Likes -->
                <div class="flex gap-1">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-emerald-600" viewBox="0 0 20 20" fill="currentColor">
                    <path d="M2 10.5a1.5 1.5 0 113 0v6a1.5 1.5 0 01-3 0v-6zM6 10.333v5.43a2 2 0 001.106 1.79l.05.025A4 4 0 008.943 18h5.416a2 2 0 001.962-1.608l1.2-6A2 2 0 0015.56 8H12V4a2 2 0 00-2-2 1 1 0 00-1 1v.667a4 4 0 01-.8 2.4L6.8 7.933a4 4 0 00-.8 2.4z" />
                  </svg>
                  <span class="text-gray-500">{{ item.numberoflike }}</span>
                </div>
              </div>

              <!-- Duration -->
              <div class="flex items-center gap-1">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-emerald-600" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z" clip-rule="evenodd" />
                </svg>
                <span class="text-gray-500">{{ item.duration }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- Footer -->
    <footer class="bg-gradient-to-br from-emerald-900 to-teal-800 text-white py-12 mt-auto">
      <div class="container mx-auto px-4">
        <div class="grid grid-cols-1 md:grid-cols-4 gap-8">
          <!-- About Section -->
          <div class="space-y-4">
            <h3 class="text-2xl font-['Dancing_Script'] font-bold">Tuem Recipe</h3>
            <p class="text-white/80 font-['Quicksand']">Your ultimate destination for discovering and sharing authentic recipes from around the world.</p>
          </div>

          <!-- Quick Links -->
          <div class="space-y-4">
            <h4 class="text-lg font-semibold font-['Quicksand']">Quick Links</h4>
            <ul class="space-y-2 text-white/80 font-['Quicksand']">
              <li><a href="/" class="hover:text-emerald-300 transition-colors">Home</a></li>
              <li><a href="/bookmarks" class="hover:text-emerald-300 transition-colors">Bookmarks</a></li>
              <li><a href="/recipe" class="hover:text-emerald-300 transition-colors">Create Recipe</a></li>
            </ul>
          </div>

          <!-- Categories -->
          <div class="space-y-4">
            <h4 class="text-lg font-semibold font-['Quicksand']">Categories</h4>
            <ul class="space-y-2 text-white/80 font-['Quicksand']">
              <li><a href="#" class="hover:text-emerald-300 transition-colors">Breakfast</a></li>
              <li><a href="#" class="hover:text-emerald-300 transition-colors">Lunch</a></li>
              <li><a href="#" class="hover:text-emerald-300 transition-colors">Dinner</a></li>
              <li><a href="#" class="hover:text-emerald-300 transition-colors">Desserts</a></li>
            </ul>
          </div>

          <!-- Contact -->
          <div class="space-y-4">
            <h4 class="text-lg font-semibold font-['Quicksand']">Connect With Us</h4>
            <div class="flex space-x-4">
              <a href="#" class="text-white/80 hover:text-emerald-300 transition-colors">
                <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24"><path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/></svg>
              </a>
              <a href="#" class="text-white/80 hover:text-emerald-300 transition-colors">
                <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24"><path d="M23.953 4.57a10 10 0 01-2.825.775 4.958 4.958 0 002.163-2.723c-.951.555-2.005.959-3.127 1.184a4.92 4.92 0 00-8.384 4.482C7.69 8.095 4.067 6.13 1.64 3.162a4.822 4.822 0 00-.666 2.475c0 1.71.87 3.213 2.188 4.096a4.904 4.904 0 01-2.228-.616v.06a4.923 4.923 0 003.946 4.827 4.996 4.996 0 01-2.212.085 4.936 4.936 0 004.604 3.417 9.867 9.867 0 01-6.102 2.105c-.39 0-.779-.023-1.17-.067a13.995 13.995 0 007.557 2.209c9.053 0 13.998-7.496 13.998-13.985 0-.21 0-.42-.015-.63A9.935 9.935 0 0024 4.59z"/></svg>
              </a>
              <a href="#" class="text-white/80 hover:text-emerald-300 transition-colors">
                <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24"><path d="M12 0C8.74 0 8.333.015 7.053.072 5.775.132 4.905.333 4.14.63c-.789.306-1.459.717-2.126 1.384S.935 3.35.63 4.14C.333 4.905.131 5.775.072 7.053.012 8.333 0 8.74 0 12s.015 3.667.072 4.947c.06 1.277.261 2.148.558 2.913.306.788.717 1.459 1.384 2.126.667.666 1.336 1.079 2.126 1.384.766.296 1.636.499 2.913.558C8.333 23.988 8.74 24 12 24s3.667-.015 4.947-.072c1.277-.06 2.148-.262 2.913-.558.788-.306 1.459-.718 2.126-1.384.666-.667 1.079-1.335 1.384-2.126.296-.765.499-1.636.558-2.913.06-1.28.072-1.687.072-4.947s-.015-3.667-.072-4.947c-.06-1.277-.262-2.149-.558-2.913-.306-.789-.718-1.459-1.384-2.126C21.319 1.347 20.651.935 19.86.63c-.765-.297-1.636-.499-2.913-.558C15.667.012 15.26 0 12 0zm0 2.16c3.203 0 3.585.016 4.85.071 1.17.055 1.805.249 2.227.415.562.217.96.477 1.382.896.419.42.679.819.896 1.381.164.422.36 1.057.413 2.227.057 1.266.07 1.646.07 4.85s-.015 3.585-.074 4.85c-.061 1.17-.256 1.805-.421 2.227-.224.562-.479.96-.899 1.382-.419.419-.824.679-1.38.896-.42.164-1.065.36-2.235.413-1.274.057-1.649.07-4.859.07-3.211 0-3.586-.015-4.859-.074-1.171-.061-1.816-.256-2.236-.421-.569-.224-.96-.479-1.379-.899-.421-.419-.69-.824-.9-1.38-.165-.42-.359-1.065-.42-2.235-.045-1.26-.061-1.649-.061-4.844 0-3.196.016-3.586.061-4.861.061-1.17.255-1.814.42-2.234.21-.57.479-.96.9-1.381.419-.419.81-.689 1.379-.898.42-.166 1.051-.361 2.221-.421 1.275-.045 1.65-.06 4.859-.06l.045.03zm0 3.678c-3.405 0-6.162 2.76-6.162 6.162 0 3.405 2.76 6.162 6.162 6.162 3.405 0 6.162-2.76 6.162-6.162 0-3.405-2.76-6.162-6.162-6.162zM12 16c-2.21 0-4-1.79-4-4s1.79-4 4-4 4 1.79 4 4-1.79 4-4 4zm7.846-10.405c0 .795-.646 1.44-1.44 1.44-.795 0-1.44-.646-1.44-1.44 0-.794.646-1.439 1.44-1.439.793-.001 1.44.645 1.44 1.439z"/></svg>
              </a>
            </div>
            <p class="text-white/80 font-['Quicksand']">Email: contact@deliciousrecipes.com</p>
          </div>
        </div>

        <div class="border-t border-white/10 mt-8 pt-8 text-center text-white/60 font-['Quicksand']">
          <p>&copy; {{ new Date().getFullYear() }} Tuem Recipe. All rights reserved.</p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import onloading from "../components/onLoading.vue";
import { ref, computed } from "@vue/reactivity";
import { onMounted, onUpdated, watchEffect } from "@vue/runtime-core";
import Card from "../components/cards.vue";
import app from "../Axios/axiosconf";
import { useStore } from "vuex";
const store = useStore();
import gql from "graphql-tag";

import { useQuery, useQueryLoading } from "@vue/apollo-composable";
import { useRouter } from 'vue-router';
const router = useRouter();
const catagory_id = ref(0);
const scrollBar = ref(null);
const query_recipe = gql`
  query($title: String, $search: String, $cataid: Int) {
    search_recipe(args: { search: $search, title: $title, cata: $cataid }) {
      name
      id
      description
      createdBy
      numberoflike
      numberofrate
      numberofcomment
      thum_image
      rate
      duration
    }
  }
`;

const variables = ref({ cataid: 0 });
watchEffect(() => {});
const option = ref(["duration", "title", "Ingradient"]);

const { result, loading, error, fetchMore: fetch_More } = useQuery(
  query_recipe,
  variables
);
onMounted(() => {
  fetch_More({ variables: { cataid: 0 } });
});

const showAllCategories = ref(false);

// All categories
const allCategories = ref([
  { title: "All", id: 0 },
  { title: "Breakfast", id: 250 },
  { title: "Lunch", id: 249 },
  { title: "Dinner", id: 251 },
  { title: "Dessert", id: 252 },
  { title: "Appetizer", id: 253 },
  { title: "Side Dish", id: 254 },
  { title: "Quick & Easy", id: 255 },
  { title: "Holiday", id: 256 },
  { title: "Soup", id: 257 },
  { title: "Vegetarian", id: 258 },
  { title: "Salad", id: 259 },
]);

// Main categories (first 4)
const mainCategories = ref([
  { title: "All", id: 0 },
  { title: "Breakfast", id: 250 },
  { title: "Lunch", id: 249 },
  { title: "Dinner", id: 251 },
]);

// Computed property to show either main or all categories
const currentCategories = computed(() => {
  return showAllCategories.value ? allCategories.value : mainCategories.value;
});

// Debounce function
function debounce(func, wait) {
  let timeout;
  return function executedFunction(...args) {
    const later = () => {
      clearTimeout(timeout);
      func(...args);
    };
    clearTimeout(timeout);
    timeout = setTimeout(later, wait);
  };
}

// Debounced search function
const debouncedSearch = debounce((value) => {
  search_recipe(searchby.value, value);
}, 300);

// Update handleSearch to use debounced function
const handleSearch = (e) => {
  debouncedSearch(e.target.value);
};

function search_recipe(title, search) {
  if (!title) {
    title = "title"; // Default search by title if no filter selected
  }
  variables.value = { 
    title, 
    search,
    cataid: catagory_id.value 
  };
}

function change_catagory(id) {
  selectedCategoryId.value = id;
  catagory_id.value = id;
  variables.value = { cataid: id };
}

const searchby = ref("");
const search = ref("");
const selectedCategoryId = ref(0);

// Enhanced filter functionality
const handleFilterChange = () => {
  if (search.value) {
    search_recipe(searchby.value, search.value);
  }
};

const showDetail = (id) => {
  router.push({ 
    name: "detail", 
    query: { productID: id } 
  });
};
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Dancing_Script:wght@400;500;600;700&display=swap');

/* Hide scrollbar but keep functionality */
.scrollbar-hide {
  -ms-overflow-style: none;
  scrollbar-width: none;
}
.scrollbar-hide::-webkit-scrollbar {
  display: none;
}

/* Animations */
@keyframes typewriter-infinite {
  0%, 100% { clip-path: inset(0 100% 0 0); }
  30%, 70% { clip-path: inset(0 0 0 0); }
}

@keyframes cursor-infinite {
  0%, 100% { opacity: 0; }
  30%, 70% { opacity: 1; }
}

@keyframes fade-in-down {
  0% {
    opacity: 0;
    transform: translateY(-20px);
  }
  100% {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-typewriter-infinite {
  display: inline-block;
  animation: typewriter-infinite 5s linear infinite;
}

.animate-cursor-infinite {
  animation: cursor-infinite 5s linear infinite;
}

.animate-fade-in-down {
  animation: fade-in-down 0.8s ease-out forwards;
}

.animate-fade-in {
  animation: fade-in 1s ease-out forwards;
}

/* Layout styles */
#app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

main {
  flex: 1;
}

/* Rest of the existing styles */
@keyframes continuous-slide-down {
  0%, 100% {
    opacity: 0;
    transform: translateY(-30px);
  }
  50% {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes continuous-slide-up {
  0%, 100% {
    opacity: 0;
    transform: translateY(30px);
  }
  50% {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes continuous-fade {
  0%, 100% {
    opacity: 0.5;
  }
  50% {
    opacity: 1;
  }
}

.animate-continuous-slide-down {
  animation: continuous-slide-down 3s ease-in-out infinite;
}

.animate-continuous-slide-up {
  animation: continuous-slide-up 3s ease-in-out infinite;
}

.animate-continuous-fade {
  animation: continuous-fade 3s ease-in-out infinite;
}

/* Add new styles for search improvements */
input[type="search"]::-webkit-search-decoration,
input[type="search"]::-webkit-search-cancel-button,
input[type="search"]::-webkit-search-results-button,
input[type="search"]::-webkit-search-results-decoration {
  display: none;
}

input[type="search"]:focus {
  box-shadow: 0 0 0 2px rgba(16, 185, 129, 0.2);
}

select {
  appearance: none;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' fill='white' viewBox='0 0 24 24'%3E%3Cpath d='M7 10l5 5 5-5z'/%3E%3C/svg%3E");
  background-repeat: no-repeat;
  background-position: right 0.5rem center;
  background-size: 1.5em;
  padding-right: 2.5rem;
}

/* Improve search container responsiveness */
.search-container {
  transition: all 0.3s ease;
}

@media (max-width: 640px) {
  .search-container {
    width: 100%;
  }
}

/* Add new styles for cards */
.line-clamp-1 {
  display: -webkit-box;
  -webkit-line-clamp: 1;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

/* Improve filter dropdown */
select:focus {
  box-shadow: 0 0 0 2px rgba(16, 185, 129, 0.2);
}

/* Add smooth hover transitions for cards */
.card-hover {
  transition: all 0.3s ease;
}

.card-hover:hover {
  transform: translateY(-4px);
}

/* Add animation for image zoom */
.image-zoom {
  transition: transform 0.5s ease;
}

.image-zoom:hover {
  transform: scale(1.1);
}

/* Add new hover effect for the view details section */
.group:hover .transform {
  transition-duration: 300ms;
}

/* Improve card interaction */
.group {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.group:hover {
  transform: translateY(-4px) scale(1.01);
}

/* Enhance image zoom effect */
.group:hover img {
  transform: scale(1.1);
}

/* Smooth transition for view details text */
.group:hover .opacity-0 {
  transition-delay: 100ms;
}

/* Add new styles for improved card design */
.group {
  backdrop-filter: blur(8px);
}

.group:hover {
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);
  transform: translateY(-4px);
}

/* Improve text contrast */
.text-emerald-800 {
  color: #065f46;
}

.text-emerald-600 {
  color: #059669;
}

/* Add smooth transitions */
.transition-transform {
  transition-duration: 500ms;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
}

/* Page layout */
#app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

main {
  flex: 1;
}

/* Remove extra space after footer */
footer {
  margin-top: 0;
}

/* Add styles for view detail button animation */
.group:hover .invisible {
  visibility: visible;
}

.group:hover .group-hover\:left-\[20\%\] {
  left: 20%;
}

.group:hover .group-hover\:right-\[20\%\] {
  right: 20%;
}

/* Center the grid items */
.grid {
  justify-items: center;
}

/* Ensure consistent card width */
.max-w-sm {
  width: 100%;
  max-width: 24rem;
}

/* Add these styles to improve mobile menu interaction */
.backdrop-blur-sm {
  backdrop-filter: blur(8px);
  -webkit-backdrop-filter: blur(8px);
}

/* Improve touch targets on mobile */
.cursor-pointer {
  cursor: pointer;
  -webkit-tap-highlight-color: transparent;
}

/* Add active state for mobile touches */
.cursor-pointer:active {
  opacity: 0.7;
}

/* Mobile Menu Styles */
.mobile-menu {
  position: fixed;
  top: 0;
  right: 0;
  width: 50%;
  height: 100vh;
  background: linear-gradient(to bottom, rgb(6 78 59), rgb(20 184 166));
  backdrop-filter: blur(8px);
  transform: translateX(100%);
  transition: transform 0.3s ease-in-out;
  z-index: 50;
}

.mobile-menu.open {
  transform: translateX(0);
}

.mobile-menu-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100vh;
  background: rgba(0, 0, 0, 0.5);
  backdrop-filter: blur(4px);
  opacity: 0;
  visibility: hidden;
  transition: all 0.3s ease-in-out;
  z-index: 40;
}

.mobile-menu-overlay.open {
  opacity: 1;
  visibility: visible;
}

.mobile-menu-content {
  padding: 2rem;
  height: 100%;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.mobile-menu-item {
  color: white;
  padding: 0.75rem 1rem;
  border-radius: 0.5rem;
  transition: all 0.2s ease-in-out;
}

.mobile-menu-item:hover {
  background: rgba(255, 255, 255, 0.1);
  transform: translateX(4px);
}
</style>

<route lang="yaml">
meta:
  layout: header
</route>

