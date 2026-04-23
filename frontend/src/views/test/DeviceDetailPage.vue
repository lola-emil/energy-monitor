<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'

import {
    Card,
    CardHeader,
    CardTitle,
    CardDescription,
    CardContent
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
    AlertTriangleIcon
} from 'lucide-vue-next'

const route = useRoute()
const router = useRouter()
const applianceId = route.params.id as string

const isLoading = ref(true)

// Mock appliance data – replace with API call
const appliance = ref<null | {
    id: string
    name: string
    location: string
    status: 'online' | 'offline'
    lastUpdate: string
}>()

// Instant measurements for this appliance
const voltage = ref<number | null>(null)
const current = ref<number | null>(null)
const power = ref<number | null>(null)
const frequency = ref<number | null>(null)
const todayEnergy = ref<number | null>(null)

// Real‑time series for last N minutes (power vs time)
const realtimeSeries = ref(
    Array.from({ length: 10 }).map((_, i) => ({
        time: `${10 + i}:00`,
        power: 400 + Math.round(Math.random() * 80)
    }))
)

// Simple “today” energy profile
const dailyEnergySeries = ref([
    { hour: '00:00', energy: 0.2 },
    { hour: '06:00', energy: 0.4 },
    { hour: '12:00', energy: 0.8 },
    { hour: '18:00', energy: 0.6 }
])

// Alerts for this appliance
const alerts = ref([
    { id: 1, time: '2026-04-22 16:05', message: 'Appliance went offline', severity: 'medium' },
    { id: 2, time: '2026-04-21 09:12', message: 'Voltage exceeded 250 V', severity: 'high' }
])

const lastUpdatedText = computed(() => appliance.value?.lastUpdate ?? '—')

onMounted(() => {
    // Simulate fetch from backend
    setTimeout(() => {
        appliance.value = {
            id: applianceId,
            name: applianceId === '1' ? 'Refrigerator' : `Appliance ${applianceId}`,
            location: 'Kitchen',
            status: 'online',
            lastUpdate: '2026-04-23 10:35'
        }
        voltage.value = 229
        current.value = 2.2
        power.value = 502
        frequency.value = 60
        todayEnergy.value = 5.32

        isLoading.value = false

        // TODO: start WebSocket / SSE subscription for this appliance here
    }, 600)
})

// TODO: onUnmounted -> close realtime connection when using WebSocket/SSE
</script>

