import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import LoginPage from '@/views/login/Index.vue'
import SignupPage from '@/views/signup/Index.vue'
import UserActivation from '@/views/user-activation/Index.vue'
import WaitlistNotification from '@/views/waitlist/UserNotification.vue'
import { requireAuth } from '@/middlewares/auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'HomePage',
      component: HomeView,
      beforeEnter: requireAuth,
    },
    {
      path: '/login',
      name: 'Login',
      component: LoginPage,
    },
    {
      path: '/signup',
      name: 'Signup',
      component: SignupPage,
    },
    {
      path: '/activation',
      name: 'UserActivation',
      component: UserActivation,
    },
    {
      path: '/waitlist/notification',
      name: 'WaitlistNotification',
      component: WaitlistNotification,
    }
  ],
})

export default router
