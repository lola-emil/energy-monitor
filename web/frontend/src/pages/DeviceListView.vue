<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import {
    Card,
    CardHeader,
    CardTitle,
    CardDescription,
    CardContent,
    CardFooter,
} from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { Input } from '@/components/ui/input'
import { Skeleton } from '@/components/ui/skeleton'
import {
    SearchIcon,
    MapPinIcon,
    PowerIcon,
    MonitorIcon,
    AlertCircleIcon,
} from 'lucide-vue-next'

import { applianceService } from '@/services/appliance.service'
import type { Appliance } from '@/types/appliance'
import AddApplianceDialog from '@/components/AddDeviceModal.vue'
import EditDeviceModal from '@/components/EditDeviceModal.vue'
import DeleteDeviceAlert from '@/components/DeleteDeviceAlert.vue'
import { formatTime } from '@/lib/time'

const appliances = ref<Appliance[] | null>([])
const isLoading = ref(true)
const errorMessage = ref('')
const search = ref('')

const filteredAppliances = computed(() => {
    const term = search.value.trim().toLowerCase()
    if (!term) return appliances.value

    return appliances.value?.filter(appliance =>
        appliance.name.toLowerCase().includes(term) ||
        String(appliance.id).includes(term) ||
        (appliance.location ?? '').toLowerCase().includes(term)
    )
})

const totalCount = computed(() => appliances.value?.length)

const onlineCount = computed(() =>
    appliances.value?.filter(appliance => appliance.status === 'online').length
)

const offlineCount = computed(() =>
    appliances.value?.filter(appliance => appliance.status === 'offline').length
)

const hasAppliances = computed(() => appliances.value ? appliances.value.length > 0 : false)
const hasFilteredAppliances = computed(() => filteredAppliances.value ? filteredAppliances.value.length > 0 : false)
const isSearching = computed(() => search.value.trim().length > 0)

const fetchAppliances = async () => {
    try {
        isLoading.value = true
        errorMessage.value = ''
        appliances.value = await applianceService.getAll()
    } catch (error: any) {
        console.error(error)
        errorMessage.value =
            error?.response?.data?.message ||
            'Failed to load appliances.'
        appliances.value = []
    } finally {
        isLoading.value = false
    }
}

const handleCreateAppliance = async (payload: {
    name: string
    location: string
}) => {
    try {
        await applianceService.create(payload)
        await fetchAppliances()
    } catch (error) {
        console.error(error)
    }
}

const handleUpdateAppliance = async (payload: {
    id: number
    name: string
    location: string
}) => {
    try {
        await applianceService.update(payload.id, {
            name: payload.name,
            location: payload.location,
        })
        await fetchAppliances()
    } catch (error) {
        console.error(error)
    }
}

const handleDeleteAppliance = async (id: number) => {
    try {
        await applianceService.delete(id)
        await fetchAppliances()
    } catch (error) {
        console.error(error)
    }
}

onMounted(fetchAppliances)
</script>

