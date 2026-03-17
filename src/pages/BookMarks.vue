<template>
  <div class="min-h-screen bg-gradient-to-br from-gray-50 to-gray-100">
    <!-- Hero Section -->
    <div class="relative">
      <div class="relative flex items-center justify-center mt-16 h-[20rem] bg-[url('/src/assets/bookmark-hero.jpg')] bg-cover bg-center bg-no-repeat">
        <div class="absolute inset-0 bg-gradient-to-b from-black/60 to-black/40"></div>
        <div class="relative z-10 text-center">
          <h1 class="text-4xl md:text-5xl font-bold text-white mb-4 font-['Dancing_Script']">Your Favorite Recipes</h1>
          <p v-if="result" class="text-xl text-white/90 font-['Quicksand']">
            You have saved {{ computedCards?.getuserbyid[0]?.numberofBookMark }} delicious recipes
          </p>
        </div>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex justify-center items-center min-h-[400px]">
      <div class="flex flex-col items-center gap-4">
        <img class="w-20 h-20" src="../assets/loading.svg" alt="Loading" />
        <p class="text-gray-600 font-['Quicksand']">Loading your recipes...</p>
      </div>
    </div>

    <!-- No Bookmarks State -->
    <div v-if="result && computedCards?.getuserbyid[0]?.numberofBookMark === 0" class="flex justify-center items-center min-h-[400px]">
      <div class="text-center">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 text-gray-400 mx-auto mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z" />
        </svg>
        <h2 class="text-2xl font-semibold text-gray-800 mb-2 font-['Quicksand']">No saved recipes yet</h2>
        <p class="text-gray-600 mb-4 font-['Quicksand']">Start exploring and save your favorite recipes!</p>
        <router-link to="/" class="inline-flex items-center px-6 py-3 bg-emerald-600 text-white rounded-lg hover:bg-emerald-700 transition-colors font-['Quicksand']">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M9.707 16.707a1 1 0 01-1.414 0l-6-6a1 1 0 010-1.414l6-6a1 1 0 011.414 1.414L5.414 9H17a1 1 0 110 2H5.414l4.293 4.293a1 1 0 010 1.414z" clip-rule="evenodd" />
          </svg>
          Explore Recipes
        </router-link>
      </div>
    </div>

    <!-- Recipe Grid -->
    <main
      v-if="result && computedCards?.getuserbyid[0]?.numberofBookMark > 0"
      class="container mx-auto px-4 py-12 grid gap-6 xl:grid-cols-4 lg:grid-cols-3 md:grid-cols-2 grid-cols-1"
    >
      <div
        v-for="item in computedCards?.getuserbyid[0].getBookMark"
        :key="item.id"
        class="group bg-white rounded-xl shadow-lg overflow-hidden hover:shadow-2xl transition-all duration-300 transform hover:-translate-y-1"
      >
        <!-- Recipe Image -->
        <div class="relative overflow-hidden aspect-w-16 aspect-h-9">
          <img 
            :src="item.thum_image" 
            :alt="item.name"
            class="w-full h-48 object-cover transform group-hover:scale-110 transition-transform duration-500"
          />
          <div class="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
          
          <!-- Remove Bookmark Button -->
          <button 
            @click.prevent="delete_book_mark(item.id)"
            class="absolute top-3 right-3 p-2 bg-white/90 rounded-full opacity-0 group-hover:opacity-100 transition-opacity hover:bg-red-50"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-red-500" viewBox="0 0 20 20" fill="currentColor">
              <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd" />
            </svg>
          </button>
        </div>
        
        <!-- Recipe Details -->
        <div class="p-4">
          <router-link :to="'/recipe/' + item.id">
            <h3 class="text-xl font-semibold text-gray-800 mb-2 font-['Quicksand'] line-clamp-1 group-hover:text-emerald-600 transition-colors">{{ item.name }}</h3>
          </router-link>
          <p class="text-gray-600 text-sm mb-3 font-['Quicksand'] line-clamp-2">{{ item.description }}</p>
          
          <div class="flex items-center justify-between text-sm text-gray-500">
            <div class="flex items-center gap-4">
              <span class="flex items-center gap-1">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-emerald-500" viewBox="0 0 20 20" fill="currentColor">
                  <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                </svg>
                {{ item.numberofrate }}
              </span>
              <span class="flex items-center gap-1">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-emerald-500" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd" d="M18 13V5a2 2 0 00-2-2H4a2 2 0 00-2 2v8a2 2 0 002 2h3l3 3 3-3h3a2 2 0 002-2zM5 7a1 1 0 011-1h8a1 1 0 110 2H6a1 1 0 01-1-1zm1 3a1 1 0 100 2h3a1 1 0 100-2H6z" clip-rule="evenodd" />
                </svg>
                {{ item.numberofcomment }}
              </span>
            </div>
            <div class="flex items-center gap-1">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-emerald-500" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z" clip-rule="evenodd" />
              </svg>
              <span class="font-medium">{{ item.duration }} min</span>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref } from "@vue/reactivity";
import gql from "graphql-tag";
import { computed, onMounted } from "@vue/runtime-core";
import { useMutation, useQuery } from "@vue/apollo-composable";
import { useStore } from "vuex";

const store = useStore();

const markedRecipe = gql`
  query($uid: uuid!) {
    getuserbyid(args: { uid: $uid }) {
      name
      numberofBookMark
      getBookMark {
        name
        thum_image
        description
        numberofrate
        rate
        numberoflike
        numberofcomment
        duration
        createdBy
        id
      }
    }
  }
`;

const mark_remove = gql`
  mutation removing($uid: uuid, $recid: Int) {
    remove_bookmark(args: { uid: $uid, recid: $recid }) {
      userid
    }
  }
`;

const { mutate: mutant_remove_mark } = useMutation(mark_remove);

const delete_book_mark = async (id) => {
  if (store.state.id && id) {
    try {
      await mutant_remove_mark({ uid: store.state.id, recid: id });
      await update_change({ variables: { uid: store.state.id } });
    } catch (error) {
      console.error('Error removing bookmark:', error);
    }
  }
};

onMounted(() => {
  update_change({ variables: { uid: store.state.id } });
});

const variables = ref({ uid: store.state.id });
const { result, loading, fetchMore: update_change, error } = useQuery(
  markedRecipe,
  variables
);

const computedCards = computed(() => result.value);
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Dancing+Script:wght@400;500;600;700&family=Quicksand:wght@300;400;500;600;700&display=swap');

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

.aspect-w-16 {
  position: relative;
  padding-bottom: 56.25%;
}

.aspect-w-16 > * {
  position: absolute;
  height: 100%;
  width: 100%;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
}
</style>

<route lang="yaml">
meta:
  layout: header
</route>
