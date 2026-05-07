<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import {
    readingService,
    type ChartPoint,
    type ReadingSummary,
} from '@/services/reading.service'
import {
    applianceService,
    type ApplianceWithReading,
} from '@/services/appliance.service'
import { alertService, type Alert } from '@/services/alert.service'

import { Card, CardHeader, CardTitle, CardDescription, CardContent } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Separator } from '@/components/ui/separator'
import { Badge } from '@/components/ui/badge'
import { Skeleton } from '@/components/ui/skeleton'
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogDescription } from '@/components/ui/dialog'
import AlertTable from '@/components/AlertTable.vue'
import VChart from 'vue-echarts'

import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { LineChart, BarChart } from 'echarts/charts'
import { GridComponent, TooltipComponent, LegendComponent } from 'echarts/components'

import { useThemeColors } from '@/composables/useThemeColors'
import { formatTime } from '@/lib/time'
import {
    BoltIcon,
    CircleDollarSignIcon,
    ActivityIcon,
    AlertTriangleIcon,
    RefreshCcwIcon,
    ZapIcon,
    BarChart3Icon,
    ShieldAlertIcon,
    CpuIcon,
} from 'lucide-vue-next'

use([
    CanvasRenderer,
    LineChart,
    BarChart,
    GridComponent,
    TooltipComponent,
    LegendComponent,
])

type Period = 'day' | 'month'

const { colors } = useThemeColors()

const selectedPeriod = ref<Period>('month')
const now = ref(new Date())

const isLoadingSummary = ref(true)
const isLoadingChart = ref(true)
const isLoadingDevices = ref(true)
const isLoadingAlerts = ref(false)
const isLoadingRecentAlerts = ref(true)
const isRefreshing = ref(false)

const summary = ref<ReadingSummary>({
    total_energy_kwh: 0,
    estimated_cost: 0,
    peak_power: 0,
    active_devices: 0,
    device_count: 0,
    billing_rate: 0,
    active_alerts: 0,
})

const chartData = ref<ChartPoint[]>([])
const activeDevices = ref<ApplianceWithReading[]>([])
const recentAlerts = ref<Alert[]>([])
const allAlerts = ref<Alert[]>([])

const isModalOpen = ref(false)
const loadError = ref<string | null>(null)

const range = computed(() => (selectedPeriod.value === 'day' ? 'today' : 'month'))
const isDayView = computed(() => selectedPeriod.value === 'day')

const lastUpdatedText = computed(() =>
    now.value.toLocaleString(undefined, {
        hour12: false,
        year: 'numeric',
        month: 'short',
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit',
    })
)

const chartTitle = computed(() =>
    isDayView.value ? "Today's Energy Usage" : 'Monthly Energy Usage'
)

const chartDescription = computed(() =>
    isDayView.value
        ? 'Consumption trend over the last 24 hours.'
        : 'Consumption trend for the current month.'
)

const chartSeriesName = computed(() =>
    isDayView.value ? 'Power Usage Today' : 'Monthly Energy Usage'
)

const hasChartData = computed(() => chartData.value.length > 0)
const hasActiveDevices = computed(() => activeDevices.value.length > 0)
const hasRecentAlerts = computed(() => recentAlerts.value.length > 0)
const activeChartData = computed(() => chartData.value)

const formatNumber = (value: number, digits = 2) => value.toFixed(digits)
const formatCurrency = (value: number) => `₱${value.toFixed(2)}`

const fetchSummary = async () => {
    try {
        isLoadingSummary.value = true
        const data = await readingService.getSummary({ range: range.value })
        summary.value = data
        now.value = new Date()
    } catch (error) {
        console.error('Failed to load summary:', error)
        loadError.value = 'Failed to load dashboard summary.'
    } finally {
        isLoadingSummary.value = false
    }
}

const fetchChart = async () => {
    try {
        isLoadingChart.value = true
        const data = await readingService.getChart(range.value)
        chartData.value = data
    } catch (error) {
        console.error('Failed to load chart:', error)
        loadError.value = 'Failed to load usage chart.'
        chartData.value = []
    } finally {
        isLoadingChart.value = false
    }
}

const fetchDevices = async () => {
    try {
        isLoadingDevices.value = true
        activeDevices.value = await applianceService.getDeviceStatus()
    } catch (error) {
        console.error('Failed to load devices:', error)
        activeDevices.value = []
    } finally {
        isLoadingDevices.value = false
    }
}

