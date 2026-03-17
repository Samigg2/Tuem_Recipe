<template>
  <div class="w-full">
    <header
      class="flex drop-shadow-lg bg-gradient-to-r from-emerald-800 to-teal-800 fixed top-0 w-full h-16 justify-between items-center px-[5%] z-50"
    >
      <div class="flex items-center gap-3">
        <img src="/src/assets/recipe-logo.svg" alt="Recipe Logo" class="w-9 h-9 transform hover:scale-110 transition-transform" />
        <div class="text-3xl font-bold text-white font-['Dancing_Script'] hover:text-emerald-200 transition-colors">
          Tuem Recipe
        </div>
      </div>

      <nav class="flex justify-end items-center">
        <ul
          ref="navigator"
          class="list-none h-screen md:h-auto divide-y divide-emerald-100/10 md:divide-y-0 items-center pt-16 md:pt-0 z-30 bg-gradient-to-r from-emerald-800 to-teal-800 md:bg-transparent text-white flex-col gap-2 overflow-hidden fixed md:relative top-0 right-0 w-1/2 md:w-auto -right-full md:right-0 duration-300 md:duration-[0ms] md:gap-3 md:flex-row flex"
        >
          <li
            @click="removeNavigator"
            v-for="item in pathes"
            class="cursor-pointer w-full md:w-auto px-4 md:px-3 text-center relative group"
            :key="item"
          >
            <router-link 
              :to="item.path" 
              class="flex items-center gap-2 text-white/90 hover:text-white font-medium text-lg transition-all duration-300 font-['Quicksand'] group-hover:scale-105"
            >
              <component :is="item.icon" class="w-5 h-5 transition-transform group-hover:rotate-6" />
              {{ item.name }}
            </router-link>
            <div class="absolute bottom-0 left-0 w-0 h-0.5 bg-emerald-400 transition-all duration-300 group-hover:w-full"></div>
          </li>

          <li
            v-if="!loginOut"
            class="cursor-pointer w-full md:w-fit px-6 md:px-0 text-center relative group"
          >
            <router-link 
              to="authentication" 
              class="flex items-center gap-2 text-white/90 hover:text-white font-medium text-lg transition-all duration-300 font-['Quicksand'] group-hover:scale-105"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 transition-transform group-hover:rotate-6" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z" clip-rule="evenodd" />
              </svg>
              LogIn
            </router-link>
            <div class="absolute bottom-0 left-0 w-0 h-0.5 bg-emerald-400 transition-all duration-300 group-hover:w-full"></div>
          </li>
        </ul>

        <span
          v-if="loginOut"
          @click="show_drop"
          class="border-2 border-emerald-400 cursor-pointer bg-emerald-700 text-center align-middle mr-8 md:ml-5 md:mr-0 text-2xl text-white w-9 h-9 font-mono rounded-full hover:bg-emerald-600 hover:border-emerald-300 transition-colors"
        >
          {{ loginOut && loginOut.toUpperCase().charAt(0) }}
        </span>

        <div
          v-if="loginOut"
          v-show="show_profile"
          ref="profile_value"
          class="w-44 fixed top-16 justify-around items-center right-1 h-36 drop-shadow-2xl shadow-2xl z-50 bg-white flex flex-col rounded-md"
        >
          <span class="text-lg text-center text-gray-600 border-2 w-8 h-8 rounded-full">
            {{loginOut && loginOut.toUpperCase().charAt(0)}}
          </span>
          <div class="flex flex-col items-center">
            <span class="text-base text-gray-600">{{ useState.state.name }}</span>
            <span class="text-sm text-gray-600">{{ useState.state.email }}</span>
          </div>
          <button
            @click="logout_event"
            class="shadow-lg drop-shadow-lg px-3 py-1 text-white rounded-md bg-emerald-700 font-semibold text-lg hover:bg-emerald-600 transition-colors"
          >
            LogOut
          </button>
        </div>
      </nav>

      <div
        v-if="bool"
        @click="burger"
        class="w-8 absolute right-[5%] top-5 z-50 flex gap-1 flex-col md:hidden cursor-pointer"
      >
        <div ref="bre_one" class="w-6 h-[0.2rem] duration-300 bg-white"></div>
        <div ref="bre_two" class="w-6 h-[0.2rem] bg-white"></div>
        <div ref="bre_three" class="w-6 h-[0.18rem] duration-300 bg-white"></div>
      </div>
    </header>
    <RouterView />
  </div>
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=Dancing+Script:wght@400;500;600;700&family=Quicksand:wght@300;400;500;600;700&display=swap');

/* Add smooth transitions for interactive elements */
.router-link-active {
  color: white;
  font-weight: 600;
}

.router-link-active svg {
  transform: scale(1.1);
}

/* Improve mobile menu animation */
@media (max-width: 767px) {
  .navigator-enter-active,
  .navigator-leave-active {
    transition: transform 0.3s ease-in-out;
  }

  .navigator-enter-from,
  .navigator-leave-to {
    transform: translateX(100%);
  }
}
</style>

<script setup>
import { ref } from "@vue/reactivity";
import { computed, onMounted, onUnmounted, watchEffect } from "@vue/runtime-core";
import { useStore } from "vuex";
import app from "../Axios/axiosconf";

