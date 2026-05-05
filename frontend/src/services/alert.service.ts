import api from "./api";

export interface Alert {
  id: number
  message: string
  severity: 'low' | 'medium' | 'high'
  triggered_at: string
}

export const alertService = {
  async getRecent(): Promise<Alert[]> {
    const res = await api.get("/alerts", { params: { limit: 5 } })
    return res.data
  },

  async getAll(): Promise<Alert[]> {
    const res = await api.get("/alerts")
    return res.data
  },

  async getAnalyticsAlerts(params: {
    range: string
    appliance_id?: number
  }) {
    const res = await api.get("/alerts/analytics", { params })
    return res.data
  },

  async getAlertsByAppliance(id: number) {
    const res = await api.get(`/alerts/appliances/${id}`)
    return res.data
  }
}