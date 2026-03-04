import AnalyticsView from '@/views/AnalyticsView.vue'
import CostMonitoringView from '@/views/CostMonitoringView.vue'
import DashboardView from '@/views/DashboardView.vue'
import DeviceListView from '@/views/DeviceListView.vue'
import SettingsView from '@/views/SettingsView.vue'
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      component: DashboardView
    },
    {
      path: "/devices",
      component: DeviceListView
    },
    {
      path: "/cost-monitor",
      component: CostMonitoringView
    },

    {
      path: "/analytics",
      component: AnalyticsView
    },

    {
      path: "/settings",
      component: SettingsView
    },
  ],
})

export default router
