import { defineStore } from 'pinia'
import axios from 'axios'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: localStorage.getItem('token') || null,
    user: null,
  }),
  actions: {
    async login(email, password) {
      const res = await axios.post('http://localhost:8080/login', { email, password })
      this.token = res.data.token
      localStorage.setItem('token', this.token)
    },
    async register(email, password) {
      await axios.post('http://localhost:8080/register', { email, password })
    },
    logout() {
      this.token = null
      localStorage.removeItem('token')
    }
  }
})
