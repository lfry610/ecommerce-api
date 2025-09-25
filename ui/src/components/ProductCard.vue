<template>
    <div class="bg-white p-4 shadow rounded">
      <h2 class="text-lg font-semibold">{{ product.name }}</h2>
      <p class="text-gray-600">${{ product.price }}</p>
      <button 
        @click="$emit('order', product)" 
        class="mt-2 bg-blue-500 text-white px-3 py-1 rounded hover:bg-blue-600">
        Buy Now
      </button>
    </div>
  </template>
  
  <script setup>
  defineProps({ product: Object })
  </script>