const fetchRecentAlerts = async () => {
    try {
        isLoadingRecentAlerts.value = true
        recentAlerts.value = await alertService.getRecent()
    } catch (error) {
        console.error('Failed to load recent alerts:', error)
        recentAlerts.value = []
    } finally {
        isLoadingRecentAlerts.value = false
    }
}

const fetchAllAlerts = async () => {
    try {
        isLoadingAlerts.value = true
        allAlerts.value = await alertService.getAll()
    } catch (error) {
        console.error('Failed to load all alerts:', error)
        allAlerts.value = []
    } finally {
        isLoadingAlerts.value = false
    }
}

const refreshDashboard = async () => {
    try {
        isRefreshing.value = true
        loadError.value = null
        await Promise.all([
            fetchSummary(),
            fetchChart(),
            fetchDevices(),
            fetchRecentAlerts(),
        ])
    } finally {
        isRefreshing.value = false
    }
}

const openModal = async () => {
    isModalOpen.value = true
    if (!allAlerts.value.length) {
        await fetchAllAlerts()
    }
}

const chartOptions = computed(() => {
    if (!colors.value) return {}

    return {
        tooltip: {
            trigger: 'axis',
            backgroundColor: colors.value.card,
            borderColor: colors.value.border,
            textStyle: {
                color: colors.value.cardForeground,
            },
            axisPointer: {
                type: 'line',
                lineStyle: {
                    color: colors.value.border,
                    opacity: 0.6,
                },
            },
        },
        grid: {
            left: 8,
            right: 8,
            top: 24,
            bottom: 8,
            containLabel: true,
        },
        xAxis: {
            type: 'category',
            data: activeChartData.value.map(item => item.label),
            boundaryGap: selectedPeriod.value !== 'day',
            axisLine: {
                lineStyle: {
                    color: colors.value.border,
                },
            },
            axisTick: {
                show: false,
            },
            axisLabel: {
                color: colors.value.mutedForeground,
            },
        },
        yAxis: {
            type: 'value',
            axisLine: {
                show: false,
            },
            axisTick: {
                show: false,
            },
            splitLine: {
                lineStyle: {
                    color: colors.value.border,
                    opacity: 0.35,
                },
            },
            axisLabel: {
                color: colors.value.mutedForeground,
            },
        },
        series: [
            {
                name: chartSeriesName.value,
                type: selectedPeriod.value === 'day' ? 'line' : 'bar',
                smooth: selectedPeriod.value === 'day',
                data: activeChartData.value.map(item => item.value),
                lineStyle: {
                    width: 3,
                    color: colors.value.chart1,
                },
                itemStyle: {
                    color: colors.value.chart1,
                    borderRadius: 6,
                },
                areaStyle:
                    selectedPeriod.value === 'day'
                        ? {
                            opacity: 0.16,
                            color: colors.value.chart1,
                        }
                        : undefined,
                barMaxWidth: 28,
                symbol: selectedPeriod.value === 'day' ? 'circle' : 'none',
                symbolSize: 7,
                showSymbol: selectedPeriod.value === 'day',
                emphasis: {
                    focus: 'series',
                },
            },
        ],
    }
})

onMounted(refreshDashboard)

watch(selectedPeriod, async () => {
    await Promise.all([fetchSummary(), fetchChart()])
})
</script>

