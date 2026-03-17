<template>
  <div class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50">
    <div class="bg-white rounded-lg shadow-lg p-8 max-w-md w-full relative">
      <button @click="$emit('close')" class="absolute top-2 right-2 text-gray-400 hover:text-gray-700 text-2xl">&times;</button>
      <h2 class="text-2xl font-bold mb-4 text-emerald-700">Buy Recipe</h2>
      <div class="mb-4">
        <p class="text-lg font-semibold">Recipe: <span class="text-gray-800">{{ recipe.name }}</span></p>
        <p class="text-lg font-semibold">Price: <span class="text-emerald-700">{{ price }} ETB</span></p>
        <p class="text-gray-600 mt-2">You are about to purchase this recipe. After payment, you will get full access to all instructions and details.</p>
      </div>
      <button @click="proceedToChapa" class="w-full bg-blue-600 text-white py-3 rounded-lg font-['Quicksand'] hover:bg-blue-700 transition-colors mt-4">Proceed to Payment</button>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  price: Number,
  recipe: Object,
});

function proceedToChapa() {
  // Replace with actual Chapa payment URL and parameters
  window.location.href = `https://chapa.co/pay`; // Placeholder
}
</script> 