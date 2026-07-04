<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import {
    Table,
    TableHeader,
    TableRow,
    TableHead,
    TableCell,
    TableBody,
} from '@/components/ui/table'
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Select, SelectTrigger, SelectContent, SelectItem, SelectValue } from '@/components/ui/select'
import { Separator } from '@/components/ui/separator'
import { Skeleton } from '@/components/ui/skeleton'
import { Badge } from '@/components/ui/badge'
import DatePicker from '@/components/DatePicker.vue'
import {
    BoltIcon,
    GaugeIcon,
    ActivityIcon,
    CalendarRangeIcon,
    AlertTriangleIcon,
    DownloadIcon,
    ChevronLeftIcon,
    ChevronRightIcon,
} from 'lucide-vue-next'
import VChart from 'vue-echarts'
import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { LineChart } from 'echarts/charts'
import { GridComponent, TooltipComponent, LegendComponent } from 'echarts/components'

import { useThemeColors } from '@/composables/useThemeColors'
import { readingService, type AnalyticsResponse } from '@/services/reading.service'
import { applianceService } from '@/services/appliance.service'
import { alertService, type Alert } from '@/services/alert.service'
import type { Appliance } from '@/types/appliance'
import { formatTime } from '@/lib/time'

use([
    CanvasRenderer,
    LineChart,
    GridComponent,
    TooltipComponent,
    LegendComponent,
])

type RangeType = 'today' | '7d' | 'month'

const { colors } = useThemeColors()

const selectedRange = ref<RangeType>('today')
const selectedAppliance = ref<string>('all')

const lastUpdated = ref(new Date())

const isLoading = ref(true)
const isLoadingTable = ref(false)
const isLoadingAlerts = ref(false)
const isAppliancesLoading = ref(false)

const appliances = ref<Appliance[]>([])
const alerts = ref<Alert[]>([])
const detailedReadings = ref<any[]>([])

const total = ref(0)
const page = ref(1)
const pageSize = ref(10)

const defaultAnalytics: AnalyticsResponse = {
    summary: {
        total_energy_kwh: 0,
        avg_power: 0,
        avg_voltage: 0,
        avg_current: 0,
        peak_power: 0,
    },
    energy: [],
    voltage_current: [],
}

const analyticsData = ref<AnalyticsResponse>(defaultAnalytics)

const selectedMonth = ref({
    year: 2026,
    month: 6,
})

const selectedApplianceId = computed<number | undefined>(() => {
    if (selectedAppliance.value === 'all') return undefined
    return Number(selectedAppliance.value)
})

const totalPages = computed(() => Math.max(1, Math.ceil(total.value / pageSize.value)))

const selectedApplianceLabel = computed(() => {
    if (selectedAppliance.value === 'all') return 'All devices'
    const found = appliances.value.find(a => String(a.id) === selectedAppliance.value)
    return found?.name ?? 'Selected device'
})

const rangeLabel = computed(() => {
    switch (selectedRange.value) {
        case 'today':
            return 'Today'
        case '7d':
            return 'Last 7 days'
        case 'month':
            return 'Monthly'
    }
})


const showMonthPicker = computed(() => selectedRange.value == 'month')

const hasEnergyData = computed(() => analyticsData.value.energy?.length > 0)
const hasVoltageData = computed(() => analyticsData.value.voltage_current?.length > 0)

const fetchAppliances = async () => {
    try {
        isAppliancesLoading.value = true
        appliances.value = await applianceService.getAll()
    } catch (err) {
        console.error('Appliance list error:', err)
        appliances.value = []
    } finally {
        isAppliancesLoading.value = false
    }
}

