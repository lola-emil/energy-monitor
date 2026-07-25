<template>
    <div class="min-h-screen bg-base-200 flex flex-col">

        <Navbar />

        <main class="flex-1 p-4 md:p-8">

            <div v-if="currentPage === 'analytics'">

                <div class="flex flex-col lg:flex-row gap-6">

                    <div class="w-full lg:w-1/4 grid grid-rows-4 gap-4">

                        <div class="card bg-primary text-primary-content shadow-lg border-0">
                            <div class="card-body p-5">
                                <div class="flex items-center gap-2 opacity-90">
                                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none"
                                        viewBox="0 0 24 24" stroke="currentColor">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                            d="M13 10V3L4 14h7v7l9-11h-7z" />
                                    </svg>
                                    <span class="font-medium">Total Consumption</span>
                                </div>
                                <div class="flex justify-between items-end mt-3">
                                    <h2 class="text-3xl font-bold">{{ summary?.total_energy_kwh }} <span
                                            class="text-lg font-normal opacity-70">kWh</span></h2>
                                    <span class="badge bg-white/20 text-white border-0 gap-1">
                                        <svg xmlns="http://www.w3.org/2000/svg" class="h-3 w-3" fill="none"
                                            viewBox="0 0 24 24" stroke="currentColor">
                                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                                d="M5 10l7-7m0 0l7 7m-7-7v18" />
                                        </svg>
                                        5.2%
                                    </span>
                                </div>
                            </div>
                        </div>

                        <div class="card bg-base-100 shadow-sm border border-base-300">
                            <div class="card-body p-5">
                                <div class="flex items-center gap-2 text-base-content/70">
                                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none"
                                        viewBox="0 0 24 24" stroke="currentColor">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                            d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                                    </svg>
                                    <span class="font-medium">Estimated Bill</span>
                                </div>
                                <div class="flex justify-between items-end mt-3">
                                    <h2 class="text-3xl font-bold text-base-content">$ {{ summary?.estimated_cost }}
                                    </h2>
                                    <span class="badge bg-error/10 text-error border-0 gap-1">
                                        <svg xmlns="http://www.w3.org/2000/svg" class="h-3 w-3" fill="none"
                                            viewBox="0 0 24 24" stroke="currentColor">
                                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                                d="M19 14l-7 7m0 0l-7-7m7 7V3" />
                                        </svg>
                                        2.1%
                                    </span>
                                </div>
                            </div>
                        </div>

                        <div class="card bg-base-100 shadow-sm border border-base-300">
                            <div class="card-body p-5">
                                <div class="flex items-center gap-2 text-base-content/70">
                                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none"
                                        viewBox="0 0 24 24" stroke="currentColor">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                            d="M9 3v2m6-2v2M9 19v2m6-2v2M5 9H3m2 6H3m18-6h-2m2 6h-2M7 19h10a2 2 0 002-2V7a2 2 0 00-2-2H7a2 2 0 00-2 2v10a2 2 0 002 2zM9 9h6v6H9V9z" />
                                    </svg>
                                    <span class="font-medium">Active Devices</span>
                                </div>
                                <div class="flex justify-between items-end mt-3">
                                    <h2 class="text-3xl font-bold text-base-content">{{ summary?.active_devices }} <span
                                            class="text-lg font-normal text-base-content/50">/ {{ summary?.device_count
                                            }}</span></h2>
                                    <span class="badge bg-success/10 text-success border-0 gap-1">
                                        <span class="w-1.5 h-1.5 rounded-full bg-success"></span>
                                        Nominal
                                    </span>
                                </div>
                            </div>
                        </div>

                        <div class="card bg-base-100 shadow-sm border border-base-300">
                            <div class="card-body p-5">
                                <div class="flex items-center gap-2 text-base-content/70">
                                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none"
                                        viewBox="0 0 24 24" stroke="currentColor">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                            d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z" />
                                    </svg>
                                    <span class="font-medium">Monthly Budget</span>
                                </div>
                                <div class="flex justify-between items-end mt-3">
                                    <h2 class="text-3xl font-bold text-base-content">${{ budgetUsed }} <span
                                            class="text-lg font-normal text-base-content/50">/ ${{ budgetTotal }}</span>
                                    </h2>
                                    <span class="badge bg-warning/10 text-warning border-0 gap-1">
                                        {{ budgetPercentage }}%
                                    </span>
                                </div>
                                <progress class="progress progress-warning w-full mt-2 h-1.5" :value="budgetPercentage"
                                    max="100"></progress>
                            </div>
                        </div>

                    </div>

                    <div class="w-full lg:w-3/4 flex flex-col gap-6">

                        <div class="card bg-base-100 shadow-sm border border-base-300 flex-1">
                            <div class="card-body p-6">

                                <div
                                    class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4 mb-6">
                                    <h2 class="text-xl font-bold text-base-content">Energy Metrics</h2>
                                    <div class="flex flex-wrap items-center gap-2">
                                        <div class="join border border-base-300 rounded-lg bg-base-200">
                                            <input type="date" v-model="startDate"
                                                class="input input-sm join-item border-none bg-transparent w-32" />
                                            <span
                                                class="join-item flex items-center px-2 text-base-content/50 text-xs">to</span>
                                            <input type="date" v-model="endDate"
                                                class="input input-sm join-item border-none bg-transparent w-32" />
                                        </div>
                                        <button @click="refreshData" class="btn btn-primary btn-sm gap-2">
                                            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none"
                                                viewBox="0 0 24 24" stroke="currentColor">
                                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                                    d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                                            </svg>
                                            Refresh
                                        </button>
                                    </div>
                                </div>

                                <div
                                    class="w-full h-96 bg-base-200/50 rounded-xl flex items-center justify-center border border-dashed border-base-300 relative">
                                    <EnergyTrendChart />
                                </div>

                            </div>
                        </div>

                        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">

                            <div class="card bg-base-100 shadow-sm border border-base-300">
                                <div class="card-body p-5">
                                    <h3 class="card-title text-base text-base-content">Top Energy Hogs</h3>
                                    <ul class="mt-2 space-y-4">
                                        <li v-for="device in topDevices" :key="device.name"
                                            class="flex items-center justify-between">
                                            <div class="flex items-center gap-3">
                                                <div class="p-2 rounded-lg" :class="device.bgColor">
                                                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4"
                                                        :class="device.iconColor" fill="none" viewBox="0 0 24 24"
                                                        stroke="currentColor">
                                                        <path stroke-linecap="round" stroke-linejoin="round"
                                                            stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
                                                    </svg>
                                                </div>
                                                <div>
                                                    <p class="font-semibold text-sm text-base-content">{{ device.name }}
                                                    </p>
                                                    <p class="text-xs text-base-content/60">{{ device.usageTime }}</p>
                                                </div>
                                            </div>
                                            <div class="text-right">
                                                <p class="font-bold text-sm text-base-content">{{ device.kWh }} kWh</p>
                                                <p class="text-xs font-medium" :class="device.textColor">{{
                                                    device.percentage }}%</p>
                                            </div>
                                        </li>
                                    </ul>
                                </div>
                            </div>

                            <div class="card bg-base-100 shadow-sm border border-base-300">
                                <div class="card-body p-5">
                                    <h3 class="card-title text-base text-base-content">Peak vs Off-Peak</h3>
                                    <div class="grid grid-cols-2 gap-3 mt-2">
                                        <div class="bg-error/5 border border-error/20 rounded-lg p-3 text-center">
                                            <p class="text-xs text-error font-semibold uppercase tracking-wide">Peak</p>
                                            <p class="text-2xl font-bold text-base-content mt-1">{{ peakUsage }}</p>
                                            <p class="text-xs text-base-content/60">kWh ({{ peakPercentage }}%)</p>
                                        </div>
                                        <div class="bg-success/5 border border-success/20 rounded-lg p-3 text-center">
                                            <p class="text-xs text-success font-semibold uppercase tracking-wide">
                                                Off-Peak</p>
                                            <p class="text-2xl font-bold text-base-content mt-1">{{ offPeakUsage }}</p>
                                            <p class="text-xs text-base-content/60">kWh ({{ offPeakPercentage }}%)</p>
                                        </div>
                                    </div>
                                    <div class="mt-4 p-3 bg-base-200 rounded-lg flex items-start gap-2">
                                        <svg xmlns="http://www.w3.org/2000/svg"
                                            class="h-4 w-4 text-warning flex-shrink-0 mt-0.5" fill="none"
                                            viewBox="0 0 24 24" stroke="currentColor">
                                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                                d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z" />
                                        </svg>
                                        <p class="text-xs text-base-content/80 leading-relaxed">Shift 50kWh of laundry
                                            to off-peak hours to save <span class="font-bold text-success">$8.00</span>.
                                        </p>
                                    </div>
                                </div>
                            </div>

                        </div>
                    </div>
                </div>
            </div>

            <div v-else-if="currentPage === 'devices'" class="max-w-4xl mx-auto">
                <div class="card bg-base-100 shadow-sm border border-base-300">
                    <div class="card-body">
                        <h2 class="text-2xl font-bold mb-4">Device Management</h2>
                        <p class="text-base-content/70">Manage your smart meters, sensors, and connected appliances.</p>
                        <div
                            class="mt-8 p-12 bg-base-200/50 rounded-xl border border-dashed border-base-300 text-center">
                            <p class="text-base-content/40 font-medium">[ Device List & Controls Go Here ]</p>
                        </div>
                    </div>
                </div>
            </div>

            <div v-else-if="currentPage === 'settings'" class="max-w-4xl mx-auto">
                <div class="card bg-base-100 shadow-sm border border-base-300">
                    <div class="card-body">
                        <h2 class="text-2xl font-bold mb-4">Settings</h2>
                        <p class="text-base-content/70">Configure billing thresholds, notification preferences, and API
                            integrations.</p>
                        <div
                            class="mt-8 p-12 bg-base-200/50 rounded-xl border border-dashed border-base-300 text-center">
                            <p class="text-base-content/40 font-medium">[ Settings Form Goes Here ]</p>
                        </div>
                    </div>
                </div>
            </div>

        </main>
    </div>
