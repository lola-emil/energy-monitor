import WSTestView from '@/views/ WSTestView.vue'
import AnalyticsView from '@/views/AnalyticsView.vue'
import CostMonitoringView from '@/views/CostMonitoringView.vue'
import DashboardView from '@/views/DashboardView.vue'
import DeviceListView from '@/views/DeviceListView.vue'
import LoginView from '@/views/LoginView.vue'
import SettingsView from '@/views/SettingsView.vue'
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      component: AnalyticsView
    },
    {
      path: "/login",
      component: LoginView
    },
    {
      path: "/devices",
      component: DeviceListView
    },
    {
      path: "/ws-test",
      component: WSTestView
    },
    {
      path: "/cost-monitor",
      component: CostMonitoringView
    },

    {
      path: "/settings",
      component: SettingsView
    },
  ],
})

export default router
