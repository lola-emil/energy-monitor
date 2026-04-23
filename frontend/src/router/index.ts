import { useAuthStore } from '@/stores/auth';
import { createRouter, createWebHistory } from 'vue-router';
import AnalyticsView from '@/views/dashboard/DashboardView.vue';
import DeviceAnalytics from '@/views/device-analytics/DeviceAnalytics.vue';
import LoginView from '@/views/auth/LoginView.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/test-dashboard",
      component: () => import("@/views/test/Dashboard.vue")
    },

    {
      path: "/test-analytics",
      component: () => import("@/views/test/DeviceAnalytics.vue")
    },

    {
      path: "/test-homepage",
      component: () => import("@/views/test/HomePage.vue")
    },

    {
      path: "/test-devicelist",
      component: () => import("@/views/test/DeviceList.vue")
    },

    {
      path: "/test-devicedetail",
      component: () => import("@/views/test/DeviceDetailPage.vue")
    },

    {
      path: "/test-settings",
      component: () => import("@/views/test/SettingsPage.vue")
    },

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
});


router.beforeEach((to, _from, next) => {
  const auth = useAuthStore();

  if (to.meta.requiresAuth && !auth.isAuthenticated) {
    return next('/auth/login');
  }
  if (to.meta.requiresGuest && auth.isAuthenticated) {
    return next('/');
  }

  next();
});


export default router;
