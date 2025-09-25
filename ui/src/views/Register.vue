<template>
    <div class="max-w-md mx-auto bg-white shadow rounded-lg p-6 mt-10">
      <h1 class="text-2xl font-bold mb-6">Create an Account</h1>
  
      <form @submit.prevent="handleRegister" class="space-y-4">
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
  
        <!-- Confirm Password -->
        <div>
          <label for="confirm" class="block text-sm font-medium text-gray-700">Confirm Password</label>
          <input
            id="confirm"
            v-model="confirmPassword"
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
          Register
        </button>
      </form>
  
      <p class="mt-4 text-sm text-gray-600">
        Already have an account?
        <router-link to="/login" class="text-blue-600 hover:underline">
          Log in
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
  const confirmPassword = ref('')
  const error = ref(null)
  
  const router = useRouter()
  const auth = useAuthStore()
  
  async function handleRegister() {
    error.value = null
    if (password.value !== confirmPassword.value) {
      error.value = "Passwords do not match"
      return
    }
  
    try {
      await auth.register(email.value, password.value)
      alert("Account created successfully! Please log in.")
      router.push('/login')
    } catch (err) {
      error.value = err.response?.data?.error || "Registration failed"
    }
  }
  </script>