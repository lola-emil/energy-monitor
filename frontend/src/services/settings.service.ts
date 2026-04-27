import api from "./api"
import type { Settings } from "@/types/settings"

export const settingsService = {
  async get(): Promise<Settings> {
    const response = await api.get("/settings")
    return response.data
  },

  async update(payload: Settings) {
    const response = await api.put("/settings", payload)
    return response.data
  },
}