</template>

<script setup lang="ts">
import EnergyTrendChart from '@/components/EnergyTrendChart.vue';
import Navbar from '@/components/Navbar.vue';
import { readingService, type ReadingSummary } from '@/services/reading.service';
import { ref, computed, onMounted } from 'vue';

const currentPage = ref('analytics');

const startDate = ref('2026-07-01');
const endDate = ref('2026-07-25');

const refreshData = () => {
    console.log(`Refreshing data from ${startDate.value} to ${endDate.value}`);
};
const budgetUsed = ref(140);
const budgetTotal = ref(200);
const budgetPercentage = computed(() => Math.round((budgetUsed.value / budgetTotal.value) * 100));

const topDevices = ref([
    { name: 'Central AC', usageTime: 'Running 14h/day', kWh: 342, percentage: 45, bgColor: 'bg-error/10', iconColor: 'text-error', textColor: 'text-error' },
    { name: 'EV Charger', usageTime: 'Charging 6h/day', kWh: 185, percentage: 24, bgColor: 'bg-warning/10', iconColor: 'text-warning', textColor: 'text-warning' },
    { name: 'Refrigerator', usageTime: 'Running 24/7', kWh: 95, percentage: 12, bgColor: 'bg-info/10', iconColor: 'text-info', textColor: 'text-info' }
]);

const peakUsage = ref(310);
const offPeakUsage = ref(465);
const totalUsage = computed(() => peakUsage.value + offPeakUsage.value);
const peakPercentage = computed(() => Math.round((peakUsage.value / totalUsage.value) * 100));
const offPeakPercentage = computed(() => 100 - peakPercentage.value);

const summary = ref<ReadingSummary>();

onMounted(async () => {
    const data = await readingService.getSummary({ range: "" })
    summary.value = data;
})
</script>