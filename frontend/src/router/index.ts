import { useAuthStore } from '@/stores/auth'
import AnalyticsView from '@/views/AnalyticsView.vue'
import DeviceAnalytics from '@/views/DeviceAnalytics.vue'
import LoginView from '@/views/LoginView.vue'
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      component: () => import("@/layout/MainLayout.vue"),
      meta: {
        requiresAuth: true
      },
      children: [
        {
          path: "",
          component: AnalyticsView
        },
        {
          path: "devices",
          children: [
            {
              path: "",
              component: DeviceAnalytics
            },
            {
              path: ":id",
              component: DeviceAnalytics
            }
          ]
        },
      ]
    },
    {
      path: "/auth",
      component: () => import("@/layout/AuthLayout.vue"),
      meta: {
        requiresGuest: true
      },
      children: [
        {
          path: "login",
          component: LoginView
        }
      ]
    },

  ],
})


router.beforeEach((to, _from, next) => {
  const auth = useAuthStore()

  if (to.meta.requiresAuth && !auth.isAuthenticated) {
    return next('/auth/login')
  }
  if (to.meta.requiresGuest && auth.isAuthenticated) {
    return next('/')
  }

  next()
})


export default router
