<script setup>
import { ref } from "@vue/reactivity";
import { computed, watchEffect } from "@vue/runtime-core";
import { useForm, useField } from "vee-validate";
import { useStore } from "vuex";
import Input from "../components/Input.vue";
import InputFile from "../components/FileInput.vue";
import TextArea from "../components/TextArea.vue";
import Dropdawn from "../components/Dropdawn.vue";
import InNumber from "../components/InNumber.vue";
import axios from "../Axios/axiosconf";
const store = useStore();
import { gql } from "graphql-tag";
import { useMutation } from "@vue/apollo-composable";

const loading = ref(false);

const mutant = gql`
  mutation(
    $cataid: Int!
    $collectionimage: String!
    $description: String!
    $duration: String!
    $ingred: String!
    $step: String!
    $thumbimage: String!
    $title: String!
    $userid: uuid
    $price: Float!
  ) {
    insertrecipe(
      args: {
        catid: $cataid
        collectionimage: $collectionimage
        description: $description
        duration: $duration
        ingred: $ingred
        step: $step
        thumbimage: $thumbimage
        title: $title
        userid: $userid
        price: $price
      }
    ) {
      name
    }
  }
`;

const { mutate } = useMutation(mutant);

const steps = ref([]);

const warning_for_ingredient = ref();
const schemeValidate = {
  nameValidate: (value) => {
    if (value && value.trim()) return true;
    else return "please Enter the Title";
  },
  descriptionValidate: (value) => {
    if (value && value.trim()) return true;
    else return "please Enter the description";
  },

  prep: (value) => {
    if (value && !isNaN(Number(value)) && Number(value) > 0) {
      return true;
    } else if (Number(value) < 0) return "must not negative";
    else return "Enter the number ";
  },

  ingredient: (value) => {
    warning_for_ingredient.value = "";

    if (ingradientInput.value && ingradientInput.value.length >= 1) return true;
    else return "at least one ingredient is required";
  },
  imageValidate: (value) => {
    if (images.value && images.value.length > 0) return true;
    else return "Image is required";
  },

  stepValidate: (value) => {
    stepwarning.value = "";
    if (steps.value.length >= 2) return true;
    else return " at least two step is required";
  },

  price: (value) => {
    if (value && !isNaN(Number(value)) && Number(value) > 0) {
      return true;
    } else if (Number(value) < 0) return "Price must not be negative";
    else return "Enter the price";
  },
};

const ingradientInput = ref([]);
const stepwarning = ref();

function addIngredient() {
  if (ingradient.value != "" && ingradient.value != undefined&&(ingradient.value+"").trim() != "") {
    ingradientInput.value.push({
      value: ingradient.value,
      id: Date.now(),
    });
    ingradient.value = "";
  } else {
    warning_for_ingredient.value = "Please Enter the ingredient";
  }
}
function addStep() {
  if (step_input.value != "" && (step_input.value+"").trim() != "" && step_input.value != undefined) {
    steps.value.push({ value: step_input.value, id: Date.now() });
    step_input.value = "";
  } else {
    stepwarning.value = "Please Enter the step";
  }
}
// it should be choised
 const choise_catagory =ref([ { title: "Launch", id: 249 },
  { title: "BreakFast", id: 250 },
  { title: "Dinner", id: 251 },
  { title: "Desert", id: 252 },
  { title: "Appitizer", id: 253 },
  { title: "SideDish", id: 254 },
  { title: "Quick and easy", id: 255 },
  { title: "Holiday", id: 256 },
  { title: "Soup", id: 257 },
  { title: "vegeterain", id: 258 },
  { title: "Salad", id: 259 },]);
  // it should be choised
const images = ref([]);
const thumbImage = ref("/src/assets/ground.jpeg");
const imagename = ref();

const showImage = ref(false);

function review(event) {
  if (event.target.files) {
    const files = [...event.target.files];
    image_warning.value = "";
    images.value = [];
    imagename.value = [];

    files.forEach((file) => {
      imagename.value.push(file.name);
      readAndPreview(file);
    });
    showImage.value = true;
  }

  function readAndPreview(file, index) {
    var reader = new FileReader();
    reader.addEventListener("load", function () {
      images.value.push(this.result);
      thumbImage.value = images.value[0];
    });

    reader.readAsDataURL(file);
  }
}

let key = 2;

const bool = true;

const image_warning = ref();

const { handleSubmit, resetForm } = useForm({ validationSchema: schemeValidate });

