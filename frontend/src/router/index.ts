import WSTestView from '@/views/ WSTestView.vue'
import DashboardView from '@/views/DashboardView.vue'
import DeviceListView from '@/views/DeviceListView.vue'
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
      path: "/ws-test",
      component: WSTestView
    }
  ],
})

export default router