const fetchAnalytics = async () => {
    try {
        isLoading.value = true

        let month: number | undefined
        let year: number | undefined

        if (selectedMonth.value.month && selectedMonth.value.year) {
            month = selectedMonth.value.month + 1
            year = selectedMonth.value.year
        }


        analyticsData.value = await readingService.getAnalytics({
            range: selectedRange.value,
            appliance_id: selectedApplianceId.value,
            month: month ? month.toString() : undefined,
            year: year ? year.toString() : undefined,
        })

        lastUpdated.value = new Date()
    } catch (err) {
        console.error('Analytics error:', err)
        analyticsData.value = defaultAnalytics
    } finally {
        isLoading.value = false
    }
}

const fetchDetailedReadings = async () => {
    try {
        isLoadingTable.value = true

        let month: number | undefined
        let year: number | undefined

        if (selectedMonth.value.month && selectedMonth.value.year) {
            month = selectedMonth.value.month + 1;
            year = selectedMonth.value.year;
        }

        const res = await readingService.getDetailedReadings({
            range: selectedRange.value,
            appliance_id: selectedApplianceId.value,
            page: page.value,
            page_size: pageSize.value,
            month: month ? month.toString() : undefined,
            year: year ? year.toString() : undefined,
        })
        detailedReadings.value = res.data || []
        total.value = res.total || 0
    } catch (err) {
        console.error('Detailed readings error:', err)
        detailedReadings.value = []
        total.value = 0
    } finally {
        isLoadingTable.value = false
    }
}

const fetchAlerts = async () => {
    try {
        isLoadingAlerts.value = true

        let month: number | undefined
        let year: number | undefined

        if (selectedMonth.value.month && selectedMonth.value.year) {
            month = selectedMonth.value.month + 1;
            year = selectedMonth.value.year;
        }

        const data = await alertService.getAnalyticsAlerts({
            range: selectedRange.value,
            appliance_id: selectedApplianceId.value,
            month: month ? month.toString() : undefined,
            year: year ? year.toString() : undefined,
        })
        alerts.value = Array.isArray(data) ? data : []
    } catch (err) {
        console.error('Alerts error:', err)
        alerts.value = []
    } finally {
        isLoadingAlerts.value = false
    }
}

const refreshAll = async () => {
    await Promise.all([
        fetchAnalytics(),
        fetchDetailedReadings(),
        fetchAlerts(),
    ])
}

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
            formatter: (params: any) => {
                const item = params?.[0]
                if (!item) return ''
                return `${item.axisValue}<br/>${Number(item.value).toFixed(2)} kWh`
            },
        },
        grid: {
            left: 10,
            right: 10,
            top: 24,
            bottom: 10,
            containLabel: true,
        },
        xAxis: {
            type: 'category',
            boundaryGap: false,
            data: analyticsData.value.energy?.map(i => i.label),
            axisLabel: {
                color: colors.value.mutedForeground,
            },
            axisLine: {
                lineStyle: {
                    color: colors.value.border,
                },
            },
            axisTick: {
                show: false,
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
            axisTick: {
                show: false,
            },
        },
        series: [
            {
                name: 'Energy Usage',
                type: 'line',
                smooth: true,
                showSymbol: true,
                symbol: 'circle',
                symbolSize: 7,
                data: analyticsData.value.energy?.map(i => i.value),
                lineStyle: {
                    width: 3,
                    color: colors.value.chart1,
                },
                itemStyle: {
                    color: colors.value.chart1,
                },
                areaStyle: {
                    opacity: 0.15,
                    color: colors.value.chart1,
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
            textStyle: {
                color: colors.value.cardForeground,
            },
        },
        legend: {
            top: 0,
            textStyle: {
                color: colors.value.mutedForeground,
            },
        },
        grid: {
            left: 20,
            right: 20,
            top: 50,
            bottom: 20,
            containLabel: true,
        },
        xAxis: {
            type: 'category',
            data: analyticsData.value.voltage_current?.map(i => i.label),
            axisLabel: {
                color: colors.value.mutedForeground,
            },
            axisLine: {
                lineStyle: {
                    color: colors.value.border,
                },
            },
            axisTick: {
                show: false,
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
                    show: false,
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
                data: analyticsData.value.voltage_current?.map(i => i.voltage),
                lineStyle: {
                    width: 3,
                    color: colors.value.chart1,
                },
                itemStyle: {
                    color: colors.value.chart1,
                },
                showSymbol: false,
            },
            {
                name: 'Current',
                type: 'line',
                smooth: true,
                yAxisIndex: 1,
                data: analyticsData.value.voltage_current?.map(i => i.current),
                lineStyle: {
                    width: 3,
                    color: colors.value.chart2,
                },
                itemStyle: {
                    color: colors.value.chart2,
                },
                showSymbol: false,
            },
        ],
    }
})

