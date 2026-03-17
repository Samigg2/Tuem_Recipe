<script setup>
const props = defineProps({ 
  rating: {
    type: Number,
    default: 0
  },
  numberOfRates: {
    type: Number,
    default: 0
  }
});
</script>

<template>
  <div class="flex flex-row gap-1">
    <template v-for="index in 5" :key="index">
      <svg
        class="w-4 h-4"
        :class="index <= (props.rating || 0) ? 'fill-yellow-400' : 'fill-gray-300'"
        viewBox="0 0 24 24"
      >
        <path d="M12 17.27 18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z" />
      </svg>
    </template>
  </div>
</template>
