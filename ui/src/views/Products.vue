<template>
    <div>
      <h1 class="text-2xl font-bold mb-4">Products</h1>
      <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
        <ProductCard 
          v-for="p in products" 
          :key="p.id" 
          :product="p" 
          @order="createOrder" />
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted } from 'vue'
  import axios from 'axios'
  import ProductCard from '../components/ProductCard.vue'
  import { useAuthStore } from '../store/auth'
  
  const products = ref([])
  const auth = useAuthStore()
  
  onMounted(async () => {
    const res = await axios.get('http://localhost:8080/products')
    products.value = res.data
  })
  
  async function createOrder(product) {
    if (!auth.token) return alert("Login required")
    await axios.post('http://localhost:8080/orders', {
      items: [{ product_id: product.id, qty: 1 }]
    }, { headers: { Authorization: `Bearer ${auth.token}` } })
    alert("Order placed!")
  }
  </script>