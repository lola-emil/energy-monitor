import api from "./api";

export interface ReadingSummary {
    total_energy_kwh: number;
    estimated_cost: number;
    peak_power: number;
    active_devices: number;
    device_count: number;
    billing_rate: number;
    active_alerts: number;
}

export interface ChartPoint {
    label: string;
    value: number;
}


export interface AnalyticsResponse {
    summary: {
        total_energy_kwh: number;
        avg_power: number;
        avg_voltage: number;
        avg_current: number;
        peak_power: number;
    };
    energy: { label: string; value: number; }[];
    voltage_current: {
        label: string;
        voltage: number;
        current: number;
    }[];
}

export const readingService = {
    async getSummary(): Promise<ReadingSummary> {
        const response = await api.get("/readings/summary");
        return response.data;
    },


    async getChart(): Promise<ChartPoint[]> {
        const res = await api.get("/readings/chart");
        return res.data;
    },

    async getAnalytics(params: {
        range: string;
        appliance_id?: number;
    }): Promise<AnalyticsResponse> {
        const res = await api.get("/readings/analytics", { params });
        return res.data;
    },
};