<template>
    <div class="flex flex-col gap-6">
        <!-- Header -->
        <div class="flex flex-col gap-3 md:flex-row md:items-center md:justify-between">
            <div class="space-y-1">
                <div class="flex items-center gap-2">
                    <Button variant="ghost" size="icon" @click="router.back()">
                        <ArrowLeftIcon class="h-4 w-4" />
                    </Button>
                    <h1 class="text-2xl font-semibold tracking-tight">
                        {{ appliance?.name || 'Appliance' }}
                    </h1>
                    <Badge v-if="appliance" :variant="appliance.status === 'online' ? 'default' : 'outline'"
                        class="ml-1 text-[11px]">
                        {{ appliance.status === 'online' ? 'Online' : 'Offline' }}
                    </Badge>
                </div>
                <div class="flex flex-wrap items-center gap-3 text-xs text-muted-foreground">
                    <span class="flex items-center gap-1">
                        <MapPinIcon class="h-3 w-3" />
                        {{ appliance?.location || '—' }}
                    </span>
                    <Separator orientation="vertical" class="hidden h-4 md:block" />
                    <span>Appliance ID: {{ appliance?.id || applianceId }}</span>
                    <Separator orientation="vertical" class="hidden h-4 md:block" />
                    <span>Last updated: {{ lastUpdatedText }}</span>
                </div>
            </div>

            <Button size="sm" variant="outline" @click="router.push({ name: 'analytics', query: { applianceId } })">
                <LineChartIcon class="mr-2 h-4 w-4" />
                View in Analytics
            </Button>
        </div>

        <!-- Instant metrics -->
        <div class="grid gap-4 md:grid-cols-2 xl:grid-cols-5">
            <!-- Voltage -->
            <Card>
                <CardHeader class="flex flex-row items-center justify-between space-y-0 pb-2">
                    <CardTitle class="text-sm font-medium">Voltage</CardTitle>
                    <ActivityIcon class="h-4 w-4 text-primary" />
                </CardHeader>
                <CardContent>
                    <p class="text-2xl font-bold">
                        <Skeleton v-if="isLoading" class="h-7 w-16" />
                        <span v-else>{{ voltage }} V</span>
                    </p>
                    <p class="mt-1 text-xs text-muted-foreground">
                        Instant RMS voltage at this appliance’s connection point.
                    </p>
                </CardContent>
            </Card>

            <!-- Current -->
            <Card>
                <CardHeader class="flex flex-row items-center justify-between space-y-0 pb-2">
                    <CardTitle class="text-sm font-medium">Current</CardTitle>
                    <ActivityIcon class="h-4 w-4 text-primary" />
                </CardHeader>
                <CardContent>
                    <p class="text-2xl font-bold">
                        <Skeleton v-if="isLoading" class="h-7 w-16" />
                        <span v-else>{{ current }} A</span>
                    </p>
                    <p class="mt-1 text-xs text-muted-foreground">
                        Instant load current drawn by this appliance.
                    </p>
                </CardContent>
            </Card>

            <!-- Power -->
            <Card>
                <CardHeader class="flex flex-row items-center justify-between space-y-0 pb-2">
                    <CardTitle class="text-sm font-medium">Power</CardTitle>
                    <PowerIcon class="h-4 w-4 text-primary" />
                </CardHeader>
                <CardContent>
                    <p class="text-2xl font-bold">
                        <Skeleton v-if="isLoading" class="h-7 w-20" />
                        <span v-else>{{ power }} W</span>
                    </p>
                    <p class="mt-1 text-xs text-muted-foreground">
                        Instant active power consumption of this appliance.
                    </p>
                </CardContent>
            </Card>

            <!-- Frequency -->
            <Card>
                <CardHeader class="flex flex-row items-center justify-between space-y-0 pb-2">
                    <CardTitle class="text-sm font-medium">Frequency</CardTitle>
                    <WavesIcon class="h-4 w-4 text-primary" />
                </CardHeader>
                <CardContent>
                    <p class="text-2xl font-bold">
                        <Skeleton v-if="isLoading" class="h-7 w-20" />
                        <span v-else>{{ frequency }} Hz</span>
                    </p>
                    <p class="mt-1 text-xs text-muted-foreground">
                        Measured line frequency at the monitored point.
                    </p>
                </CardContent>
            </Card>

            <!-- Today's energy -->
            <Card>
                <CardHeader class="flex flex-row items-center justify-between space-y-0 pb-2">
                    <CardTitle class="text-sm font-medium">Today&apos;s energy</CardTitle>
                    <PowerIcon class="h-4 w-4 text-primary" />
                </CardHeader>
                <CardContent>
                    <p class="text-2xl font-bold">
                        <Skeleton v-if="isLoading" class="h-7 w-24" />
                        <span v-else>{{ todayEnergy }} kWh</span>
                    </p>
                    <p class="mt-1 text-xs text-muted-foreground">
                        Energy consumed by this appliance today.
                    </p>
                </CardContent>
            </Card>
        </div>

        <!-- Charts -->
        <div class="grid gap-4 lg:grid-cols-3">
            <!-- Real-time power chart -->
            <Card class="lg:col-span-2">
                <CardHeader>
                    <CardTitle>Real-time power</CardTitle>
                    <CardDescription>
                        Recent power readings for this appliance (auto-refreshing).
                    </CardDescription>
                </CardHeader>
                <CardContent>
                    <!-- Replace with your chart component -->
                    <div class="h-64 rounded-md border border-dashed border-muted p-3 text-xs text-muted-foreground">
                        <p class="mb-2 font-medium">Chart placeholder</p>
                        <ul class="space-y-1">
                            <li v-for="p in realtimeSeries" :key="p.time">
                                {{ p.time }} — {{ p.power }} W
                            </li>
                        </ul>
                    </div>
                </CardContent>
            </Card>

            <!-- Today usage chart -->
            <Card>
                <CardHeader>
                    <CardTitle>Today&apos;s usage</CardTitle>
                    <CardDescription>
                        Approximate energy usage profile for this appliance today.
                    </CardDescription>
                </CardHeader>
                <CardContent>
                    <div class="h-64 rounded-md border border-dashed border-muted p-3 text-xs text-muted-foreground">
                        <p class="mb-2 font-medium">Chart placeholder</p>
                        <ul class="space-y-1">
                            <li v-for="p in dailyEnergySeries" :key="p.hour">
                                {{ p.hour }} — {{ p.energy }} kWh
                            </li>
                        </ul>
                    </div>
                </CardContent>
            </Card>
        </div>

        <!-- Alerts -->
        <Card>
            <CardHeader class="flex flex-row items-center justify-between">
                <div>
                    <CardTitle>Recent alerts for this appliance</CardTitle>
                    <CardDescription>
                        Anomalies detected on this appliance’s connection point.
                    </CardDescription>
                </div>
                <Badge variant="outline">{{ alerts.length }} total</Badge>
            </CardHeader>
            <CardContent>
                <div v-if="!alerts.length"
                    class="flex flex-col items-center justify-center py-6 text-sm text-muted-foreground">
                    <AlertTriangleIcon class="mb-2 h-5 w-5" />
                    No alerts recorded for this appliance.
                </div>
                <div v-else class="space-y-2 text-xs">
                    <div v-for="alert in alerts" :key="alert.id"
                        class="flex items-start justify-between rounded-md border px-3 py-2">
                        <div>
                            <p class="text-sm font-medium">{{ alert.message }}</p>
                            <p class="text-[11px] text-muted-foreground">{{ alert.time }}</p>
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
</template>