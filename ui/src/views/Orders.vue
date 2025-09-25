<template>
    <div>
      <h1 class="text-2xl font-bold mb-4">My Orders</h1>
      <ul>
        <li v-for="o in orders" :key="o.id" class="p-2 border-b">
          Order #{{ o.id }} - ${{ o.total }}
        </li>
      </ul>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted } from 'vue'
  import axios from 'axios'
  import { useAuthStore } from '../store/auth'
  
  const orders = ref([])
  const auth = useAuthStore()
  
  onMounted(async () => {
    if (!auth.token) return
    const res = await axios.get('http://localhost:8080/orders', {
      headers: { Authorization: `Bearer ${auth.token}` }
    })
    orders.value = res.data
  })
  </script>