// Updated icons with more detailed SVGs
const HomeIcon = {
  template: `<svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
    <path d="M10.707 2.293a1 1 0 00-1.414 0l-7 7a1 1 0 001.414 1.414L4 10.414V17a1 1 0 001 1h2a1 1 0 001-1v-2a1 1 0 011-1h2a1 1 0 011 1v2a1 1 0 001 1h2a1 1 0 001-1v-6.586l.293.293a1 1 0 001.414-1.414l-7-7z" />
  </svg>`
};

const CreateIcon = {
  template: `<svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
    <path d="M17.414 2.586a2 2 0 00-2.828 0L7 10.172V13h2.828l7.586-7.586a2 2 0 000-2.828z" />
    <path fill-rule="evenodd" d="M2 6a2 2 0 012-2h4a1 1 0 010 2H4v10h10v-4a1 1 0 112 0v4a2 2 0 01-2 2H4a2 2 0 01-2-2V6z" clip-rule="evenodd" />
  </svg>`
};

const BookmarkIcon = {
  template: `<svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
    <path d="M5 4a2 2 0 012-2h6a2 2 0 012 2v14l-5-2.5L5 18V4z" />
  </svg>`
};

const MyRecipeIcon = {
  template: `<svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
    <path d="M9 2a1 1 0 000 2h2a1 1 0 100-2H9z" />
    <path fill-rule="evenodd" d="M4 5a2 2 0 012-2 3 3 0 003 3h2a3 3 0 003-3 2 2 0 012 2v11a2 2 0 01-2 2H6a2 2 0 01-2-2V5zm3 4a1 1 0 000 2h.01a1 1 0 100-2H7zm3 0a1 1 0 000 2h3a1 1 0 100-2h-3zm-3 4a1 1 0 100 2h.01a1 1 0 100-2H7zm3 0a1 1 0 100 2h3a1 1 0 100-2h-3z" clip-rule="evenodd" />
  </svg>`
};

const bre_one = ref(null);
const bre_two = ref(null);
const bre_three = ref(null);
const navigator = ref(null);
import { useRouter } from "vue-router";
const router = useRouter();
const store = useStore();
const burger = () => {
  if (bre_one.value.classList.contains("rotate-45")) {
    bre_one.value.classList.remove("rotate-45");
    bre_two.value.classList.add("duration-300");
    bre_two.value.classList.remove("invisible");
    bre_three.value.classList.remove("-rotate-45");
    navigator.value.classList.remove("right-0");
    navigator.value.classList.add("-right-full");
    bre_three.value.classList.remove("translate-y-[-6px]");
    bre_one.value.classList.remove("translate-y-[8px]");
  } else {
    bre_one.value.classList.add("rotate-45");
    bre_two.value.classList.remove("duration-300");
    bre_one.value.classList.add("translate-y-[8px]");
    bre_two.value.classList.add("invisible");
    bre_three.value.classList.add("translate-y-[-6px]");
    bre_three.value.classList.add("-rotate-45");
    navigator.value.classList.remove("-right-full");
    navigator.value.classList.add("right-0");
  }
};
const profile_value = ref(null);
const show_profile = ref(false);

onMounted(() => {
  document.addEventListener("click", (e) => {
    if (loginOut) {
      const events = e.target;
      if (events == profile_value) {
        if (show_profile) {
          show_profile.value = false;
        }
      }
    }
  });
});

const show_drop = () => {
  show_profile.value = !show_profile.value;
};

const useState = useStore();

const bool = ref(true);

const resize = () => {
  if (window.innerWidth > 767) {
    bool.value = false;
    if (navigator.value.classList.contains("left-0")) {
      navigator.value.classList.remove("left-0");
      navigator.value.classList.add("left-[120%]");
    }
  } else {
    bool.value = true;
  }
};
const login = ref(useState.state.name);

const pathes = [
  { path: "/", name: "Home", icon: HomeIcon },
  { path: "recipe", name: "createRecipe", icon: CreateIcon },
  { path: "bookMarks", name: "BookMark", icon: BookmarkIcon },
  { path: "YourRecipe", name: "myRecipe", icon: MyRecipeIcon },
];
watchEffect(() => {
  login.value = useState.getters.getName;
});
const loginOut = computed({
  get() {
    return login.value;
  },
});

window.addEventListener("resize", resize);
onUnmounted(() => {
  window.removeEventListener("resize", resize);
});

const logout_event = async () => {
  show_profile.value = !show_profile.value;
  try {
    const data = await app.post("/user/logout");
    if (data.status === 200) {
      store.dispatch("setToken", "");
      store.dispatch("setId", "");
      store.dispatch("setEmail", "");
      store.dispatch("setName", "");
      router.push("/");
    }
  } catch (e) {}
};

const removeNavigator = () => {
  if (window.innerWidth <= 767) {
    if (bre_one.value && bre_one.value.classList.contains("rotate-45")) {
      bre_one.value.classList.toggle("rotate-45");
      bre_two.value.classList.toggle("duration-300");
      bre_two.value.classList.toggle("invisible");
      bre_three.value.classList.toggle("-rotate-45");
      navigator.value.classList.toggle("right-0");
      navigator.value.classList.toggle("-right-full");
      bre_three.value.classList.toggle("translate-y-[-6px]");
      bre_one.value.classList.toggle("translate-y-[8px]");
    }
  }
};
</script>
