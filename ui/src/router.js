import { createRouter, createWebHistory } from 'vue-router'
import Login from './views/Login.vue'
import Register from './views/Register.vue'
import Products from './views/Products.vue'
import Orders from './views/Orders.vue'

const routes = [
  { path: '/', component: Products },
  { path: '/login', component: Login },
  { path: '/register', component: Register },
  { path: '/orders', component: Orders },
]

export default createRouter({
  history: createWebHistory(),
  routes,
})