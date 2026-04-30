import api from "./api"

export interface ReadingSummary {
    total_energy_kwh: number
    estimated_cost: number
    peak_power: number
    active_devices: number
    active_alerts: number
}

export interface ChartPoint {
    label: string
    value: number
}


export const readingService = {
    async getSummary(): Promise<ReadingSummary> {
        const response = await api.get("/readings/summary")
        return response.data
    },


    async getChart(): Promise<ChartPoint[]> {
        const res = await api.get("/readings/chart")
        return res.data
    }
}