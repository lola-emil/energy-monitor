<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import {
    readingService,
    type ChartPoint,
    type ReadingSummary,
} from "@/services/reading.service"
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Separator } from '@/components/ui/separator'
import { Badge } from '@/components/ui/badge'
import { Skeleton } from '@/components/ui/skeleton'
import VChart from 'vue-echarts'
import { use } from 'echarts/core'

import {
    CanvasRenderer
} from 'echarts/renderers'

import {
    LineChart
} from 'echarts/charts'

import {
    GridComponent,
    TooltipComponent,
    LegendComponent
} from 'echarts/components'

import { useThemeColors } from '@/composables/useThemeColors'
import { BoltIcon, CircleDollarSignIcon, ActivityIcon, AlertTriangleIcon } from 'lucide-vue-next'

use([
    CanvasRenderer,
    LineChart,
    GridComponent,
    TooltipComponent,
    LegendComponent,
])


const monthlyChartData = ref([
    { label: '1', value: 32 },
    { label: '5', value: 28 },
    { label: '10', value: 41 },
    { label: '15', value: 35 },
    { label: '20', value: 48 },
    { label: '25', value: 39 },
    { label: '30', value: 44 },
])

const dailyChartData = ref([
    { label: '00:00', value: 120 },
    { label: '04:00', value: 90 },
    { label: '08:00', value: 310 },
    { label: '12:00', value: 430 },
    { label: '16:00', value: 390 },
    { label: '20:00', value: 280 },
    { label: '23:00', value: 160 },
])

const { colors } = useThemeColors()

// ---- Mock data (replace with API calls) ----
const isLoading = ref(true)

const now = ref(new Date())
const period = ref<'month' | 'day'>('month')

const isLoadingChart = ref(false)
const chartData = ref<ChartPoint[]>([])

const totalEnergyKwh = ref(1284)       // this month
const ratePerKwh = ref(10.25)         // from settings API
const activeDevices = ref([
    { id: 1, name: 'Main Meter', status: 'online', power: 430 },
    { id: 2, name: 'Office Load', status: 'online', power: 210 },
    { id: 3, name: 'Spare Meter', status: 'offline', power: 0 }
])

// simple sample alerts
const recentAlerts = ref([
    { id: 1, time: '2026-04-23 10:30', message: 'Voltage exceeded 250 V', severity: 'high' },
    { id: 2, time: '2026-04-22 16:05', message: 'Device Office Load went offline', severity: 'medium' }
])

const summary = ref<ReadingSummary>({
    total_energy_kwh: 0,
    estimated_cost: 0,
    peak_power: 0,
    active_devices: 0,
    active_alerts: 0,
})

const activeChartData = computed(() => chartData.value)

const fetchSummary = async () => {
    try {
        isLoading.value = true

        const data = await readingService.getSummary()
        summary.value = data
    } catch (error) {
        console.error("Failed to load summary:", error)
    } finally {
        isLoading.value = false
    }
}

const fetchChart = async () => {
    try {
        isLoadingChart.value = true

        const data = await readingService.getChart()
        chartData.value = data

    } catch (err) {
        console.error("Failed to load chart:", err)
    } finally {
        isLoadingChart.value = false
    }
}

const lastUpdatedText = computed(() =>
    now.value.toLocaleString(undefined, { hour12: false })
)

// const activeChartData = computed(() =>
//     period.value === 'day'
//         ? dailyChartData.value
//         : monthlyChartData.value
// )

const chartSeriesName = computed(() =>
    period.value === 'day'
        ? 'Power Usage Today'
        : 'Monthly Energy Usage'
)

