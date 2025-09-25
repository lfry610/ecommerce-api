<template>
    <div class="max-w-md mx-auto bg-white shadow rounded-lg p-6 mt-10">
      <h1 class="text-2xl font-bold mb-6">Login</h1>
  
      <form @submit.prevent="handleLogin" class="space-y-4">
        <!-- Email -->
        <div>
          <label for="email" class="block text-sm font-medium text-gray-700">Email</label>
          <input
            id="email"
            v-model="email"
            type="email"
            required
            placeholder="you@example.com"
            class="mt-1 w-full border rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
          />
        </div>
  
        <!-- Password -->
        <div>
          <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
          <input
            id="password"
            v-model="password"
            type="password"
            required
            placeholder="••••••••"
            class="mt-1 w-full border rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
          />
        </div>
  
        <!-- Error message -->
        <p v-if="error" class="text-red-600 text-sm">{{ error }}</p>
  
        <!-- Submit -->
        <button
          type="submit"
          class="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700 transition"
        >
          Log In
        </button>
      </form>
  
      <p class="mt-4 text-sm text-gray-600">
        Don’t have an account?
        <router-link to="/register" class="text-blue-600 hover:underline">
          Register
        </router-link>
      </p>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  import { useRouter } from 'vue-router'
  import { useAuthStore } from '../store/auth'
  
  const email = ref('')
  const password = ref('')
  const error = ref(null)
  
  const router = useRouter()
  const auth = useAuthStore()
  
  async function handleLogin() {
    error.value = null
    try {
      await auth.login(email.value, password.value)
      router.push('/') // Redirect to products page after login
    } catch (err) {
      error.value = err.response?.data?.error || "Login failed"
    }
  }
  </script>