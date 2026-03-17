<script setup>
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import gql from "graphql-tag";
import { useMutation, useQuery } from "@vue/apollo-composable";
import { computed, onMounted, ref, watchEffect } from "@vue/runtime-core";
import Recipe_Card from "../components/Recipe_Card.vue";

const recipe_query = gql`
  query($uid: uuid) {
    getuserbyid(args: { uid: $uid }) {
      getYourRecipe {
        name
        duration
        description
        id
        thum_image
      }
    }
  }
`;

const remove_mutation = gql`
  mutation($uid: uuid, $recid: Int) {
    remove_recipes(args: { uid: $uid, recid: $recid }) {
      name
    }
  }
`;

const store = useStore();
const router = useRouter();

// Check authentication on mount
onMounted(() => {
  if (!store.state.id) {
    router.push('/authentication');
    return;
  }
  fetchRecipes();
});

const variables = ref({ uid: store.state.id });
const { result, loading: onLoading, error: onError, refetch } = useQuery(
  recipe_query,
  variables,
  {
    fetchPolicy: 'network-only',
    notifyOnNetworkStatusChange: true,
  }
);

const { mutate: remove_recipe_by_owener } = useMutation(remove_mutation);

const computednumber = computed(() => {
  if (result.value?.getuserbyid?.[0]) {
    return result.value;
  }
  return { getuserbyid: [{ getYourRecipe: [] }] };
});

const fetchRecipes = async () => {
  try {
    await refetch();
  } catch (error) {
    console.error('Error fetching recipes:', error);
  }
};

const remove_recipe = async (id) => {
  if (!id || !store.state.id) return;
  
  try {
    await remove_recipe_by_owener({ 
      uid: store.state.id, 
      recid: id 
    });
    await fetchRecipes();
  } catch (error) {
    console.error('Error removing recipe:', error);
  }
};
</script>
<template>
  <div class="relative pb-9 bg-gradient-to-b from-white to-slate-50">
    <!-- Hero Section -->
    <div class="relative rounded-b-3xl overflow-hidden">
      <div class="relative flex items-center justify-center h-[16rem] bg-[url('/src/assets/catering-2778755_1920.jpg')] bg-cover bg-center bg-no-repeat">
        <div class="absolute inset-0 bg-gradient-to-b from-black/60 to-black/40"></div>
        <div class="relative z-10 text-center">
          <h1 class="text-4xl md:text-5xl font-bold text-white mb-4 font-['Dancing_Script']">
            Welcome, {{ store.state.name }}
          </h1>
          <h2 class="text-2xl md:text-3xl text-white/90 font-['Quicksand']">
            Your Recipe Collection
          </h2>
        </div>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="onLoading" class="flex w-full h-[200px] justify-center items-center">
      <img class="w-20 h-20" src="../assets/loading.svg" alt="Loading..." />
    </div>

    <!-- Error State -->
    <div v-else-if="onError" class="flex flex-col items-center justify-center py-16">
      <div class="text-center max-w-md mx-auto p-6 bg-white rounded-lg shadow-lg">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 text-red-500 mb-4 mx-auto" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
        </svg>
        <h3 class="text-xl font-semibold text-gray-900 mb-2 font-['Quicksand']">Error Loading Recipes</h3>
        <p class="text-gray-600 mb-6 font-['Quicksand']">An error occurred while loading your recipes. Please try again.</p>
        <button 
          @click="fetchRecipes"
          class="px-6 py-2 bg-emerald-600 text-white rounded-lg hover:bg-emerald-700 transition-colors font-['Quicksand'] flex items-center gap-2 mx-auto"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
          Retry
        </button>
      </div>
    </div>

    <!-- No Recipes State -->
    <div v-else-if="!computednumber.getuserbyid[0]?.getYourRecipe?.length" class="flex flex-col items-center justify-center py-16">
      <div class="text-center">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 text-gray-400 mb-4 mx-auto" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 13h6m-3-3v6m-9 1V7a2 2 0 012-2h6l2 2h6a2 2 0 012 2v8a2 2 0 01-2 2H5a2 2 0 01-2-2z" />
        </svg>
        <h3 class="text-xl font-semibold text-gray-900 mb-2 font-['Quicksand']">No Recipes Yet</h3>
        <p class="text-gray-600 mb-6 font-['Quicksand']">Start creating your first recipe!</p>
        <a 
          href="/recipe" 
          class="inline-flex items-center px-6 py-3 bg-emerald-600 text-white rounded-lg hover:bg-emerald-700 transition-colors font-['Quicksand']"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
          </svg>
          Create Recipe
        </a>
      </div>
    </div>

    <!-- Recipe Grid -->
    <main
      v-else
      class="container mx-auto grid gap-6 px-4 mt-8 xl:grid-cols-4 lg:grid-cols-3 md:grid-cols-2 sm:grid-cols-2 grid-cols-1"
    >
      <Recipe_Card
        v-for="item in computednumber.getuserbyid[0].getYourRecipe"
        :key="item.id"
        :unique="item.id"
        :name="item.name"
        :description="item.description"
        :src="item.thum_image"
        :duration="item.duration"
        @delete_recipe="remove_recipe"
        class="w-full"
      />
    </main>
  </div>
</template>

<style scoped>
.grid {
  justify-items: center;
}

@media (min-width: 640px) {
  .grid {
    justify-items: stretch;
  }
}
</style>

<route lang="yaml">
meta:
  layout: header
</route>