const chartOptions = computed(() => {
    if (!colors.value) return {}

    return {
        tooltip: {
            trigger: 'axis',
            backgroundColor: `hsl(${colors.value.card})`,
            borderColor: `hsl(${colors.value.border})`,
            textStyle: {
                color: `hsl(${colors.value.cardForeground})`,
            },
        },

        grid: {
            left: 10,
            right: 10,
            top: 20,
            bottom: 10,
            containLabel: true,
        },

        xAxis: {
            type: 'category',
            boundaryGap: false,
            data: activeChartData.value.map(item => item.label),
            axisLine: {
                lineStyle: {
                    color: colors.value.border,
                },
            },
            axisLabel: {
                color: colors.value.mutedForeground,
            },
        },

        yAxis: {
            type: 'value',
            axisLine: {
                lineStyle: {
                    color: colors.value.border,
                },
            },
            splitLine: {
                lineStyle: {
                    color: colors.value.border,
                    opacity: 0.3,
                },
            },
            axisLabel: {
                color: colors.value.mutedForeground,
            },
        },
        series: [
            {
                name: chartSeriesName.value,
                type: 'line',
                smooth: true,

                data: activeChartData.value.map(item => item.value),

                lineStyle: {
                    width: 3,
                    color: colors.value.chart1,
                },

                areaStyle: {
                    opacity: 0.18,
                    color: colors.value.chart1,
                },

                symbol: 'circle',
                symbolSize: 8,
                showSymbol: true,

                emphasis: {
                    focus: 'none',
                    scale: false,
                    lineStyle: {
                        width: 3,
                    },
                    areaStyle: {
                        opacity: 0.18,
                    },
                },

                blur: {
                    lineStyle: {
                        opacity: 1,
                    },
                    areaStyle: {
                        opacity: 0.18,
                    },
                },
            },
        ]
    }
})

onMounted(() => {
    fetchSummary()
    fetchChart()
})
</script>