const { value: title, errorMessage: errorTitle } = useField("nameValidate");
const { value: description, errorMessage: errorDescription } = useField(
  "descriptionValidate"
);

const { value: ingradient, errorMessage: errorIngredient } = useField("ingredient");
const { value: step_input, errorMessage: errorStep } = useField("stepValidate");
const { value: imagevalue, errorMessage: errMage } = useField("imageValidate");
const { value: prep, errorMessage: errorPrep } = useField("prep");
const { value: price, errorMessage: errorPrice } = useField("price");

const catagory = ref("249");

let prep_Time = ref("min");
const removeIngerdeint = (inde) => {
  ingradientInput.value = ingradientInput.value.filter((value) => {
    return value.id != inde;
  });
};

function removestep(index) {
  steps.value = steps.value.filter((value) => {
    return value.id != index;
  });
}
function editStep(index) {
  let value = steps.value.find((e) => {
    return e.id == index;
  });
  step_input.value = value.value;

  steps.value = steps.value.filter((value) => {
    return value.id != index;
  });
}

function editIngredient(index) {
  let value = ingradientInput.value.find((e) => {
    return e.id == index;
  });
  ingradient.value = value.value;

  ingradientInput.value = ingradientInput.value.filter((value) => {
    return value.id != index;
  });
}

const imageLink = ref("");
const temporaryImage = ref();
const clearAll=()=>{
 images.value = [];
    ingradientInput.value = [];
    steps.value = [];
    resetForm();
}

const upload = async () => {
  try {
    if (images.value.length >= 1) {
      const featureIndex = images.value.indexOf(thumbImage.value);
      for (let x = 0; x < images.value.length; x++) {
        const uploading = await axios.post("/user/uploadImage", {
          filename: imagename.value[x],
          base64img: images.value[x].split(",")[1],
        });
        if (x == 0) imageLink.value = uploading.data.url;
        else imageLink.value = imageLink.value + "$#@!%" + uploading.data.url;

        if (featureIndex == x) {
          temporaryImage.value = uploading.data.url;
        }
      }
    }
  } catch (e) {}
};

function choose_feature(value) {
  thumbImage.value = value;
}

const submit = handleSubmit(
  async (values) => {
    loading.value = true;
    await upload();

    let upload_ingredeint_text = "";
    ingradientInput.value.forEach((element, index) => {
      if (index === 0) upload_ingredeint_text = `${element.value}`;
      else upload_ingredeint_text = `${upload_ingredeint_text} $#@!% ${element.value}`;
    });

    let upload_step_text = "";
    steps.value.forEach((element, index) => {
      if (index === 0) upload_step_text = `${element.value}`;
      else upload_step_text = `${upload_step_text} $#@!% ${element.value}`;
    });
    mutate({
      cataid: Number(catagory.value),
      collectionimage: imageLink.value,
      description: values.descriptionValidate,
      duration: values.prep + " " + prep_Time.value,
      ingred: upload_ingredeint_text,
      step: upload_step_text,
      thumbimage: temporaryImage.value,
      title: values.nameValidate,
      userid: store.state.id,
      price: values.price,
    });
    images.value = [];
    ingradientInput.value = [];
    steps.value = [];
    resetForm();
    loading.value = false;
  },
  (error) => {
    if (ingradientInput.value.length < 2) {
      warning_for_ingredient.value = "at Least two ingredient is required";
    }
    if (steps.value.length < 1) {
      stepwarning.value = "at leat one Step is required";
    }
    if (images.value.length < 1) {
      image_warning.value = "Image is required";
    }
  }
);
</script>

