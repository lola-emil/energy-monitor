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
    async getSummary(params: { range: string, month?: string, year?: string; }): Promise<ReadingSummary> {
        const response = await api.get("/readings/summary", { params });
        return response.data;
    },

    async getChart(params: { range: string, month?: string, year?: string; }): Promise<ChartPoint[]> {
        const res = await api.get("/readings/chart", {
            params
        });
        return res.data;
    },

    async getAnalytics(params: {
        range: string;
        appliance_id?: number;
        month?: string,
        year?: string,
    }): Promise<AnalyticsResponse> {
        const res = await api.get("/readings/analytics", {
            params: {
                ...params,
            },
        });
        return res.data;
    },

    async getDetailedReadings(params: {
        range: string,
        appliance_id?: number;
        page?: number,
        page_size?: number,

        month?: string,
        year?: string,
    }) {
        const res = await api.get("/readings/detailed", { params });
        return res.data;
    }
};
