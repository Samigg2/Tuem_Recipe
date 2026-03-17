<script setup>
import SignUp from "../components/Signup.vue";
import Login from "../components/Login.vue";
import { provide, ref } from "vue";

const show = ref(null);

const title = ref("Welcome Back");
const direction = ref("Already have an account?");
provide("func", loSign);

function loSign() {
  if (title.value === "Welcome Back") {
    title.value = "Join Us Today";
    direction.value = "Create your account";
  } else {
    title.value = "Welcome Back";
    direction.value = "Already have an account?";
  }
  show.value = !show.value;
}
</script>
<template>
  <div class="min-h-screen relative flex bg-gradient-to-br from-emerald-900 to-teal-800 justify-center items-center w-screen py-10">
    <div class="absolute inset-0 bg-[url('/src/assets/pattern.svg')] opacity-10"></div>
    
    <div class="relative z-10 flex flex-row w-[90%] md:w-[80%] lg:w-[70%] xl:w-[60%] rounded-2xl shadow-2xl bg-white overflow-hidden">
      <!-- Left Panel -->
      <div class="w-[40%] hidden sm:flex bg-gradient-to-br from-emerald-600 to-teal-600 h-full flex-col gap-6 justify-center items-center p-8">
        <div class="w-32 h-32 rounded-full bg-white/10 backdrop-blur-sm flex items-center justify-center">
          <img
            src="/src/assets/recipe-logo.svg"
            class="w-20 h-20"
            alt="Recipe Logo"
          />
        </div>

        <div class="text-center">
          <h1 class="text-white text-3xl font-['Dancing_Script'] font-bold mb-3">{{ title }}</h1>
          <p class="text-white/90 font-['Quicksand'] text-lg mb-6">{{ direction }}</p>
          <button
            @click="loSign()"
            class="rounded-full px-8 py-2 text-white border-2 border-white/80 font-['Quicksand'] transition-all duration-300 hover:bg-white hover:text-emerald-600 hover:scale-105"
          >
            {{ !show ? "Sign In" : "Sign Up" }}
          </button>
        </div>
      </div>

      <!-- Mobile Toggle -->
      <div class="sm:hidden w-full bg-emerald-600 p-4 text-center">
        <button
          @click="loSign()"
          class="text-white font-['Quicksand'] underline decoration-2 underline-offset-4"
        >
          {{ !show ? "Switch to Sign In" : "Switch to Sign Up" }}
        </button>
      </div>

      <!-- Right Panel -->
      <div class="flex-1 relative">
        <SignUp v-show="!show" />
        <Login v-show="show" />
      </div>
    </div>
  </div>
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=Dancing+Script:wght@400;500;600;700&family=Quicksand:wght@300;400;500;600;700&display=swap');

.auth-input {
  @apply w-full px-4 py-2 rounded-lg border border-gray-300 focus:border-emerald-500 focus:ring-2 focus:ring-emerald-200 outline-none transition-all duration-300;
}

.auth-button {
  @apply w-full py-3 rounded-lg bg-emerald-600 text-white font-medium hover:bg-emerald-700 transition-colors duration-300 focus:ring-4 focus:ring-emerald-200;
}
</style>
