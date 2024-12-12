import { createRouter, createWebHistory } from 'vue-router'
import DefaultLayout from '@/layout/Index.vue'
import HomeView from '@/views/HomeView.vue'
import ListStore from '@/views/stores/Index.vue'
import LoginPage from '@/views/login/Index.vue'
import SignupPage from '@/views/signup/Index.vue'
import UserActivation from '@/views/user-activation/Index.vue'
import WaitlistNotification from '@/views/waitlist/UserNotification.vue'

import AdminListUser from '@/views/admin/users/Index.vue'
import { requireAuth } from '@/middlewares/auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
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
    },
    {
      path: '/',
      component: DefaultLayout,
      beforeEnter: requireAuth,
      children: [
        {
          path: '',
          name: 'HomePage',
          component: HomeView,
        },
        {
          path: 'stores',
          name: 'ListStore',
          component: ListStore,
        },
        {
          path: 'admin/users',
          name: 'AdminListUser',
          component: AdminListUser,
        },
      ],
    },
  ],
})

export default router
