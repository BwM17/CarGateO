import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import UserView from '@/views/UserView.vue'
import CarView from '@/views/CarView.vue'
import LogView from '@/views/LogView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Dashboard',
      component: HomeView,
    },
    {
      path: '/cars',
      name: 'Car Management',
      component: CarView,
    },
    {
      path: '/user',
      name: 'User Management',
      component: UserView,
    },
    {
      path: '/logs',
      name: 'Logs',
      component: LogView,
    },
  ],
})

export default router
