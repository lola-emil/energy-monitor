import { createRouter, createWebHistory } from 'vue-router';
import { useAuthStore } from "@/stores/auth";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/auth",
      children: [
        {
          path: "/login",
          component: () => import("@/pages/LoginView.vue")
        },
      ]
    },
    {
      path: "/",
      children: [
        {
          path: "",
          component: () => import("@/pages/DashboardView.vue")
        },
        {
          path: "devices",
          component: () => import("@/pages/DevicePage.vue")
        },
        {
          path: "settings",
          component: () => import("@/pages/SettingsPage.vue")
        }
      ],
    }
  ],
});

router.beforeEach((to, from, next) => {
  const auth = useAuthStore();

  if (to.path !== "/login" && !auth.isAuthenticated) {
    return next("/login");
  }

  if (to.path === "/login" && auth.isAuthenticated) {
    return next("/");
  }

  next();
});

export default router;
