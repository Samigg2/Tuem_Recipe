<script setup>
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import gql from 'graphql-tag';
import Icon_with from "../components/Icon_with.vue";
import { useStore } from "vuex";
import { useQuery, useMutation } from "@vue/apollo-composable";
import { watchEffect } from "@vue/runtime-core";
import Rating from "../components/Rating.vue";
import { hasDirectives } from "@apollo/client/utilities";
import RateDisplay from "../components/RateDisplay.vue";

const toggler = ref(true);
const useState = useStore();
const recipeID = useRoute().query;
const router = useRouter();
const add_delete_Book = gql`
  mutation($uid: uuid, $recid: Int) {
    addorremovebookmark(args: { uid: $uid, recid: $recid }) {
      recipeid
      userid
    }
  }
`;
const like_unlike_query = gql`
  mutation($useID: uuid, $recip: Int) {
    likeorunlik(args: { uid: $useID, rid: $recip }) {
      userid
      isliked
    }
  }
`;
const Rating_query = gql`
  mutation($uid: uuid, $recid: Int, $rate: Int) {
    rateadding(args: { uid: $uid, rid: $recid, rate_number: $rate }) {
      rate
    }
  }
`;

const addComment_query = gql`
  mutation($useID: uuid, $recip: Int, $text: String) {
    addcomment(args: { uid: $useID, rid: $recip, comm: $text }) {
      comment
      userid
    }
  }
`;
const { mutate: Rate_Dispatch } = useMutation(Rating_query);
const { mutate: Like_Unlike } = useMutation(like_unlike_query);

const {
  mutate: addBooKMark,
  loading: BooKLoading,
  error: error_Mark,
  onDone: complate_B_mark,
} = useMutation(add_delete_Book);

const variables = ref({
  recipeid: recipeID.productID,
  userid: useState.state.id ? useState.state.id : null,
});
const { mutate: add_comments } = useMutation(addComment_query);

let like = ref(4);
let comment = ref(5);
let timer = ref("9 m");
const favorite = ref();
const bookmark = ref();
const ClickedFavorite = () => {
  if (
    useState.state.id != "" &&
    useState.state.id != null &&
    recipeID.productID != null &&
    recipeID.productID != ""
  ) {
    Like_Unlike({ useID: useState.state.id, recip: recipeID.productID }).then((data) => {
      updateMark({
        variables: {
          recipeid: recipeID.productID,
          userid: useState.state.id ? useState.state.id : null,
        },
      });
      if (data.data.likeorunlik && data.data.likeorunlik[0].isliked) {
        favorite.value = "fill-yellow-800";
      } else {
        favorite.value = "";
      }
    });
  } else {
    router.push("/authentication");
  }
};
onMounted(() => {
  updateMark({
    variables: {
      recipeid: recipeID.productID,
      userid: useState.state.id ? useState.state.id : null,
    },
  });
});

const bookMarkClick = () => {
  toggler.value = !toggler.value;
  if (
    useState.state.id != "" &&
    useState.state.id != null &&
    recipeID.productID != null &&
    recipeID.productID != ""
  ) {
    addBooKMark({ uid: useState.state.id, recid: recipeID.productID })
      .then((result) => {
        updateMark({
          variables: {
            recipeid: recipeID.productID,
            userid: useState.state.id ? useState.state.id : null,
          },
        });

        if (result.data.addorremovebookmark.length > 0) {
          bookmark.value = "fill-yellow-800";
        } else bookmark.value = "";
      })
      .catch(() => {
        alert("error happened");
      });
  } else {
    router.push("/authentication");
  }
};

const getProductInfo = gql`
  query getRecip($recipeid: Int, $userid: uuid) {
    getrecipebyid(args: { recipeid: $recipeid }) {
      id
      name
      duration
      rate
      numberoflike
      numberofrate
      numberofcomment
      thum_image
      description
      createdBy

      getDirection {
        steps
      }
      getIngeredent {
        items
      }
      is_she_he_markit(args: { uid: $userid })
      userLikeRateComment(args: { uid: $userid }) {
        comment
        isliked
        rate
      }

      getImage {
        images
      }
      getReview {
        comment
        rate
        getReviewer
      }
    }
  }
`;

