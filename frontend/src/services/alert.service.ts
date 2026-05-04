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
}