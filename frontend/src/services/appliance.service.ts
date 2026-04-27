import api from "./api"
import type { Appliance } from "@/types/appliance"

export const applianceService = {
  async getAll(): Promise<Appliance[]> {
    const response = await api.get("/appliances")
    return response.data
  },

  async getById(id: number): Promise<Appliance> {
    const response = await api.get(`/appliances/${id}`)
    return response.data
  },

  async create(payload: Partial<Appliance>) {
    const response = await api.post("/appliances", payload)
    return response.data
  },

  async update(id: number, payload: Partial<Appliance>) {
    const response = await api.put(`/appliances/${id}`, payload)
    return response.data
  },

  async delete(id: number) {
    await api.delete(`/appliances/${id}`)
  },
}