<template>
    <div class="min-h-screen bg-muted/20">
        <div class="mx-auto w-full space-y-6 px-4 py-6 sm:px-6 lg:px-8">
            <!-- Hero header -->
            <section class="rounded-2xl border bg-background px-5 py-5 shadow-sm">
                <div class="flex flex-col gap-5 lg:flex-row lg:items-center lg:justify-between">
                    <div class="space-y-3">
                        <div class="flex items-center gap-3">
                            <div class="rounded-2xl bg-primary/10 p-2.5 text-primary">
                                <BarChart3Icon class="h-5 w-5" />
                            </div>
                            <div>
                                <h1 class="text-2xl font-semibold tracking-tight">Overview</h1>
                                <p class="text-sm text-muted-foreground">
                                    Monitor energy usage, device activity, and recent system alerts at a glance.
                                </p>
                            </div>
                        </div>

                        <div class="flex flex-wrap items-center gap-2">
                            <Badge variant="secondary" class="rounded-full px-3 py-1">
                                {{ isDayView ? 'Today' : 'This Month' }}
                            </Badge>
                            <Badge variant="outline" class="rounded-full px-3 py-1">
                                {{ summary.device_count }} registered devices
                            </Badge>
                            <Badge variant="outline" class="rounded-full px-3 py-1">
                                {{ summary.active_alerts }} active alerts
                            </Badge>
                        </div>
                    </div>

                    <div class="flex flex-col gap-3 sm:flex-row sm:flex-wrap sm:items-center">
                        <div class="flex gap-3">
                            <div class="inline-flex items-center rounded-xl border bg-muted/30 p-1">
                                <Button variant="ghost" size="sm" class="rounded-lg" :class="selectedPeriod === 'day'
                                    ? 'bg-primary text-primary-foreground hover:bg-primary hover:text-primary-foreground'
                                    : ''" @click="selectedPeriod = 'day'">
                                    Today
                                </Button>
                                <Button variant="ghost" size="sm" class="rounded-lg" :class="selectedPeriod === 'month'
                                    ? 'bg-primary text-primary-foreground hover:bg-primary hover:text-primary-foreground'
                                    : ''" @click="selectedPeriod = 'month'">
                                    This Month
                                </Button>
                            </div>

                            <div
                                class="flex items-center gap-2 rounded-xl border bg-background px-3 py-2 text-xs text-muted-foreground">
                                <span class="inline-block h-2 w-2 rounded-full bg-emerald-500" />
                                <span>Last updated: {{ lastUpdatedText }}</span>
                            </div>
                        </div>

                        <Button variant="outline" size="sm" class="gap-2 rounded-xl" :disabled="isRefreshing"
                            @click="refreshDashboard">
                            <RefreshCcwIcon class="h-4 w-4" :class="{ 'animate-spin': isRefreshing }" />
                            Refresh
                        </Button>
                    </div>
                </div>
            </section>

            <!-- Error banner -->
            <div v-if="loadError"
                class="rounded-xl border border-destructive/30 bg-destructive/5 px-4 py-3 text-sm text-destructive">
                {{ loadError }}
            </div>

            <!-- KPI cards -->
            <section class="grid gap-4 md:grid-cols-2 xl:grid-cols-4">
                <Card class="group rounded-2xl shadow-sm transition-all hover:-translate-y-0.5 hover:shadow-md">
                    <CardContent class="p-5">
                        <div class="flex items-start justify-between">
                            <div class="space-y-2">
                                <p class="text-sm font-medium text-muted-foreground">
                                    Total Energy
                                </p>
                                <div class="text-2xl font-semibold tracking-tight">
                                    <Skeleton v-if="isLoadingSummary" class="h-8 w-24" />
                                    <span v-else>{{ formatNumber(summary.total_energy_kwh) }} kWh</span>
                                </div>
                                <p class="text-xs text-muted-foreground">
                                    {{ isDayView ? 'Recorded today' : 'Recorded this month' }}
                                </p>
                            </div>
                            <div
                                class="rounded-xl bg-primary/10 p-2 text-primary transition-colors group-hover:bg-primary/15">
                                <BoltIcon class="h-4 w-4" />
                            </div>
                        </div>
                    </CardContent>
                </Card>

                <Card class="group rounded-2xl shadow-sm transition-all hover:-translate-y-0.5 hover:shadow-md">
                    <CardContent class="p-5">
                        <div class="flex items-start justify-between">
                            <div class="space-y-2">
                                <p class="text-sm font-medium text-muted-foreground">
                                    Estimated Bill
                                </p>
                                <div class="text-2xl font-semibold tracking-tight">
                                    <Skeleton v-if="isLoadingSummary" class="h-8 w-28" />
                                    <span v-else>{{ formatCurrency(summary.estimated_cost) }}</span>
                                </div>
                                <p class="text-xs text-muted-foreground">
                                    Rate: ₱{{ summary.billing_rate }}/kWh
                                </p>
                            </div>
                            <div
                                class="rounded-xl bg-primary/10 p-2 text-primary transition-colors group-hover:bg-primary/15">
                                <CircleDollarSignIcon class="h-4 w-4" />
                            </div>
                        </div>
                    </CardContent>
                </Card>

                <Card class="group rounded-2xl shadow-sm transition-all hover:-translate-y-0.5 hover:shadow-md">
                    <CardContent class="p-5">
                        <div class="flex items-start justify-between">
                            <div class="space-y-2">
                                <p class="text-sm font-medium text-muted-foreground">
                                    Active Devices
                                </p>
                                <div class="text-2xl font-semibold tracking-tight">
                                    <Skeleton v-if="isLoadingSummary" class="h-8 w-12" />
                                    <span v-else>{{ summary.active_devices }}</span>
                                </div>
                                <p class="text-xs text-muted-foreground">
                                    {{ summary.active_devices }} online out of {{ summary.device_count }}
                                </p>
                            </div>
                            <div
                                class="rounded-xl bg-primary/10 p-2 text-primary transition-colors group-hover:bg-primary/15">
                                <CpuIcon class="h-4 w-4" />
                            </div>
                        </div>
                    </CardContent>
                </Card>

                <Card class="group rounded-2xl shadow-sm transition-all hover:-translate-y-0.5 hover:shadow-md">
                    <CardContent class="p-5">
                        <div class="flex items-start justify-between">
                            <div class="space-y-2">
                                <p class="text-sm font-medium text-muted-foreground">
                                    Peak Power
                                </p>
                                <div class="text-2xl font-semibold tracking-tight">
                                    <Skeleton v-if="isLoadingSummary" class="h-8 w-20" />
                                    <span v-else>{{ formatNumber(summary.peak_power) }} W</span>
                                </div>
                                <p class="text-xs text-muted-foreground">
                                    Highest observed during this period
                                </p>
                            </div>
                            <div
                                class="rounded-xl bg-primary/10 p-2 text-primary transition-colors group-hover:bg-primary/15">
                                <ZapIcon class="h-4 w-4" />
                            </div>
                        </div>
                    </CardContent>
                </Card>
            </section>

            <!-- Main content -->
            <section class="grid gap-6 xl:grid-cols-12">
                <!-- Chart -->
                <Card class="rounded-2xl shadow-sm xl:col-span-8">
                    <CardHeader class="border-b pb-4">
                        <div class="flex flex-col gap-2 sm:flex-row sm:items-center sm:justify-between">
                            <div>
                                <CardTitle class="text-base">{{ chartTitle }}</CardTitle>
                                <CardDescription>
                                    {{ chartDescription }}
                                </CardDescription>
                            </div>

                            <Badge variant="outline" class="w-fit rounded-full px-3 py-1 text-[11px]">
                                {{ isDayView ? '24-hour view' : 'Monthly view' }}
                            </Badge>
                        </div>
                    </CardHeader>

                    <CardContent class="p-4 sm:p-6">
                        <div class="rounded-xl border bg-muted/20 p-3">
                            <div class="h-80">
                                <div v-if="isLoadingChart" class="flex h-full items-center justify-center">
                                    <div class="w-full space-y-3">
                                        <Skeleton class="h-4 w-32" />
                                        <Skeleton class="h-65 w-full rounded-xl" />
                                    </div>
                                </div>

                                <div v-else-if="!hasChartData"
                                    class="flex h-full flex-col items-center justify-center rounded-xl border border-dashed text-sm text-muted-foreground">
                                    <BoltIcon class="mb-2 h-5 w-5" />
                                    No chart data available for this period.
                                </div>

                                <VChart v-else class="h-full w-full" :option="chartOptions" autoresize />
                            </div>
                        </div>
                    </CardContent>
                </Card>

                <!-- Devices -->
                <Card class="rounded-2xl shadow-sm xl:col-span-4">
                    <CardHeader class="border-b pb-4">
                        <div class="flex items-start justify-between gap-3">
                            <div>
                                <CardTitle class="text-base">Device Status</CardTitle>
                                <CardDescription>
                                    Live overview of connected devices and current power draw.
                                </CardDescription>
                            </div>
                            <Badge variant="outline" class="rounded-full">
                                {{ activeDevices.length }} listed
                            </Badge>
                        </div>
                    </CardHeader>

                    <CardContent class="p-4 sm:p-5">
                        <div class="space-y-3">
                            <template v-if="isLoadingDevices">
                                <div v-for="n in 4" :key="n"
                                    class="flex items-center justify-between rounded-xl border px-3 py-3">
                                    <div class="space-y-2">
                                        <Skeleton class="h-4 w-28" />
                                        <Skeleton class="h-3 w-16" />
                                    </div>
                                    <div class="flex items-center gap-2">
                                        <Skeleton class="h-5 w-16" />
                                        <Skeleton class="h-4 w-12" />
                                    </div>
                                </div>
                            </template>

                            <template v-else-if="!hasActiveDevices">
                                <div
                                    class="flex min-h-65 flex-col items-center justify-center rounded-xl border border-dashed bg-muted/20 px-6 text-center">
                                    <ActivityIcon class="mb-2 h-5 w-5 text-muted-foreground" />
                                    <p class="text-sm font-medium">No active devices</p>
                                    <p class="text-xs text-muted-foreground">
                                        Online devices and live power readings will appear here.
                                    </p>
                                </div>
                            </template>

                            <template v-else>
                                <div v-for="device in activeDevices" :key="device.id"
                                    class="flex items-center justify-between rounded-xl border bg-muted/20 px-4 py-3 transition-colors hover:bg-muted/40">
                                    <div class="min-w-0 space-y-1">
                                        <p class="truncate text-sm font-medium">
                                            {{ device.name }}
                                        </p>
                                        <div class="flex items-center gap-2 text-xs text-muted-foreground">
                                            <span class="inline-block h-2 w-2 rounded-full"
                                                :class="device.status === 'online' ? 'bg-emerald-500' : 'bg-slate-400'" />
                                            <span>
                                                {{ device.status === 'online' ? 'Currently online' : 'Currently offline'
                                                }}
                                            </span>
                                        </div>
                                    </div>

                                    <div class="ml-4 flex items-center gap-2">
                                        <Badge :variant="device.status === 'online' ? 'default' : 'outline'"
                                            class="rounded-full px-2.5">
                                            {{ device.status === 'online' ? 'Online' : 'Offline' }}
                                        </Badge>
                                        <div v-if="device.status === 'online'"
                                            class="rounded-lg bg-background px-2.5 py-1 text-sm font-semibold">
                                            {{ device.power }} W
                                        </div>
                                    </div>
                                </div>
                            </template>
                        </div>
                    </CardContent>
                </Card>
            </section>

            <!-- Alerts -->
            <section>
                <Card class="rounded-2xl shadow-sm">
                    <CardHeader class="border-b pb-4">
                        <div class="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
                            <div>
                                <CardTitle class="text-base">Recent Alerts</CardTitle>
                                <CardDescription>
                                    Latest anomalies and warning events detected by the system.
                                </CardDescription>
                            </div>

                            <Button variant="outline" size="sm" class="rounded-xl" @click="openModal">
                                View all alerts
                            </Button>
                        </div>
                    </CardHeader>

                    <CardContent class="p-4 sm:p-5">
                        <div v-if="isLoadingRecentAlerts" class="space-y-3">
                            <div v-for="n in 3" :key="n"
                                class="flex items-start justify-between rounded-xl border px-4 py-3">
                                <div class="space-y-2">
                                    <Skeleton class="h-4 w-56" />
                                    <Skeleton class="h-3 w-24" />
                                </div>
                                <Skeleton class="h-5 w-14" />
                            </div>
                        </div>

                        <div v-else-if="!hasRecentAlerts"
                            class="flex min-h-45 flex-col items-center justify-center rounded-xl border border-dashed bg-muted/20 px-6 text-center">
                            <ShieldAlertIcon class="mb-2 h-5 w-5 text-muted-foreground" />
                            <p class="text-sm font-medium">No recent alerts</p>
                            <p class="text-xs text-muted-foreground">
                                Your system looks stable right now.
                            </p>
                        </div>

                        <div v-else class="grid gap-3 md:grid-cols-2">
                            <div v-for="alert in recentAlerts" :key="alert.id"
                                class="rounded-xl border bg-muted/20 p-4 transition-colors hover:bg-muted/40">
                                <div class="flex items-start justify-between gap-3">
                                    <div class="min-w-0">
                                        <p class="text-sm font-medium leading-5">
                                            {{ alert.message }}
                                        </p>
                                        <p class="mt-2 text-xs text-muted-foreground">
                                            {{ formatTime(alert.triggered_at) }}
                                        </p>
                                    </div>

                                    <Badge :variant="alert.severity === 'high' ? 'destructive' : 'outline'"
                                        class="shrink-0 rounded-full uppercase text-[10px]">
                                        {{ alert.severity }}
                                    </Badge>
                                </div>
                            </div>
                        </div>
                    </CardContent>
                </Card>
            </section>

            <Dialog v-model:open="isModalOpen">
                <DialogContent class="max-w-3xl rounded-2xl">
                    <DialogHeader>
                        <DialogTitle>All Alerts</DialogTitle>
                        <DialogDescription>
                            Full list of detected anomalies and system warnings.
                        </DialogDescription>
                    </DialogHeader>
                    <AlertTable :alerts="allAlerts" :loading="isLoadingAlerts" />
                </DialogContent>
            </Dialog>
        </div>
    </div>
</template>