const { result, loading, error, fetchMore: updateMark, networkStatus } = useQuery(
  getProductInfo,
  variables
);

const show = ref([]);
const showImage = ref();
const allImage = ref([]);
const description = ref("");

watchEffect(() => {
  if (result && result.value) {
    const detail = result.value;
    show.value = [
      {
        number: result.value.getrecipebyid.numberoflike,
      },
      {
        number: result.value.getrecipebyid.rate,
      },
      {
        number: result.value.getrecipebyid.numberofcomment,
      },
      {
        number: result.value.getrecipebyid.duration,
      },
    ];
    comment.value = show.value[2].number;
    timer.value = show.value[3].number;
    like.value = show.value[0].number;

    showImage.value = result.value.getrecipebyid.thum_image;
    allImage.value =
      result.value.getrecipebyid.getImage.length > 0
        ? result.value.getrecipebyid.getImage[0].images
        : [];

    description.value = result.value.getrecipebyid.description;

    bookmark.value = result.value.getrecipebyid.is_she_he_markit ? "fill-yellow-800" : "";
    favorite.value = result.value.getrecipebyid.userLikeRateComment[0].isliked
      ? "fill-yellow-800"
      : "";
  }
});
const hasText = ref(null);
const put_comment = () => {
  if (hasText && (hasText.value+"").trim() !="") {
    if (
      useState.state.id != "" &&
      useState.state.id != null &&
      recipeID.productID != null &&
      recipeID.productID != ""
    ) {
      add_comments({
        useID: useState.state.id,
        recip: recipeID.productID,
        text: hasText.value,
      })
        .then((data) => {
          hasText.value = "";
          updateMark({
            variables: {
              recipeid: recipeID.productID,
              userid: useState.state.id ? useState.state.id : null,
            },
          });
        })
        .catch((e) => {});
    } else {
      router.push("/authentication");
    }
  }
};

const id = ref("1");
const changeImage = (src) => {
  showImage.value = src;
};

const rating_method = (rates) => {
  if (
    useState.state.id != "" &&
    useState.state.id != null &&
    recipeID.productID != null &&
    recipeID.productID != ""
  ) {
    Rate_Dispatch({
      uid: useState.state.id,
      recid: recipeID.productID,
      rate: rates,
    }).then((e) => {
      updateMark({
        variables: {
          recipeid: recipeID.productID,
          userid: useState.state.id ? useState.state.id : null,
        },
      });
    });
  } else {
    router.push("/authentication");
  }
};
</script>