const prevPage = () => {
    if (page.value > 1) page.value--
}

const nextPage = () => {
    if (page.value < totalPages.value) page.value++
}

onMounted(async () => {
    await fetchAppliances()
    await refreshAll()
})

watch([selectedRange, selectedAppliance, selectedMonth], async () => {
    page.value = 1
    await refreshAll()
})

watch(page, fetchDetailedReadings)
</script>

<template>
    <div class="min-h-screen bg-muted/20">
        <div class="mx-auto w-full space-y-6 px-4 py-6 sm:px-6 lg:px-8">
            <!-- Header -->
            <section
                class="flex flex-col gap-4 rounded-2xl border bg-background px-5 py-5 shadow-sm md:flex-row md:items-center md:justify-between">
                <div class="space-y-1">
                    <h1 class="text-2xl font-semibold tracking-tight">Analytics</h1>
                    <p class="text-sm text-muted-foreground">
                        Explore historical energy usage, electrical measurements, and alerts across devices.
                    </p>
                </div>

                <div class="flex flex-col gap-3 lg:flex-row lg:items-center">
                    <!-- Device -->
                    <div class="flex items-center gap-2">
                        <span class="whitespace-nowrap text-xs text-muted-foreground">Device:</span>
                        <Select v-model="selectedAppliance">
                            <SelectTrigger class="w-45">
                                <SelectValue
                                    :placeholder="isAppliancesLoading ? 'Loading devices...' : 'Select device'" />
                            </SelectTrigger>
                            <SelectContent>
                                <SelectItem value="all">All devices</SelectItem>
                                <SelectItem v-for="dev in appliances" :key="dev.id" :value="String(dev.id)">
                                    {{ dev.name }}
                                </SelectItem>
                            </SelectContent>
                        </Select>
                    </div>

                    <!-- Range -->
                    <div class="flex items-center gap-2">
                        <span class="whitespace-nowrap text-xs text-muted-foreground">Range:</span>
                        <div class="inline-flex rounded-xl border bg-muted/30 p-1">
                            <Button size="sm" variant="ghost" class="rounded-lg" :class="selectedRange === 'today'
                                ? 'bg-primary text-primary-foreground hover:bg-primary hover:text-primary-foreground'
                                : ''" @click="selectedRange = 'today'">
                                Today
                            </Button>
                            <Button size="sm" variant="ghost" class="rounded-lg" :class="selectedRange === '7d'
                                ? 'bg-primary text-primary-foreground hover:bg-primary hover:text-primary-foreground'
                                : ''" @click="selectedRange = '7d'">
                                Last 7 days
                            </Button>
                            <Button size="sm" variant="ghost" class="rounded-lg" :class="selectedRange === 'month'
                                ? 'bg-primary text-primary-foreground hover:bg-primary hover:text-primary-foreground'
                                : ''" @click="selectedRange = 'month'">
                                Monthly
                            </Button>
                        </div>
                    </div>

                    <div v-if="showMonthPicker">
                        <span>For the month of: </span>
                        <DatePicker v-model="selectedMonth" />
                    </div>

                    <Separator orientation="vertical" class="hidden h-6 lg:block" />

                </div>
            </section>

            <!-- Filter summary -->
            <section class="flex flex-wrap items-center gap-2">
                <Badge variant="secondary" class="rounded-full px-3 py-1">
                    {{ selectedApplianceLabel }}
                </Badge>
                <Badge variant="outline" class="rounded-full px-3 py-1">
                    {{ rangeLabel }}
                </Badge>
            </section>

            <!-- Summary cards -->
            <section class="grid gap-4 md:grid-cols-2 xl:grid-cols-5">
                <Card class="rounded-2xl shadow-sm">
                    <CardContent class="p-5">
                        <div class="flex items-start justify-between">
                            <div class="space-y-2">
                                <p class="text-sm font-medium text-muted-foreground">Energy Used</p>
                                <div class="text-2xl font-semibold tracking-tight">
                                    <Skeleton v-if="isLoading" class="h-8 w-24" />
                                    <span v-else>{{ analyticsData.summary.total_energy_kwh.toFixed(2) }} kWh</span>
                                </div>
                                <p class="text-xs text-muted-foreground">Total for selected range.</p>
                            </div>
                            <div class="rounded-xl bg-primary/10 p-2 text-primary">
                                <BoltIcon class="h-4 w-4" />
                            </div>
                        </div>
                    </CardContent>
                </Card>

                <Card class="rounded-2xl shadow-sm">
                    <CardContent class="p-5">
                        <div class="flex items-start justify-between">
                            <div class="space-y-2">
                                <p class="text-sm font-medium text-muted-foreground">Average Power</p>
                                <div class="text-2xl font-semibold tracking-tight">
                                    <Skeleton v-if="isLoading" class="h-8 w-24" />
                                    <span v-else>{{ analyticsData.summary.avg_power.toFixed(2) }} W</span>
                                </div>
                                <p class="text-xs text-muted-foreground">Mean power draw over time.</p>
                            </div>
                            <div class="rounded-xl bg-primary/10 p-2 text-primary">
                                <GaugeIcon class="h-4 w-4" />
                            </div>
                        </div>
                    </CardContent>
                </Card>

                <Card class="rounded-2xl shadow-sm">
                    <CardContent class="p-5">
                        <div class="flex items-start justify-between">
                            <div class="space-y-2">
                                <p class="text-sm font-medium text-muted-foreground">Average Voltage</p>
                                <div class="text-2xl font-semibold tracking-tight">
                                    <Skeleton v-if="isLoading" class="h-8 w-20" />
                                    <span v-else>{{ analyticsData.summary.avg_voltage.toFixed(2) }} V</span>
                                </div>
                                <p class="text-xs text-muted-foreground">Indicates supply stability.</p>
                            </div>
                            <div class="rounded-xl bg-primary/10 p-2 text-primary">
                                <ActivityIcon class="h-4 w-4" />
                            </div>
                        </div>
                    </CardContent>
                </Card>

                <Card class="rounded-2xl shadow-sm">
                    <CardContent class="p-5">
                        <div class="flex items-start justify-between">
                            <div class="space-y-2">
                                <p class="text-sm font-medium text-muted-foreground">Average Current</p>
                                <div class="text-2xl font-semibold tracking-tight">
                                    <Skeleton v-if="isLoading" class="h-8 w-20" />
                                    <span v-else>{{ analyticsData.summary.avg_current.toFixed(2) }} A</span>
                                </div>
                                <p class="text-xs text-muted-foreground">Load drawn during the range.</p>
                            </div>
                            <div class="rounded-xl bg-primary/10 p-2 text-primary">
                                <ActivityIcon class="h-4 w-4" />
                            </div>
                        </div>
                    </CardContent>
                </Card>

                <Card class="rounded-2xl shadow-sm">
                    <CardContent class="p-5">
                        <div class="flex items-start justify-between">
                            <div class="space-y-2">
                                <p class="text-sm font-medium text-muted-foreground">Peak Power</p>
                                <div class="text-2xl font-semibold tracking-tight">
                                    <Skeleton v-if="isLoading" class="h-8 w-24" />
                                    <span v-else>{{ analyticsData.summary.peak_power.toFixed(2) }} W</span>
                                </div>
                                <p class="text-xs text-muted-foreground">Highest recorded value.</p>
                            </div>
                            <div class="rounded-xl bg-primary/10 p-2 text-primary">
                                <BoltIcon class="h-4 w-4" />
                            </div>
                        </div>
                    </CardContent>
                </Card>
            </section>

            <!-- Charts -->
            <section class="grid gap-6 xl:grid-cols-2">
                <Card class="rounded-2xl shadow-sm">
                    <CardHeader class="border-b pb-4">
                        <CardTitle class="text-base">Energy Usage Over Time</CardTitle>
                        <CardDescription>
                            Energy consumption trend for {{ selectedApplianceLabel.toLowerCase() }} during {{
                                rangeLabel.toLowerCase() }}.
                        </CardDescription>
                    </CardHeader>
                    <CardContent class="p-4 sm:p-6">
                        <div class="rounded-xl border bg-muted/20 p-3">
                            <div class="h-80">
                                <div v-if="isLoading" class="flex h-full items-center justify-center">
                                    <div class="w-full space-y-3">
                                        <Skeleton class="h-4 w-32" />
                                        <Skeleton class="h-65 w-full rounded-xl" />
                                    </div>
                                </div>

                                <div v-else-if="!hasEnergyData"
                                    class="flex h-full flex-col items-center justify-center rounded-xl border border-dashed text-sm text-muted-foreground">
                                    <BoltIcon class="mb-2 h-5 w-5" />
                                    No energy data available for the selected filters.
                                </div>

                                <VChart v-else class="h-full w-full" :option="energyChartOptions" autoresize />
                            </div>
                        </div>
                    </CardContent>
                </Card>

                <Card class="rounded-2xl shadow-sm">
                    <CardHeader class="border-b pb-4">
                        <CardTitle class="text-base">Voltage & Current Trends</CardTitle>
                        <CardDescription>
                            Helps analyze electrical stability and current draw patterns.
                        </CardDescription>
                    </CardHeader>
                    <CardContent class="p-4 sm:p-6">
                        <div class="rounded-xl border bg-muted/20 p-3">
                            <div class="h-80">
                                <div v-if="isLoading" class="flex h-full items-center justify-center">
                                    <div class="w-full space-y-3">
                                        <Skeleton class="h-4 w-40" />
                                        <Skeleton class="h-65 w-full rounded-xl" />
                                    </div>
                                </div>

                                <div v-else-if="!hasVoltageData"
                                    class="flex h-full flex-col items-center justify-center rounded-xl border border-dashed text-sm text-muted-foreground">
                                    <ActivityIcon class="mb-2 h-5 w-5" />
                                    No voltage/current data available for the selected filters.
                                </div>

                                <VChart v-else class="h-full w-full" :option="voltageChartOptions" autoresize />
                            </div>
                        </div>
                    </CardContent>
                </Card>
            </section>

            <!-- Table and alerts -->
            <section class="grid gap-6 xl:grid-cols-12">
                <!-- Detailed readings -->
                <Card class="rounded-2xl shadow-sm xl:col-span-8">
                    <CardHeader class="border-b pb-4">
                        <div class="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
                            <div>
                                <CardTitle class="text-base">Detailed Readings</CardTitle>
                                <CardDescription>
                                    Raw measurements for the selected filters.
                                </CardDescription>
                            </div>
                            <Button variant="outline" size="sm" class="gap-2 rounded-xl">
                                <DownloadIcon class="h-4 w-4" />
                                Export CSV
                            </Button>
                        </div>
                    </CardHeader>

                    <CardContent class="space-y-4 p-4 sm:p-6">
                        <div class="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
                            <p class="text-sm text-muted-foreground">
                                Showing page {{ page }} of {{ totalPages }}
                            </p>

                            <div class="flex items-center gap-2">
                                <Button variant="outline" size="sm" :disabled="page === 1" @click="prevPage">
                                    <ChevronLeftIcon class="h-4 w-4" />
                                    Previous
                                </Button>
                                <Button variant="outline" size="sm" :disabled="page >= totalPages" @click="nextPage">
                                    Next
                                    <ChevronRightIcon class="h-4 w-4" />
                                </Button>
                            </div>
                        </div>

                        <div class="overflow-hidden rounded-xl border">
                            <div class="overflow-x-auto">
                                <Table>
                                    <TableHeader>
                                        <TableRow class="bg-muted/30 hover:bg-muted/30">
                                            <TableHead>Timestamp</TableHead>
                                            <TableHead>Voltage (V)</TableHead>
                                            <TableHead>Current (A)</TableHead>
                                            <TableHead>Power (W)</TableHead>
                                            <TableHead>Energy (kWh)</TableHead>
                                        </TableRow>
                                    </TableHeader>

                                    <TableBody>
                                        <template v-if="isLoadingTable">
                                            <TableRow v-for="n in 6" :key="n">
                                                <TableCell colspan="5">
                                                    <Skeleton class="h-4 w-full" />
                                                </TableCell>
                                            </TableRow>
                                        </template>

                                        <template v-else-if="detailedReadings.length">
                                            <TableRow v-for="r in detailedReadings" :key="r.timestamp"
                                                class="hover:bg-muted/30">
                                                <TableCell class="whitespace-nowrap">{{ formatTime(r.timestamp) }}
                                                </TableCell>
                                                <TableCell>{{ r.voltage }}</TableCell>
                                                <TableCell>{{ r.current }}</TableCell>
                                                <TableCell>{{ r.power }}</TableCell>
                                                <TableCell>{{ r.energy_kwh }}</TableCell>
                                            </TableRow>
                                        </template>

                                        <TableRow v-else>
                                            <TableCell colspan="5"
                                                class="py-8 text-center text-sm text-muted-foreground">
                                                No readings found for the selected filters.
                                            </TableCell>
                                        </TableRow>
                                    </TableBody>
                                </Table>
                            </div>
                        </div>
                    </CardContent>
                </Card>

                <!-- Alerts -->
                <Card class="rounded-2xl shadow-sm xl:col-span-4">
                    <CardHeader class="border-b pb-4">
                        <div class="flex items-start justify-between gap-3">
                            <div>
                                <CardTitle class="text-base">Alerts in This Period</CardTitle>
                                <CardDescription>
                                    Anomalies detected for the selected device and range.
                                </CardDescription>
                            </div>
                            <Badge variant="outline" class="rounded-full">
                                {{ alerts.length }} total
                            </Badge>
                        </div>
                    </CardHeader>

                    <CardContent class="p-4 sm:p-5 flex-1">
                        <div v-if="isLoadingAlerts" class="space-y-3">
                            <div v-for="n in 4" :key="n" class="rounded-xl border p-4">
                                <Skeleton class="h-4 w-40" />
                                <Skeleton class="mt-2 h-3 w-24" />
                            </div>
                        </div>

                        <div v-else-if="!alerts.length"
                            class="flex h-full flex-col items-center justify-center rounded-xl border border-dashed bg-muted/20 px-6 text-center">
                            <AlertTriangleIcon class="mb-2 h-5 w-5 text-muted-foreground" />
                            <p class="text-sm font-medium">No alerts found</p>
                            <p class="text-xs text-muted-foreground">
                                No anomalies were detected for the selected filters.
                            </p>
                        </div>

                        <div v-else class="space-y-3">
                            <div v-for="alert in alerts" :key="alert.id"
                                class="rounded-xl border bg-muted/20 p-4 transition-colors hover:bg-muted/40">
                                <div class="flex items-start justify-between gap-3">
                                    <div class="min-w-0">
                                        <p class="text-sm font-medium leading-5">
                                            {{ alert.message }}
                                        </p>
                                        <p class="mt-2 text-xs text-muted-foreground">
                                            {{ alert.name }} | {{ alert.triggered_at }}
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
        </div>
    </div>
</template>