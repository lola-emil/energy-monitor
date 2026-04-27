<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'

// shadcn-vue components (update paths to your project)
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Select, SelectTrigger, SelectContent, SelectItem, SelectValue } from '@/components/ui/select'
import { Separator } from '@/components/ui/separator'
import { Skeleton } from '@/components/ui/skeleton'
import { Badge } from '@/components/ui/badge'
import {
    BoltIcon,
    GaugeIcon,
    ActivityIcon,
    CalendarRangeIcon,
    AlertTriangleIcon
} from 'lucide-vue-next'
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

use([
    CanvasRenderer,
    LineChart,
    GridComponent,
    TooltipComponent,
    LegendComponent,
])

const { colors } = useThemeColors()

const activeEnergySeries = computed(() => {
    if (selectedRange.value === 'today') {
        return [
            { label: '00:00', value: 120 },
            { label: '04:00', value: 90 },
            { label: '08:00', value: 310 },
            { label: '12:00', value: 430 },
            { label: '16:00', value: 390 },
            { label: '20:00', value: 280 },
        ]
    }

    if (selectedRange.value === '7d') {
        return [
            { label: 'Mon', value: 410 },
            { label: 'Tue', value: 390 },
            { label: 'Wed', value: 450 },
            { label: 'Thu', value: 430 },
            { label: 'Fri', value: 470 },
            { label: 'Sat', value: 520 },
            { label: 'Sun', value: 480 },
        ]
    }

    return energySeries.value
})

// ------------ State (mock data for now) ------------
const isLoading = ref(true)

const devices = ref([
    { id: 'all', name: 'All devices' },
    { id: '1', name: 'Main Meter' },
    { id: '2', name: 'Office Load' }
])

const selectedDeviceId = ref('all')
const selectedRange = ref<'today' | '7d' | 'month'>('month')
const lastUpdated = ref(new Date())

// Summary metrics for selected device + range
const totalEnergy = ref(1284)   // kWh
const avgPower = ref(480)       // W
const avgVoltage = ref(229)     // V
const avgCurrent = ref(2.1)     // A
const peakPower = ref(910)      // W

// Simple sample timeseries for charts
const energySeries = ref([
    { label: 'Jan', value: 600 },
    { label: 'Feb', value: 580 },
    { label: 'Mar', value: 620 },
    { label: 'Apr', value: 590 },
    { label: 'May', value: 610 },
    { label: 'Jun', value: 570 },
    { label: 'Jul', value: 600 },
    { label: 'Aug', value: 640 },
    { label: 'Sep', value: 700 },
    { label: 'Oct', value: 780 },
    { label: 'Nov', value: 820 },
    { label: 'Dec', value: 960 }
])

const voltageCurrentSeries = ref([
    { time: '00:00', voltage: 228, current: 1.8 },
    { time: '04:00', voltage: 230, current: 1.6 },
    { time: '08:00', voltage: 231, current: 2.4 },
    { time: '12:00', voltage: 229, current: 2.9 },
    { time: '16:00', voltage: 227, current: 2.3 },
    { time: '20:00', voltage: 230, current: 2.0 }
])

// Table rows
const readings = ref([
    { time: '2026-04-23 10:00', voltage: 229, current: 2.1, power: 482, energy: 0.48 },
    { time: '2026-04-23 10:05', voltage: 230, current: 2.0, power: 460, energy: 0.46 },
    { time: '2026-04-23 10:10', voltage: 231, current: 2.3, power: 531, energy: 0.53 }
])

// Alerts in selected period
const alerts = ref([
    { id: 1, time: '2026-04-22 16:05', device: 'Office Load', message: 'Device went offline', severity: 'medium' },
    { id: 2, time: '2026-04-21 09:12', device: 'Main Meter', message: 'Voltage exceeded 250 V', severity: 'high' }
])

const lastUpdatedText = computed(() =>
    lastUpdated.value.toLocaleString(undefined, { hour12: false })
)