<template>
  <div class="min-h-screen bg-gradient-to-b from-white to-slate-50">
    <!-- Loading State -->
    <div v-if="loading" class="flex w-full h-[80vh] justify-center items-center">
      <img class="w-20 h-20" src="../assets/loading.svg" alt="Loading..." />
    </div>

    <div v-else-if="result" class="container mx-auto px-4 py-8">
      <!-- Recipe Header -->
      <div class="relative rounded-xl overflow-hidden mb-8">
        <div class="aspect-w-16 aspect-h-9 max-h-[600px]">
          <img 
            :src="showImage" 
            :alt="result.getrecipebyid.name"
            class="w-full h-full object-cover"
          />
        </div>
        <div class="absolute inset-0 bg-gradient-to-t from-black/70 to-transparent"></div>
        <div class="absolute bottom-0 left-0 right-0 p-8">
          <h1 class="text-5xl font-bold text-white mb-4 font-['Dancing_Script']">{{ result.getrecipebyid.name }}</h1>
          <div class="flex flex-wrap gap-4 text-white">
            <div class="flex items-center gap-2">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              <span class="font-['Quicksand']">{{ result.getrecipebyid.duration }}</span>
            </div>
            <div class="flex items-center gap-2">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
              </svg>
              <span class="font-['Quicksand']">{{ result.getrecipebyid.numberoflike }} Likes</span>
            </div>
            <div class="flex items-center gap-2">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
              </svg>
              <span class="font-['Quicksand']">{{ result.getrecipebyid.numberofcomment }} Comments</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Recipe Content -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Left Column - Description and Ingredients -->
        <div class="lg:col-span-2 space-y-8">
          <!-- Image Gallery -->
          <div v-if="allImage && allImage.length > 0" class="bg-white rounded-xl p-6 shadow-sm">
            <h2 class="text-2xl font-bold text-emerald-800 mb-4 font-['Dancing_Script']">Gallery</h2>
            <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
              <div 
                v-for="(image, index) in allImage" 
                :key="index" 
                class="relative aspect-w-1 aspect-h-1 rounded-lg overflow-hidden cursor-pointer group"
                @click="changeImage(image)"
              >
                <img 
                  :src="image" 
                  :alt="'Recipe image ' + (index + 1)"
                  class="w-full h-full object-cover transform group-hover:scale-110 transition-transform duration-300"
                />
                <div class="absolute inset-0 bg-black/30 opacity-0 group-hover:opacity-100 transition-opacity"></div>
              </div>
            </div>
          </div>

          <!-- Description -->
          <div class="bg-white rounded-xl p-6 shadow-sm">
            <h2 class="text-2xl font-bold text-emerald-800 mb-4 font-['Dancing_Script']">Description</h2>
            <p class="text-gray-700 font-['Quicksand'] leading-relaxed">{{ result.getrecipebyid.description }}</p>
          </div>

          <!-- Ingredients -->
          <div class="bg-white rounded-xl p-6 shadow-sm">
            <h2 class="text-2xl font-bold text-emerald-800 mb-4 font-['Dancing_Script']">Ingredients</h2>
            <ul class="space-y-2">
              <li 
                v-for="ingredient in result.getrecipebyid.getIngeredent[0].items" 
                :key="ingredient"
                class="flex items-center gap-2 font-['Quicksand']"
              >
                <svg class="w-5 h-5 text-emerald-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                </svg>
                {{ ingredient }}
              </li>
            </ul>
          </div>

          <!-- Instructions -->
          <div class="bg-white rounded-xl p-6 shadow-sm">
            <h2 class="text-2xl font-bold text-emerald-800 mb-4 font-['Dancing_Script']">Instructions</h2>
            <ol class="space-y-4">
              <li 
                v-for="(step, index) in result.getrecipebyid.getDirection[0].steps" 
                :key="index"
                class="flex gap-4 font-['Quicksand']"
              >
                <span class="flex-shrink-0 w-8 h-8 bg-emerald-100 text-emerald-600 rounded-full flex items-center justify-center font-bold">
                  {{ index + 1 }}
                </span>
                <p class="text-gray-700 leading-relaxed">{{ step }}</p>
              </li>
            </ol>
          </div>
        </div>

        <!-- Right Column - Additional Info -->
        <div class="space-y-8">
          <!-- Author Info -->
          <div class="bg-white rounded-xl p-6 shadow-sm">
            <h2 class="text-2xl font-bold text-emerald-800 mb-4 font-['Dancing_Script']">Author</h2>
            <div class="flex items-center gap-4">
              <div class="w-16 h-16 rounded-full bg-emerald-100 flex items-center justify-center">
                <svg class="w-8 h-8 text-emerald-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                </svg>
              </div>
              <div>
                <h3 class="text-lg font-semibold text-gray-900 font-['Quicksand']">{{ result.getrecipebyid.createdBy }}</h3>
                <p class="text-gray-500 font-['Quicksand']">Recipe Creator</p>
              </div>
            </div>
          </div>

          <!-- Rating Section -->
          <div class="bg-white rounded-xl p-6 shadow-sm">
            <h2 class="text-2xl font-bold text-emerald-800 mb-4 font-['Dancing_Script']">Rating</h2>
            
            <!-- Overall Rating Display -->
            <div class="mb-6">
              <RateDisplay 
                :averageNumberOFRate="result.getrecipebyid.rate" 
                :numberofpeople="result.getrecipebyid.numberofrate"
              />
            </div>

            <!-- User Rating -->
            <div class="border-t pt-4">
              <h3 class="text-lg font-semibold text-gray-900 mb-3 font-['Quicksand']">Rate this recipe</h3>
              <Rating 
                :rate_number="result.getrecipebyid.userLikeRateComment[0].rate"
                @method="rating_method"
              />
            </div>
          </div>

          <!-- Actions -->
          <div class="bg-white rounded-xl p-6 shadow-sm space-y-4">
            <button 
              @click="ClickedFavorite"
              :class="{'bg-emerald-700': favorite === 'fill-yellow-800'}"
              class="w-full bg-emerald-600 text-white py-3 rounded-lg font-['Quicksand'] hover:bg-emerald-700 transition-colors flex items-center justify-center gap-2"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
              </svg>
              {{ favorite === 'fill-yellow-800' ? 'Liked' : 'Like Recipe' }}
            </button>
            <button 
              @click="bookMarkClick"
              :class="{'bg-emerald-50 border-emerald-700': bookmark === 'fill-yellow-800'}"
              class="w-full border-2 border-emerald-600 text-emerald-600 py-3 rounded-lg font-['Quicksand'] hover:bg-emerald-50 transition-colors flex items-center justify-center gap-2"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z" />
              </svg>
              {{ bookmark === 'fill-yellow-800' ? 'Saved' : 'Save Recipe' }}
            </button>
          </div>

          <!-- Nutrition Info -->
          <div class="bg-white rounded-xl p-6 shadow-sm">
            <h2 class="text-2xl font-bold text-emerald-800 mb-4 font-['Dancing_Script']">Nutrition Facts</h2>
            <div class="space-y-2">
              <div v-for="(value, label) in result.getrecipebyid.nutrition" :key="label"
                class="flex justify-between items-center py-2 border-b border-gray-100 last:border-0 font-['Quicksand']">
                <span class="text-gray-600">{{ label }}</span>
                <span class="font-medium text-gray-900">{{ value }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Comments Section -->
      <div class="mt-12 bg-white rounded-xl p-6 shadow-sm">
        <h2 class="text-2xl font-bold text-emerald-800 mb-6 font-['Dancing_Script']">Comments</h2>
        
        <!-- Comment Form -->
        <form @submit.prevent="put_comment" class="mb-8">
          <textarea
            v-model="hasText"
            rows="3"
            placeholder="Write a comment..."
            class="w-full p-4 rounded-lg border border-gray-200 focus:border-emerald-500 outline-none transition-colors font-['Quicksand'] resize-none"
          ></textarea>
          <div class="flex justify-end mt-4">
            <button
              type="submit"
              class="px-6 py-2 bg-emerald-600 text-white rounded-lg hover:bg-emerald-700 transition-colors font-['Quicksand'] disabled:opacity-50 disabled:cursor-not-allowed"
              :disabled="!hasText || hasText.trim() === ''"
            >
              Post Comment
            </button>
          </div>
        </form>

        <!-- Comments List -->
        <div class="space-y-6">
          <div v-for="comment in result.getrecipebyid.getReview" :key="comment.id" class="border-b border-gray-100 last:border-0 pb-6">
            <div class="flex items-start gap-4">
              <div class="w-10 h-10 rounded-full bg-emerald-100 flex items-center justify-center flex-shrink-0">
                <span class="text-emerald-600 font-semibold font-['Quicksand']">{{ comment.getReviewer.charAt(0).toUpperCase() }}</span>
              </div>
              <div class="flex-1">
                <div class="flex items-center gap-2 mb-1">
                  <h4 class="font-semibold text-gray-900 font-['Quicksand']">{{ comment.getReviewer }}</h4>
                  <span class="text-sm text-gray-500 font-['Quicksand']">{{ comment.date }}</span>
                </div>
                <p class="text-gray-700 font-['Quicksand']">{{ comment.comment }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<route lang="yaml">
meta:
  layout: header
</route>