<template>
  <div class="min-h-screen bg-gradient-to-br from-gray-50 to-gray-100">
    <!-- Hero Section -->
    <div class="relative">
      <div class="relative flex items-center justify-center mt-16 h-[20rem] bg-[url('/src/assets/recipe-hero.jpg')] bg-cover bg-center bg-no-repeat">
        <div class="absolute inset-0 bg-gradient-to-b from-black/60 to-black/40"></div>
        <div class="relative z-10 text-center">
          <h1 class="text-4xl md:text-5xl font-bold text-white mb-4 font-['Dancing_Script']">Create Your Recipe</h1>
          <p class="text-xl text-white/90 font-['Quicksand']">Share your culinary masterpiece with the world</p>
        </div>
      </div>
    </div>

    <!-- Form Section -->
    <div class="container mx-auto px-4 py-12">
      <form @submit.prevent="submit" class="max-w-4xl mx-auto bg-white rounded-2xl shadow-xl p-8">
        <!-- Title and Category -->
        <div class="grid md:grid-cols-2 gap-6 mb-8">
          <div class="relative">
            <label class="block text-gray-700 text-sm font-medium mb-2 font-['Quicksand']">Recipe Title</label>
            <Input
              v-model="title"
              type="text"
              placeholder="Enter your recipe name"
              class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:border-emerald-500 focus:ring-2 focus:ring-emerald-200 outline-none transition-all duration-300 font-['Quicksand']"
            />
            <span class="absolute text-sm text-red-600 mt-1 font-['Quicksand']">{{ errorTitle }}</span>
          </div>

          <div class="relative">
            <label class="block text-gray-700 text-sm font-medium mb-2 font-['Quicksand']">Category</label>
            <select
              v-model="catagory"
              class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:border-emerald-500 focus:ring-2 focus:ring-emerald-200 outline-none transition-all duration-300 font-['Quicksand'] appearance-none bg-white"
            >
              <option v-for="cata in choise_catagory" :key="cata.id" :value="cata.id">
                {{ cata.title }}
              </option>
            </select>
          </div>
        </div>

        <!-- Description -->
        <div class="mb-8">
          <label class="block text-gray-700 text-sm font-medium mb-2 font-['Quicksand']">Description</label>
          <textarea
            v-model="description"
            placeholder="Describe your recipe..."
            class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:border-emerald-500 focus:ring-2 focus:ring-emerald-200 outline-none transition-all duration-300 font-['Quicksand'] h-32 resize-none"
          ></textarea>
          <span class="text-sm text-red-600 mt-1 font-['Quicksand']">{{ errorDescription }}</span>
        </div>

        <!-- Images Upload -->
        <div class="mb-8">
          <label class="block text-gray-700 text-sm font-medium mb-2 font-['Quicksand']">Recipe Images</label>
          <div class="border-2 border-dashed border-gray-300 rounded-lg p-6 text-center">
            <InputFile :isMultiple="bool" v-model="imagevalue" :review="review" />
            <p class="text-sm text-gray-500 mt-2 font-['Quicksand']">Upload multiple images (max 5)</p>
            <span class="text-sm text-red-600 mt-1 font-['Quicksand']">{{ image_warning }}</span>
          </div>

          <!-- Image Preview -->
          <div v-if="images.length > 0" class="mt-4">
            <label class="block text-gray-700 text-sm font-medium mb-2 font-['Quicksand']">Select Thumbnail</label>
            <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
              <div v-for="(item, index) in images" :key="index" class="relative group">
                <img :src="item" class="w-full h-32 object-cover rounded-lg" />
                <div class="absolute inset-0 bg-black/50 opacity-0 group-hover:opacity-100 transition-opacity rounded-lg flex items-center justify-center">
                  <input
                    type="radio"
                    v-model="thumbImage"
                    :value="item"
                    @change="choose_feature(item)"
                    class="w-5 h-5 accent-emerald-500"
                  />
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Ingredients -->
        <div class="mb-8">
          <h2 class="text-2xl font-semibold mb-4 font-['Dancing_Script'] text-gray-800">Ingredients</h2>
          <div class="flex gap-4 mb-4">
            <Input
              v-model="ingradient"
              type="text"
              placeholder="e.g., 2 cups flour"
              class="flex-1"
            />
            <button
              type="button"
              @click="addIngredient"
              class="px-6 py-2 bg-emerald-600 text-white rounded-lg hover:bg-emerald-700 transition-colors font-['Quicksand']"
            >
              Add
            </button>
          </div>
          <span class="text-sm text-red-600 mt-1 font-['Quicksand']">{{ warning_for_ingredient }}</span>

          <!-- Ingredients List -->
          <div class="space-y-2">
            <div
              v-for="item in ingradientInput"
              :key="item.id"
              class="flex items-center justify-between p-3 bg-gray-50 rounded-lg group hover:bg-gray-100 transition-colors"
            >
              <span class="font-['Quicksand']">{{ item.value }}</span>
              <div class="flex gap-2">
                <button
                  @click="editIngredient(item.id)"
                  type="button"
                  class="text-emerald-600 hover:text-emerald-700"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z" />
                  </svg>
                </button>
                <button
                  @click="removeIngerdeint(item.id)"
                  type="button"
                  class="text-red-600 hover:text-red-700"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd" d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z" clip-rule="evenodd" />
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- Steps -->
        <div class="mb-8">
          <h2 class="text-2xl font-semibold mb-4 font-['Dancing_Script'] text-gray-800">Cooking Steps</h2>
          <div class="flex gap-4 mb-4">
            <Input
              v-model="step_input"
              type="text"
              placeholder="Describe the step..."
              class="flex-1"
            />
            <button
              type="button"
              @click="addStep"
              class="px-6 py-2 bg-emerald-600 text-white rounded-lg hover:bg-emerald-700 transition-colors font-['Quicksand']"
            >
              Add Step
            </button>
          </div>
          <span class="text-sm text-red-600 mt-1 font-['Quicksand']">{{ stepwarning }}</span>

          <!-- Steps List -->
          <div class="space-y-2">
            <div
              v-for="(step, index) in steps"
              :key="step.id"
              class="flex items-start gap-4 p-4 bg-gray-50 rounded-lg group hover:bg-gray-100 transition-colors"
            >
              <span class="flex-none w-8 h-8 flex items-center justify-center bg-emerald-600 text-white rounded-full font-['Quicksand']">
                {{ index + 1 }}
              </span>
              <div class="flex-1">
                <p class="font-['Quicksand']">{{ step.value }}</p>
              </div>
              <div class="flex gap-2">
                <button
                  @click="editStep(step.id)"
                  type="button"
                  class="text-emerald-600 hover:text-emerald-700"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z" />
                  </svg>
                </button>
                <button
                  @click="removestep(step.id)"
                  type="button"
                  class="text-red-600 hover:text-red-700"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd" d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z" clip-rule="evenodd" />
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- Duration -->
        <div class="mb-8">
          <h2 class="text-2xl font-semibold mb-4 font-['Dancing_Script'] text-gray-800">Cooking Time</h2>
          <div class="flex gap-4 items-center">
            <div class="flex-1">
              <Input
                v-model="prep"
                type="number"
                placeholder="Duration"
                class="w-full"
              />
              <span class="text-sm text-red-600 mt-1 font-['Quicksand']">{{ errorPrep }}</span>
            </div>
            <select
              v-model="prep_Time"
              class="px-4 py-2 rounded-lg border border-gray-300 focus:border-emerald-500 focus:ring-2 focus:ring-emerald-200 outline-none transition-all duration-300 font-['Quicksand']"
            >
              <option value="min">Minutes</option>
              <option value="hr">Hours</option>
            </select>
          </div>
        </div>

        <!-- Price -->
        <div class="mb-8">
          <h2 class="text-2xl font-semibold mb-4 font-['Dancing_Script'] text-gray-800">Price</h2>
          <div class="flex gap-4 items-center">
            <div class="flex-1">
              <InNumber
                v-model="price"
                name="price"
                type="number"
                placeholder="Enter price in ETB"
                unique="recipe-price"
              />
              <span class="text-sm text-red-600 mt-1 font-['Quicksand']">{{ errorPrice }}</span>
            </div>
          </div>
        </div>

        <!-- Submit Buttons -->
        <div class="flex gap-4 justify-end">
          <button
            type="button"
            @click="clearAll"
            class="px-6 py-2 border-2 border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors font-['Quicksand']"
          >
            Clear All
          </button>
          <button
            type="submit"
            class="px-8 py-2 bg-emerald-600 text-white rounded-lg hover:bg-emerald-700 transition-colors font-['Quicksand'] flex items-center gap-2"
            :disabled="loading"
          >
            <span>Create Recipe</span>
            <svg v-if="loading" class="animate-spin h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=Dancing+Script:wght@400;500;600;700&family=Quicksand:wght@300;400;500;600;700&display=swap');

/* Add smooth transitions */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

/* Custom scrollbar */
::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}

::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 4px;
}

::-webkit-scrollbar-thumb {
  background: #888;
  border-radius: 4px;
}

::-webkit-scrollbar-thumb:hover {
  background: #555;
}

/* Input focus styles */
input:focus, textarea:focus, select:focus {
  box-shadow: 0 0 0 2px rgba(16, 185, 129, 0.2);
}

/* Placeholder styles */
::placeholder {
  color: #9ca3af;
  opacity: 1;
}

/* Number input styles */
input[type="number"]::-webkit-inner-spin-button,
input[type="number"]::-webkit-outer-spin-button {
  -webkit-appearance: none;
  margin: 0;
}

input[type="number"] {
  -moz-appearance: textfield;
}
</style>

<route lang="yaml">
meta:
  layout: header
</route>
