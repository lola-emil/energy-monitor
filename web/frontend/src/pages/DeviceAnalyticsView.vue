<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import {
    Card,
    CardHeader,
    CardTitle,
    CardDescription,
    CardContent,
} from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { Skeleton } from '@/components/ui/skeleton'
import { Separator } from '@/components/ui/separator'
import {
    PowerIcon,
    ActivityIcon,
    WavesIcon,
    MapPinIcon,
    ArrowLeftIcon,
    LineChartIcon,
    AlertTriangleIcon,
    BoltIcon,
    PlugZapIcon,
} from 'lucide-vue-next'
import VChart from 'vue-echarts'
import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { LineChart } from 'echarts/charts'
import { GridComponent, TooltipComponent, LegendComponent } from 'echarts/components'

import { useThemeColors } from '@/composables/useThemeColors'
import { sseService } from '@/services/sse.service'
import { applianceService } from '@/services/appliance.service'
import { alertService } from '@/services/alert.service'
import type { Appliance } from '@/types/appliance'
import { formatTime } from '@/lib/time'

type ReadingUpdate = {
    time: string; power: number
}

use([
    CanvasRenderer,
    LineChart,
    GridComponent,
    TooltipComponent,
    LegendComponent,
])

const { colors } = useThemeColors()
const route = useRoute()
const router = useRouter()

const applianceId = computed(() => Number(route.params.id))

const isLoading = ref(true)
const isLoadingAlerts = ref(false)
const loadError = ref<string | null>(null)

const appliance = ref<Appliance | null>(null)

const voltage = ref<number | null>(null)
const current = ref<number | null>(null)
const power = ref<number | null>(null)
const frequency = ref<number | null>(null)
const todayEnergy = ref<number | null>(null)

const realtimeSeries = ref<ReadingUpdate[]>([])

const alerts = ref<any[]>([])

const applianceStatus = computed(() =>
    appliance.value?.status === 'online' ? 'Online' : 'Offline'
)

const lastUpdatedText = computed(() => {
    if (!appliance.value?.updated_at) return '—'
    return formatTime(appliance.value.updated_at)
})

const hasRealtimeData = computed(() => realtimeSeries.value.length > 0)

const formatMetric = (value: number | null, unit: string, digits = 2) => {
    if (value === null || value === undefined) return `0.00 ${unit}`
    return `${value.toFixed(digits)} ${unit}`
}

const fetchAppliance = async () => {
    if (!applianceId.value) return

    try {
        isLoading.value = true
        loadError.value = null
        const data = await applianceService.getById(applianceId.value)
        appliance.value = data
    } catch (error) {
        console.error(error)
        loadError.value = 'Failed to load appliance details.'
        appliance.value = null
    } finally {
        isLoading.value = false
    }
}

const fetchAlerts = async () => {
    if (!applianceId.value) return

    try {
        isLoadingAlerts.value = true
        const data = await alertService.getAlertsByAppliance(applianceId.value)
        alerts.value = Array.isArray(data) ? data : []
    } catch (error) {
        console.error(error)
        alerts.value = []
    } finally {
        isLoadingAlerts.value = false
    }
}

const updateRealtimeChart = (payload: any) => {
    realtimeSeries.value.push({
        time: new Date(payload.timestamp).toLocaleTimeString([], {
            hour: '2-digit',
            minute: '2-digit',
            second: '2-digit',
        }),
        power: Number(payload.power ?? 0),
    })

    if (realtimeSeries.value.length > 20) {
        realtimeSeries.value.shift()
    }
}

const handleLiveReading = (payload: any) => {
    if (!appliance.value?.device_code) return
    if (payload.device_code !== appliance.value.device_code) return

    voltage.value = payload.voltage ?? null
    current.value = payload.current ?? null
    power.value = payload.power ?? null
    frequency.value = payload.frequency ?? 60
    todayEnergy.value = payload.energy_kwh ?? null

    updateRealtimeChart(payload)
}

const refreshPage = async () => {
    await Promise.all([fetchAppliance(), fetchAlerts()])
}

const MAX_POINTS = 20

const realtimePowerChartOptions = computed(() => {
    if (!colors.value) return {}

    const points = realtimeSeries.value.slice(-MAX_POINTS)

    const labels = [
        ...points.map(i => i.time),
        ...Array(MAX_POINTS - points.length).fill(''),
    ]

    const values = [
        ...points.map(i => i.power),
        ...Array(MAX_POINTS - points.length).fill(null),
    ]

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

                if (!item || item.value == null) {
                    return ''
                }

                return `${item.axisValue}<br/>${Number(item.value).toFixed(2)} W`
            },
        },

        grid: {
            left: 16,
            right: 16,
            top: 20,
            bottom: 20,
            containLabel: true,
        },

        xAxis: {
            type: 'category',
            boundaryGap: false,
            data: labels,

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
                name: 'Power',
                type: 'line',
                smooth: true,
                showSymbol: false,

                data: values,

                lineStyle: {
                    width: 3,
                    color: colors.value.chart1,
                },

                itemStyle: {
                    color: colors.value.chart1,
                },

                areaStyle: {
                    opacity: 0.14,
                    color: colors.value.chart1,
                },
            },
        ],
    }
})

