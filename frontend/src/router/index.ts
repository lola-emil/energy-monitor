import { createRouter, createWebHistory } from 'vue-router';
import { useAuthStore } from "@/stores/auth"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [


    {
      path: "/auth",
      component: () => import("@/layouts/AuthLayout.vue"),
      children: [
        {
          path: "/login",
          component: () => import("@/pages/LoginView.vue")
        },
      ]
    },

    {
      path: "/",
      component: () => import("@/layouts/MainLayout.vue"),
      children: [
        {
          path: "",
          component: () => import("@/pages/Overview.vue")
        },
        {
          path: "/settings",
          component: () => import("@/pages/SettingsView.vue")
        },

        {
          path: "/analytics",
          component: () => import("@/pages/AnalyticsView.vue")
        },
        {
          path: "/devices",
          component: () => import("@/pages/DeviceListView.vue"),
        },

        {
          path: "/devices/:id",
          component: () => import("@/pages/DeviceAnalyticsView.vue"),
        },

        {
          path: "/account-settings",
          component: () => import("@/pages/ProfileSettings.vue")
        }
      ]
    },
  ],
});

router.beforeEach((to, from, next) => {
  const auth = useAuthStore()

  if (to.path !== "/login" && !auth.isAuthenticated) {
    return next("/login")
  }

  if (to.path === "/login" && auth.isAuthenticated) {
    return next("/")
  }

  next()
})

export default router;