<template>
    <div class="min-h-screen bg-muted/20">
        <div class="mx-auto w-full space-y-6 px-4 py-6 sm:px-6 lg:px-8">
            <!-- Header -->
            <section
                class="flex flex-col gap-4 rounded-2xl border bg-background px-5 py-5 shadow-sm md:flex-row md:items-center md:justify-between">
                <div class="space-y-1">
                    <h1 class="text-2xl font-semibold tracking-tight">Devices</h1>
                    <p class="text-sm text-muted-foreground">
                        View, manage, and monitor all registered appliances in your system.
                    </p>
                </div>

                <AddApplianceDialog @submit="handleCreateAppliance" />
            </section>

            <!-- Error -->
            <div v-if="errorMessage"
                class="rounded-xl border border-destructive/30 bg-destructive/5 px-4 py-3 text-sm text-destructive">
                {{ errorMessage }}
            </div>

            <!-- Summary and search -->
            <section
                class="flex flex-col gap-4 rounded-2xl border bg-background px-5 py-4 shadow-sm lg:flex-row lg:items-center lg:justify-between">
                <div class="flex flex-wrap gap-2">
                    <Badge variant="outline" class="rounded-full px-3 py-1">
                        Total: {{ totalCount }}
                    </Badge>
                    <Badge variant="outline"
                        class="rounded-full border-emerald-500/30 bg-emerald-500/10 px-3 py-1 text-emerald-600 dark:text-emerald-400">
                        Online: {{ onlineCount }}
                    </Badge>
                    <Badge variant="outline"
                        class="rounded-full border-red-500/30 bg-red-500/10 px-3 py-1 text-red-600 dark:text-red-400">
                        Offline: {{ offlineCount }}
                    </Badge>
                </div>

                <div class="relative w-full lg:w-[320px]">
                    <SearchIcon class="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
                    <Input v-model="search" class="pl-9" placeholder="Search by name, ID, or location" />
                </div>
            </section>

            <!-- Loading state -->
            <section v-if="isLoading" class="grid gap-4 md:grid-cols-2 xl:grid-cols-3">
                <Card v-for="i in 6" :key="i" class="rounded-2xl shadow-sm">
                    <CardHeader class="space-y-2">
                        <Skeleton class="h-5 w-32" />
                        <Skeleton class="h-4 w-24" />
                    </CardHeader>
                    <CardContent class="space-y-3">
                        <Skeleton class="h-4 w-full" />
                        <Skeleton class="h-4 w-3/4" />
                        <Skeleton class="h-4 w-2/3" />
                    </CardContent>
                    <CardFooter class="flex gap-2">
                        <Skeleton class="h-9 w-24" />
                        <Skeleton class="h-9 w-20" />
                        <Skeleton class="h-9 w-20" />
                    </CardFooter>
                </Card>
            </section>

            <!-- Empty state: no appliances at all -->
            <section v-else-if="!hasAppliances"
                class="rounded-2xl border border-dashed bg-background px-6 py-12 text-center shadow-sm">
                <div class="mx-auto max-w-md space-y-3">
                    <div class="mx-auto flex h-12 w-12 items-center justify-center rounded-full bg-muted">
                        <MonitorIcon class="h-5 w-5 text-muted-foreground" />
                    </div>
                    <h2 class="text-lg font-semibold">No devices yet</h2>
                    <p class="text-sm text-muted-foreground">
                        You haven’t added any appliances yet. Add your first device to start monitoring usage and
                        activity.
                    </p>
                    <div class="pt-2">
                        <AddApplianceDialog @submit="handleCreateAppliance" />
                    </div>
                </div>
            </section>

            <!-- Empty state: search results -->
            <section v-else-if="!hasFilteredAppliances"
                class="rounded-2xl border border-dashed bg-background px-6 py-12 text-center shadow-sm">
                <div class="mx-auto max-w-md space-y-3">
                    <div class="mx-auto flex h-12 w-12 items-center justify-center rounded-full bg-muted">
                        <AlertCircleIcon class="h-5 w-5 text-muted-foreground" />
                    </div>
                    <h2 class="text-lg font-semibold">No matching devices</h2>
                    <p class="text-sm text-muted-foreground">
                        No devices matched your search for
                        <span class="font-medium text-foreground">“{{ search }}”</span>.
                        Try a different name, ID, or location.
                    </p>
                    <div>
                        <Button variant="outline" size="sm" @click="search = ''">
                            Clear search
                        </Button>
                    </div>
                </div>
            </section>

            <!-- Device cards -->
            <section v-else class="grid gap-4 md:grid-cols-2 xl:grid-cols-3">
                <Card v-for="appliance in filteredAppliances" :key="appliance.id"
                    class="group flex flex-col justify-between rounded-2xl shadow-sm transition-all hover:-translate-y-0.5 hover:shadow-md">
                    <CardHeader class="space-y-3 pb-4">
                        <div class="flex items-start justify-between gap-3">
                            <div class="min-w-0">
                                <CardTitle class="truncate text-base font-semibold">
                                    {{ appliance.name }}
                                </CardTitle>
                                <CardDescription class="mt-1 flex items-center gap-1 text-xs">
                                    <MapPinIcon class="h-3 w-3" />
                                    <span class="truncate">{{ appliance.location || 'No location set' }}</span>
                                </CardDescription>
                            </div>

                            <Badge :variant="appliance.status === 'online' ? 'default' : 'outline'"
                                class="shrink-0 rounded-full text-[11px]">
                                {{ appliance.status === 'online' ? 'Online' : 'Offline' }}
                            </Badge>
                        </div>
                    </CardHeader>

                    <CardContent class="space-y-3 pt-0 text-sm">
                        <div class="rounded-xl border bg-muted/20 px-3 py-3">
                            <div class="flex items-center justify-between">
                                <span class="flex items-center gap-2 text-muted-foreground">
                                    <PowerIcon class="h-4 w-4" />
                                    Current power
                                </span>
                                <span class="font-medium">
                                    <span v-if="appliance.status === 'online'">—</span>
                                    <span v-else class="text-muted-foreground">Offline</span>
                                </span>
                            </div>
                        </div>

                        <div class="space-y-2 text-xs text-muted-foreground">
                            <div class="flex items-center justify-between">
                                <span>Device ID</span>
                                <span class="font-medium text-foreground">{{ appliance.id }}</span>
                            </div>
                            <div class="flex items-center justify-between">
                                <span>Last update</span>
                                <span class="font-medium text-foreground">
                                    {{ formatTime(appliance.updated_at) }}
                                </span>
                            </div>
                        </div>
                    </CardContent>

                    <CardFooter class="flex flex-wrap items-center justify-end gap-2 pt-4">
                        <Button size="sm" variant="outline" class="rounded-xl"
                            @click="$router.push(`/devices/${appliance.id}`)">
                            View Monitor
                        </Button>

                        <EditDeviceModal :appliance="appliance" @submit="handleUpdateAppliance" />

                        <DeleteDeviceAlert :appliance="appliance" @submit="handleDeleteAppliance" />
                    </CardFooter>
                </Card>
            </section>
        </div>
    </div>
</template>