onMounted(async () => {
    await refreshPage()
    sseService.connect(handleLiveReading)
})

onUnmounted(() => {
    if (typeof sseService.disconnect === 'function') {
        sseService.disconnect()
    }
})

watch(
    () => route.params.id,
    async () => {
        realtimeSeries.value = []
        voltage.value = null
        current.value = null
        power.value = null
        frequency.value = null
        todayEnergy.value = null
        await refreshPage()
    }
)

watch(() => realtimeSeries.value, () => {
    console.log(realtimeSeries);
})
</script>

<template>
    <div class="min-h-screen bg-muted/20">
        <div class="mx-auto w-full space-y-6 px-4 py-6 sm:px-6 lg:px-8">
            <!-- Header -->
            <section class="rounded-2xl border bg-background px-5 py-5 shadow-sm">
                <div class="flex flex-col gap-4 lg:flex-row lg:items-start lg:justify-between">
                    <div class="space-y-3">
                        <div class="flex items-center gap-2">
                            <Button variant="ghost" size="icon" class="rounded-xl" @click="router.back()">
                                <ArrowLeftIcon class="h-4 w-4" />
                            </Button>

                            <div>
                                <div class="flex flex-wrap items-center gap-2">
                                    <h1 class="text-2xl font-semibold tracking-tight">
                                        {{ appliance?.name || 'Appliance' }}
                                    </h1>
                                    <Badge v-if="appliance"
                                        :variant="appliance.status === 'online' ? 'default' : 'outline'"
                                        class="rounded-full text-[11px]">
                                        {{ applianceStatus }}
                                    </Badge>
                                </div>
                                <p class="text-sm text-muted-foreground">
                                    Live energy and electrical metrics for this appliance.
                                </p>
                            </div>
                        </div>

                        <div class="flex flex-wrap items-center gap-3 text-xs text-muted-foreground">
                            <span class="flex items-center gap-1">
                                <MapPinIcon class="h-3.5 w-3.5" />
                                {{ appliance?.location || 'No location set' }}
                            </span>

                            <Separator orientation="vertical" class="hidden h-4 md:block" />

                            <span>Appliance ID: {{ appliance?.id || applianceId }}</span>

                            <Separator orientation="vertical" class="hidden h-4 md:block" />

                            <span>Last updated: {{ lastUpdatedText }}</span>
                        </div>
                    </div>

                    <div class="flex flex-wrap items-center gap-2">
                        <Button size="sm" variant="outline" class="rounded-xl" @click="refreshPage">
                            Refresh
                        </Button>
                    </div>
                </div>
            </section>

            <!-- Error state -->
            <div v-if="loadError"
                class="rounded-xl border border-destructive/30 bg-destructive/5 px-4 py-3 text-sm text-destructive">
                {{ loadError }}
            </div>

            <!-- Live metrics -->
            <section class="grid gap-4 md:grid-cols-2 xl:grid-cols-4">
                <Card class="rounded-2xl shadow-sm">
                    <CardContent class="p-5">
                        <div class="flex items-start justify-between">
                            <div class="space-y-2">
                                <p class="text-sm font-medium text-muted-foreground">Voltage</p>
                                <div class="text-2xl font-semibold tracking-tight">
                                    <Skeleton v-if="isLoading" class="h-8 w-16" />
                                    <span v-else>{{ formatMetric(voltage, 'V') }}</span>
                                </div>
                                <p class="text-xs text-muted-foreground">
                                    Instant RMS voltage at the appliance connection.
                                </p>
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
                                <p class="text-sm font-medium text-muted-foreground">Current</p>
                                <div class="text-2xl font-semibold tracking-tight">
                                    <Skeleton v-if="isLoading" class="h-8 w-16" />
                                    <span v-else>{{ formatMetric(current, 'A') }}</span>
                                </div>
                                <p class="text-xs text-muted-foreground">
                                    Instant load current drawn by this appliance.
                                </p>
                            </div>
                            <div class="rounded-xl bg-primary/10 p-2 text-primary">
                                <PlugZapIcon class="h-4 w-4" />
                            </div>
                        </div>
                    </CardContent>
                </Card>

                <Card class="rounded-2xl shadow-sm">
                    <CardContent class="p-5">
                        <div class="flex items-start justify-between">
                            <div class="space-y-2">
                                <p class="text-sm font-medium text-muted-foreground">Power</p>
                                <div class="text-2xl font-semibold tracking-tight">
                                    <Skeleton v-if="isLoading" class="h-8 w-20" />
                                    <span v-else>{{ formatMetric(power, 'W') }}</span>
                                </div>
                                <p class="text-xs text-muted-foreground">
                                    Current active power consumption.
                                </p>
                            </div>
                            <div class="rounded-xl bg-primary/10 p-2 text-primary">
                                <PowerIcon class="h-4 w-4" />
                            </div>
                        </div>
                    </CardContent>
                </Card>

                <Card class="rounded-2xl shadow-sm">
                    <CardContent class="p-5">
                        <div class="flex items-start justify-between">
                            <div class="space-y-2">
                                <p class="text-sm font-medium text-muted-foreground">Frequency</p>
                                <div class="text-2xl font-semibold tracking-tight">
                                    <Skeleton v-if="isLoading" class="h-8 w-20" />
                                    <span v-else>{{ formatMetric(frequency, 'Hz') }}</span>
                                </div>
                                <p class="text-xs text-muted-foreground">
                                    Measured line frequency at the monitored point.
                                </p>
                            </div>
                            <div class="rounded-xl bg-primary/10 p-2 text-primary">
                                <WavesIcon class="h-4 w-4" />
                            </div>
                        </div>
                    </CardContent>
                </Card>

                <!-- <Card class="rounded-2xl shadow-sm">
                    <CardContent class="p-5">
                        <div class="flex items-start justify-between">
                            <div class="space-y-2">
                                <p class="text-sm font-medium text-muted-foreground">Today's Energy</p>
                                <div class="text-2xl font-semibold tracking-tight">
                                    <Skeleton v-if="isLoading" class="h-8 w-24" />
                                    <span v-else>{{ formatMetric(todayEnergy, 'kWh') }}</span>
                                </div>
                                <p class="text-xs text-muted-foreground">
                                    Total energy consumed by this appliance today.
                                </p>
                            </div>
                            <div class="rounded-xl bg-primary/10 p-2 text-primary">
                                <BoltIcon class="h-4 w-4" />
                            </div>
                        </div>
                    </CardContent>
                </Card> -->
            </section>

            <!-- Real-time chart -->
            <section>
                <Card class="rounded-2xl shadow-sm">
                    <CardHeader class="border-b pb-4">
                        <div class="flex flex-col gap-2 sm:flex-row sm:items-center sm:justify-between">
                            <div>
                                <CardTitle class="text-base">Real-time Power</CardTitle>
                                <CardDescription>
                                    Latest rolling power readings for this appliance.
                                </CardDescription>
                            </div>

                            <Badge variant="outline" class="w-fit rounded-full">
                                Last {{ realtimeSeries.length }} points
                            </Badge>
                        </div>
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

                                <div v-else-if="!hasRealtimeData"
                                    class="flex h-full flex-col items-center justify-center rounded-xl border border-dashed text-sm text-muted-foreground">
                                    <PowerIcon class="mb-2 h-5 w-5" />
                                    Waiting for live readings for this appliance.
                                </div>

                                <VChart v-else class="h-full w-full" :option="realtimePowerChartOptions" autoresize />
                            </div>
                        </div>
                    </CardContent>
                </Card>
            </section>

            <!-- Alerts -->
            <section>
                <Card class="rounded-2xl shadow-sm">
                    <CardHeader class="border-b pb-4">
                        <div class="flex items-start justify-between gap-3">
                            <div>
                                <CardTitle class="text-base">Recent Alerts</CardTitle>
                                <CardDescription>
                                    Recent anomalies detected for this appliance.
                                </CardDescription>
                            </div>
                            <Badge variant="outline" class="rounded-full">
                                {{ alerts.length }} total
                            </Badge>
                        </div>
                    </CardHeader>

                    <CardContent class="p-4 sm:p-5">
                        <div v-if="isLoadingAlerts" class="space-y-3">
                            <div v-for="n in 4" :key="n" class="rounded-xl border p-4">
                                <Skeleton class="h-4 w-44" />
                                <Skeleton class="mt-2 h-3 w-24" />
                            </div>
                        </div>

                        <div v-else-if="!alerts.length"
                            class="flex min-h-45 flex-col items-center justify-center rounded-xl border border-dashed bg-muted/20 px-6 text-center">
                            <AlertTriangleIcon class="mb-2 h-5 w-5 text-muted-foreground" />
                            <p class="text-sm font-medium">No alerts recorded</p>
                            <p class="text-xs text-muted-foreground">
                                This appliance has no detected anomalies at the moment.
                            </p>
                        </div>

                        <div v-else class="space-y-3">
                            <div v-for="alert in alerts" :key="alert.id"
                                class="rounded-xl border bg-muted/20 p-4 transition-colors hover:bg-muted/40">
                                <div class="flex items-start justify-between gap-3">
                                    <div class="min-w-0">
                                        <p class="text-sm font-medium">{{ alert.message }}</p>
                                        <p class="mt-2 text-xs text-muted-foreground">
                                            {{ alert.time || formatTime(alert.triggered_at) }}
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