const energyChartOptions = computed(() => {
    if (!colors.value) return {}

    return {
        tooltip: {
            trigger: 'axis',
            backgroundColor: colors.value.card,
            borderColor: colors.value.border,
            textStyle: {
                color: colors.value.cardForeground,
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
            data: activeEnergySeries.value.map(i => i.label),
            axisLabel: {
                color: colors.value.mutedForeground,
            },
            axisLine: {
                lineStyle: {
                    color: colors.value.border,
                },
            },
        },

        yAxis: {
            type: 'value',
            axisLabel: {
                color: colors.value.mutedForeground,
            },
            splitLine: {
                lineStyle: {
                    color: colors.value.border,
                    opacity: 0.3,
                },
            },
        },

        series: [
            {
                name: 'Energy Usage',
                type: 'line',
                smooth: true,
                showSymbol: true,
                symbol: 'circle',
                symbolSize: 8,

                data: activeEnergySeries.value.map(i => i.value),

                lineStyle: {
                    width: 3,
                    color: colors.value.chart1,
                },

                areaStyle: {
                    opacity: 0.18,
                    color: colors.value.chart1,
                },

                emphasis: {
                    focus: 'none',
                    scale: false,
                },
            },
        ],
    }
})


const voltageChartOptions = computed(() => {
    if (!colors.value) return {}

    return {
        tooltip: {
            trigger: 'axis',
            backgroundColor: colors.value.card,
            borderColor: colors.value.border,
        },

        // legend: {
        //     textStyle: {
        //         color: colors.value.mutedForeground,
        //     },
        // },

        grid: {
            left: 20,
            right: 20,
            top: 50,
            bottom: 20,
            containLabel: true,
        },

        xAxis: {
            type: 'category',
            data: voltageCurrentSeries.value.map(i => i.time),
            axisLabel: {
                color: colors.value.mutedForeground,
            },
        },

        yAxis: [
            {
                type: 'value',
                name: 'Voltage',
                min: 180,
                max: 260,
                interval: 20,

                axisLabel: {
                    color: colors.value.mutedForeground,
                },

                nameTextStyle: {
                    color: colors.value.mutedForeground,
                },

                splitLine: {
                    show: true,
                    lineStyle: {
                        color: colors.value.border,
                        opacity: 0.3,
                    },
                },

                axisLine: {
                    show: false,
                },

                axisTick: {
                    show: false,
                },
            },

            {
                type: 'value',
                name: 'Current',
                min: 0,
                max: 3.5,
                interval: 0.5,

                axisLabel: {
                    color: colors.value.mutedForeground,
                },

                nameTextStyle: {
                    color: colors.value.mutedForeground,
                },

                splitLine: {
                    show: false, // 🔥 important fix
                },

                axisLine: {
                    show: false,
                },

                axisTick: {
                    show: false,
                },
            },
        ],

        series: [
            {
                name: 'Voltage',
                type: 'line',
                smooth: true,
                data: voltageCurrentSeries.value.map(i => i.voltage),
                lineStyle: {
                    width: 3,
                    color: colors.value.chart1,
                },
            },
            {
                name: 'Current',
                type: 'line',
                smooth: true,
                yAxisIndex: 1,
                data: voltageCurrentSeries.value.map(i => i.current),
                lineStyle: {
                    width: 3,
                    color: colors.value.chart2,
                },
            },
        ],
    }
})

onMounted(() => {
    // simulate loading
    setTimeout(() => {
        isLoading.value = false
    }, 700)
})

// TODO: add watchers on selectedDeviceId / selectedRange to refetch real data from your API.
</script>

<template>
    <div class="px-5 mt-5">
        <div class="flex flex-col gap-6">
            <!-- Header / filters -->
            <div class="flex flex-col gap-4 md:flex-row md:items-center md:justify-between">
                <div>
                    <h1 class="text-2xl font-semibold tracking-tight">Analytics</h1>
                    <p class="text-sm text-muted-foreground">
                        Detailed historical and real-time energy data, with charts and tables.
                    </p>
                </div>

                <div class="flex flex-wrap items-center gap-3 text-xs text-muted-foreground">
                    <!-- Device selector -->
                    <div class="flex items-center gap-2">
                        <span class="whitespace-nowrap">Device:</span>
                        <Select v-model="selectedDeviceId">
                            <SelectTrigger class="w-45">
                                <SelectValue placeholder="Select device" />
                            </SelectTrigger>
                            <SelectContent>
                                <SelectItem v-for="dev in devices" :key="dev.id" :value="dev.id">
                                    {{ dev.name }}
                                </SelectItem>
                            </SelectContent>
                        </Select>
                    </div>

                    <!-- Range selector -->
                    <div class="flex items-center gap-2">
                        <span class="whitespace-nowrap">Range:</span>
                        <div class="inline-flex rounded-md border bg-background p-0.5">
                            <Button size="sm" variant="ghost"
                                :class="selectedRange === 'today' ? 'bg-primary text-primary-foreground' : ''"
                                @click="selectedRange = 'today'">
                                Today
                            </Button>
                            <Button size="sm" variant="ghost"
                                :class="selectedRange === '7d' ? 'bg-primary text-primary-foreground' : ''"
                                @click="selectedRange = '7d'">
                                Last 7 days
                            </Button>
                            <Button size="sm" variant="ghost"
                                :class="selectedRange === 'month' ? 'bg-primary text-primary-foreground' : ''"
                                @click="selectedRange = 'month'">
                                This month
                            </Button>
                        </div>
                    </div>

                    <Separator orientation="vertical" class="hidden h-6 md:block" />
                    <div class="flex items-center gap-1">
                        <CalendarRangeIcon class="h-3 w-3" />
                        <span>Last updated: {{ lastUpdatedText }}</span>
                    </div>
                </div>
            </div>

            <!-- Summary cards -->
            <div class="grid gap-4 md:grid-cols-2 xl:grid-cols-5">
                <!-- Energy -->
                <Card>
                    <CardHeader class="flex flex-row items-center justify-between space-y-0 pb-2">
                        <CardTitle class="text-sm font-medium">Energy used</CardTitle>
                        <BoltIcon class="h-4 w-4 text-primary" />
                    </CardHeader>
                    <CardContent>
                        <p class="text-2xl font-bold">
                            <Skeleton v-if="isLoading" class="h-7 w-24" />
                            <span v-else>{{ totalEnergy }} kWh</span>
                        </p>
                        <p class="mt-1 text-xs text-muted-foreground">
                            Total in selected period.
                        </p>
                    </CardContent>
                </Card>

                <!-- Avg Power -->
                <Card>
                    <CardHeader class="flex flex-row items-center justify-between space-y-0 pb-2">
                        <CardTitle class="text-sm font-medium">Average power</CardTitle>
                        <GaugeIcon class="h-4 w-4 text-primary" />
                    </CardHeader>
                    <CardContent>
                        <p class="text-2xl font-bold">
                            <Skeleton v-if="isLoading" class="h-7 w-24" />
                            <span v-else>{{ avgPower }} W</span>
                        </p>
                        <p class="mt-1 text-xs text-muted-foreground">
                            From all readings in this range.
                        </p>
                    </CardContent>
                </Card>

                <!-- Avg Voltage -->
                <Card>
                    <CardHeader class="flex flex-row items-center justify-between space-y-0 pb-2">
                        <CardTitle class="text-sm font-medium">Average voltage</CardTitle>
                        <ActivityIcon class="h-4 w-4 text-primary" />
                    </CardHeader>
                    <CardContent>
                        <p class="text-2xl font-bold">
                            <Skeleton v-if="isLoading" class="h-7 w-20" />
                            <span v-else>{{ avgVoltage }} V</span>
                        </p>
                        <p class="mt-1 text-xs text-muted-foreground">
                            Indicates supply stability.
                        </p>
                    </CardContent>
                </Card>

                <!-- Avg Current -->
                <Card>
                    <CardHeader class="flex flex-row items-center justify-between space-y-0 pb-2">
                        <CardTitle class="text-sm font-medium">Average current</CardTitle>
                        <ActivityIcon class="h-4 w-4 text-primary" />
                    </CardHeader>
                    <CardContent>
                        <p class="text-2xl font-bold">
                            <Skeleton v-if="isLoading" class="h-7 w-20" />
                            <span v-else>{{ avgCurrent }} A</span>
                        </p>
                        <p class="mt-1 text-xs text-muted-foreground">
                            Load drawn over time.
                        </p>
                    </CardContent>
                </Card>

                <!-- Peak power -->
                <Card>
                    <CardHeader class="flex flex-row items-center justify-between space-y-0 pb-2">
                        <CardTitle class="text-sm font-medium">Peak power</CardTitle>
                        <BoltIcon class="h-4 w-4 text-primary" />
                    </CardHeader>
                    <CardContent>
                        <p class="text-2xl font-bold">
                            <Skeleton v-if="isLoading" class="h-7 w-24" />
                            <span v-else>{{ peakPower }} W</span>
                        </p>
                        <p class="mt-1 text-xs text-muted-foreground">
                            Highest recorded in this range.
                        </p>
                    </CardContent>
                </Card>
            </div>

            <!-- Charts row -->
            <div class="grid gap-4 lg:grid-cols-2">
                <!-- Energy usage chart -->
                <Card>
                    <CardHeader>
                        <CardTitle>Energy usage over time</CardTitle>
                        <CardDescription>
                            kWh vs time for the selected device and period.
                        </CardDescription>
                    </CardHeader>
                    <CardContent>
                        <!-- Replace with real chart component (Chart.js, ECharts, etc.) -->
                        <div
                            class="h-64 rounded-md border border-dashed border-muted text-xs text-muted-foreground">
                            <VChart class="h-72 w-full" :option="energyChartOptions" autoresize />
                        </div>
                    </CardContent>
                </Card>

                <!-- Voltage & current trends chart -->
                <Card>
                    <CardHeader>
                        <CardTitle>Voltage & current trends</CardTitle>
                        <CardDescription>
                            Helps analyze supply quality and load profile.
                        </CardDescription>
                    </CardHeader>
                    <CardContent>
                        <!-- Replace with dual‑axis chart -->
                        <div
                            class="h-64 rounded-md border border-dashed border-muted text-xs text-muted-foreground">
                            <VChart class="h-72 w-full" :option="voltageChartOptions" autoresize />
                        </div>
                    </CardContent>
                </Card>
            </div>

            <!-- Table + alerts -->
            <div class="grid gap-4 lg:grid-cols-3">
                <!-- Historical data table -->
                <Card class="lg:col-span-2">
                    <CardHeader class="flex flex-row items-center justify-between">
                        <div>
                            <CardTitle>Detailed readings</CardTitle>
                            <CardDescription>
                                Tabular view of measurements in the selected range.
                            </CardDescription>
                        </div>
                        <Button variant="outline" size="sm">
                            Export CSV
                        </Button>
                    </CardHeader>
                    <CardContent>
                        <div class="overflow-x-auto">
                            <table class="w-full text-left text-xs">
                                <thead class="border-b text-[11px] uppercase text-muted-foreground">
                                    <tr>
                                        <th class="py-2 pr-3">Timestamp</th>
                                        <th class="py-2 px-3">Voltage (V)</th>
                                        <th class="py-2 px-3">Current (A)</th>
                                        <th class="py-2 px-3">Power (W)</th>
                                        <th class="py-2 px-3">Energy (kWh)</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    <tr v-for="row in readings" :key="row.time" class="border-b last:border-0">
                                        <td class="py-2 pr-3 align-middle text-[11px]">
                                            {{ row.time }}
                                        </td>
                                        <td class="py-2 px-3 align-middle">{{ row.voltage }}</td>
                                        <td class="py-2 px-3 align-middle">{{ row.current }}</td>
                                        <td class="py-2 px-3 align-middle">{{ row.power }}</td>
                                        <td class="py-2 px-3 align-middle">{{ row.energy.toFixed(2) }}</td>
                                    </tr>
                                </tbody>
                            </table>
                        </div>
                    </CardContent>
                </Card>

                <!-- Alerts summary -->
                <Card>
                    <CardHeader class="flex flex-row items-center justify-between">
                        <div>
                            <CardTitle>Alerts in this period</CardTitle>
                            <CardDescription>
                                Anomalies detected for the selected filters.
                            </CardDescription>
                        </div>
                        <Badge variant="outline">
                            {{ alerts.length }} total
                        </Badge>
                    </CardHeader>
                    <CardContent>
                        <div v-if="!alerts.length"
                            class="flex flex-col items-center justify-center py-6 text-sm text-muted-foreground">
                            <AlertTriangleIcon class="mb-2 h-5 w-5" />
                            No alerts in the selected range.
                        </div>
                        <div v-else class="space-y-2 text-xs">
                            <div v-for="alert in alerts" :key="alert.id"
                                class="flex items-start justify-between rounded-md border px-3 py-2">
                                <div>
                                    <p class="font-medium text-sm">{{ alert.message }}</p>
                                    <p class="text-[11px] text-muted-foreground">
                                        {{ alert.time }} · {{ alert.device }}
                                    </p>
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

        <br>
        <br>
    </div>
</template>