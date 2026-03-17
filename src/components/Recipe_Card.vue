<template>
  <article class="group w-full bg-white rounded-xl shadow-lg overflow-hidden hover:shadow-2xl transition-all duration-300 transform hover:-translate-y-1">
    <!-- View Detail Button -->
    <div class="relative">
      <img 
        :src="props.src" 
        :alt="props.name"
        class="w-full h-48 object-cover transform group-hover:scale-110 transition-transform duration-500"
      />
      <div class="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent opacity-0 group-hover:opacity-100 transition-opacity"></div>
      <button
        @click="show(unique)"
        class="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 px-6 py-2 bg-emerald-600 text-white rounded-lg opacity-0 group-hover:opacity-100 transition-all duration-300 transform scale-95 group-hover:scale-100 font-['Quicksand'] hover:bg-emerald-700"
      >
        View Details
      </button>
    </div>

    <div class="p-4">
      <!-- Title and Description -->
      <div class="mb-4">
        <h3 class="text-xl font-semibold text-gray-900 mb-2 font-['Quicksand'] line-clamp-1 group-hover:text-emerald-600 transition-colors">
          {{ props.name }}
        </h3>
        <p class="text-gray-600 text-sm font-['Quicksand'] line-clamp-2">
          {{ props.description }}
        </p>
      </div>

      <!-- Duration -->
      <div class="flex justify-end mb-4">
        <span class="flex items-center gap-1 text-emerald-600 font-medium text-sm">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z" clip-rule="evenodd" />
          </svg>
          {{ props.duration }}
        </span>
      </div>

      <!-- Stats Row -->
      <div class="flex items-center justify-between text-sm text-gray-600 mb-4">
        <!-- Rating Stars -->
        <div class="flex items-center gap-1">
          <div class="flex">
            <template v-for="index in 5" :key="index">
              <svg 
                class="w-4 h-4"
                :class="index <= Math.round(props.rating) ? 'text-yellow-500' : 'text-gray-300'"
                xmlns="http://www.w3.org/2000/svg" 
                viewBox="0 0 20 20" 
                fill="currentColor"
              >
                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
              </svg>
            </template>
          </div>
          <span class="font-medium">{{ props.rating || 0 }}</span>
        </div>

        <!-- Comments -->
        <div class="flex items-center gap-1">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-emerald-500" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M18 13V5a2 2 0 00-2-2H4a2 2 0 00-2 2v8a2 2 0 002 2h3l3 3 3-3h3a2 2 0 002-2zM5 7a1 1 0 011-1h8a1 1 0 110 2H6a1 1 0 01-1-1zm1 3a1 1 0 100 2h3a1 1 0 100-2H6z" clip-rule="evenodd" />
          </svg>
          <span class="font-medium">{{ props.comments || 0 }}</span>
        </div>

        <!-- Likes -->
        <div class="flex items-center gap-1">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-emerald-500" viewBox="0 0 20 20" fill="currentColor">
            <path d="M2 10.5a1.5 1.5 0 113 0v6a1.5 1.5 0 01-3 0v-6zM6 10.333v5.43a2 2 0 001.106 1.79l.05.025A4 4 0 008.943 18h5.416a2 2 0 001.962-1.608l1.2-6A2 2 0 0015.56 8H12V4a2 2 0 00-2-2 1 1 0 00-1 1v.667a4 4 0 01-.8 2.4L6.8 7.933a4 4 0 00-.8 2.4z" />
          </svg>
          <span class="font-medium">{{ props.likes || 0 }}</span>
        </div>
      </div>

      <!-- Delete Button -->
      <button
        @click.stop="remove_your_recipe(props.unique)"
        class="w-full px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 transition-colors font-['Quicksand'] flex items-center justify-center gap-2"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
          <path fill-rule="evenodd" d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z" clip-rule="evenodd" />
        </svg>
        Delete Recipe
      </button>
    </div>
  </article>
</template>

<script setup>
const props = defineProps({
  src: String,
  description: String,
  name: String,
  duration: String,
  unique: Number,
  rating: Number,
  comments: Number,
  likes: Number,
});

import { useRouter } from "vue-router";
const use = useRouter();

function show(id) {
  use.push({ name: "detail", query: { productID: id } });
}

const emits = defineEmits(["delete_recipe"]);

const remove_your_recipe = (id) => {
  if (confirm('Are you sure you want to delete this recipe?')) {
    emits("delete_recipe", id);
  }
};
</script>

<style scoped>
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
</style>