<template>
    <div class="px-5 my-5">
        <div class="flex flex-col gap-6">
            <!-- Header -->
            <div class="flex-col md:flex-row flex md:items-center justify-between">
                <div class="">
                    <h1 class="text-2xl font-semibold tracking-tight">Overview</h1>
                    <p class="text-sm text-muted-foreground">
                        Quick summary of your energy usage and devices.
                    </p>
                </div>

                <div class="flex-col md:flex-row md:flex hidden items-center lg:gap-3 text-xs text-muted-foreground">
                    <div class="flex items-center gap-3">
                        <span>Period:</span>
                        <div class="inline-flex rounded-md border bg-background p-0.5">
                            <Button variant="ghost" size="sm"
                                :class="period === 'day' ? 'bg-primary text-primary-foreground' : ''"
                                @click="period = 'day'">
                                Today
                            </Button>
                            <Button variant="ghost" size="sm"
                                :class="period === 'month' ? 'bg-primary text-primary-foreground' : ''"
                                @click="period = 'month'">
                                This Month
                            </Button>
                        </div>
                    </div>
                    <Separator orientation="vertical" class="h-6" />
                    <span>Last updated: {{ lastUpdatedText }}</span>
                </div>
            </div>

            <!-- KPI Cards -->
            <div class="grid gap-4 md:grid-cols-2 xl:grid-cols-4">
                <!-- Total Energy -->
                <Card>
                    <CardHeader class="flex flex-row items-center justify-between space-y-0 pb-2">
                        <CardTitle class="text-sm font-medium">
                            Total Energy (this month)
                        </CardTitle>
                        <BoltIcon class="h-4 w-4 text-primary" />
                    </CardHeader>
                    <CardContent>
                        <p class="text-2xl font-bold">
                            <Skeleton v-if="isLoading" class="h-7 w-24" />
                            <span v-else>{{ summary.total_energy_kwh.toFixed(2) }} kWh</span>
                        </p>
                        <p class="text-xs text-muted-foreground mt-1">
                            Based on all active devices.
                        </p>
                    </CardContent>
                </Card>

                <!-- Estimated Bill -->
                <Card>
                    <CardHeader class="flex flex-row items-center justify-between space-y-0 pb-2">
                        <CardTitle class="text-sm font-medium">
                            Estimated Bill (this month)
                        </CardTitle>
                        <CircleDollarSignIcon class="h-4 w-4 text-primary" />
                    </CardHeader>
                    <CardContent>
                        <p class="text-2xl font-bold">
                            <Skeleton v-if="isLoading" class="h-7 w-28" />
                            <span v-else>₱{{ summary.estimated_cost.toFixed(2) }}</span>
                        </p>
                        <p class="text-xs text-muted-foreground mt-1">
                            Using rate of ₱{{ ratePerKwh }} per kWh (editable in Settings).
                        </p>
                    </CardContent>
                </Card>

                <!-- Active Devices -->
                <Card>
                    <CardHeader class="flex flex-row items-center justify-between space-y-0 pb-2">
                        <CardTitle class="text-sm font-medium">
                            Active Devices
                        </CardTitle>
                        <ActivityIcon class="h-4 w-4 text-primary" />
                    </CardHeader>
                    <CardContent>
                        <p class="text-2xl font-bold">
                            <Skeleton v-if="isLoading" class="h-7 w-10" />
                            <span v-else>{{ summary.active_devices }}</span>
                        </p>
                        <p class="text-xs text-muted-foreground mt-1">
                            Online out of {{ activeDevices.length }} total devices.
                        </p>
                    </CardContent>
                </Card>

                <!-- Optional: Average Power Today -->
                <Card>
                    <CardHeader class="flex flex-row items-center justify-between space-y-0 pb-2">
                        <CardTitle class="text-sm font-medium">
                            Peak Power
                        </CardTitle>
                        <BoltIcon class="h-4 w-4 text-primary" />
                    </CardHeader>
                    <CardContent>
                        <p class="text-2xl font-bold">
                            <Skeleton v-if="isLoading" class="h-7 w-20" />
                            <span v-else>{{ summary.peak_power.toFixed(2) }} W</span>
                        </p>
                        <p class="text-xs text-muted-foreground mt-1">
                            Simple mean of readings collected today.
                        </p>
                    </CardContent>
                </Card>
            </div>

            <!-- Main content: chart + active devices -->
            <div class="grid gap-4 lg:grid-cols-3">
                <!-- Usage chart placeholder -->
                <Card class="lg:col-span-2">
                    <CardHeader>
                        <CardTitle>
                            {{ period === 'day' ? 'Today\'s Energy Usage' : 'Monthly Energy Usage' }}
                        </CardTitle>
                        <CardDescription>
                            Visual summary of your consumption {{ period === 'day' ? 'over the last 24 hours.' :
                                'this month.' }}
                        </CardDescription>
                    </CardHeader>
                    <CardContent>
                        <!-- Replace this div with your actual chart component -->
                        <div class="h-80">
                            <VChart class="h-72 w-full" :option="chartOptions" autoresize />

                        </div>
                    </CardContent>
                </Card>

                <!-- Active devices list -->
                <Card>
                    <CardHeader>
                        <CardTitle>Active Devices</CardTitle>
                        <CardDescription>
                            Quick view of your devices and current power draw.
                        </CardDescription>
                    </CardHeader>
                    <CardContent class="space-y-3">
                        <template v-if="!activeDevices.length">
                            <p class="text-sm text-muted-foreground">No devices registered yet.</p>
                        </template>
                        <template v-else>
                            <div v-for="device in activeDevices" :key="device.id"
                                class="flex items-center justify-between rounded-md border px-3 py-2">
                                <div>
                                    <p class="text-sm font-medium leading-none">
                                        {{ device.name }}
                                    </p>
                                    <p class="text-xs text-muted-foreground">
                                        {{ device.status === 'online' ? 'Online' : 'Offline' }}
                                    </p>
                                </div>
                                <div class="flex items-center gap-2">
                                    <Badge :variant="device.status === 'online' ? 'default' : 'outline'"
                                        class="text-xs">
                                        {{ device.status === 'online' ? 'Online' : 'Offline' }}
                                    </Badge>
                                    <span class="text-sm font-medium" v-if="device.status === 'online'">
                                        {{ device.power }} W
                                    </span>
                                </div>
                            </div>
                        </template>
                    </CardContent>
                </Card>
            </div>

            <!-- Alerts preview -->
            <Card>
                <CardHeader class="flex flex-row items-center justify-between">
                    <div>
                        <CardTitle>Recent Alerts</CardTitle>
                        <CardDescription>
                            Last few anomalies detected by the system.
                        </CardDescription>
                    </div>
                    <Button variant="ghost" size="sm" to="/alerts">
                        View all
                    </Button>
                </CardHeader>
                <CardContent>
                    <div v-if="!recentAlerts.length"
                        class="flex flex-col items-center justify-center py-6 text-sm text-muted-foreground">
                        <AlertTriangleIcon class="mb-2 h-5 w-5" />
                        No alerts yet. Your system is stable.
                    </div>
                    <div v-else class="space-y-2">
                        <div v-for="alert in recentAlerts" :key="alert.id"
                            class="flex items-start justify-between rounded-md border px-3 py-2 text-sm">
                            <div>
                                <p class="font-medium">{{ alert.message }}</p>
                                <p class="text-xs text-muted-foreground">{{ alert.time }}</p>
                            </div>
                            <Badge :variant="alert.severity === 'high' ? 'destructive' : 'outline'"
                                class="uppercase text-[10px]">
                                {{ alert.severity }}
                            </Badge>
                        </div>
                    </div>
                </CardContent>
            </Card>
        </div>
    </div>
</template>