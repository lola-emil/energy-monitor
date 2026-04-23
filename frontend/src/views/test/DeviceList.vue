<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import {
    Card,
    CardHeader,
    CardTitle,
    CardDescription,
    CardContent,
    CardFooter
} from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { Input } from '@/components/ui/input'
import { Skeleton } from '@/components/ui/skeleton'
import { PlusIcon, PowerIcon, MapPinIcon } from 'lucide-vue-next'

const isLoading = ref(true)

// Mock data – replace with API call
const appliances = ref([
    {
        id: '1',
        name: 'Refrigerator',
        location: 'Kitchen',
        status: 'online',
        power: 120,
        lastUpdate: '2026-04-23 10:32'
    },
    {
        id: '2',
        name: 'Aircon #1',
        location: 'Living Room',
        status: 'online',
        power: 780,
        lastUpdate: '2026-04-23 10:31'
    },
    {
        id: '3',
        name: 'Water Pump',
        location: 'Basement',
        status: 'offline',
        power: 0,
        lastUpdate: '2026-04-22 08:10'
    }
])

const search = ref('')

const filteredAppliances = computed(() => {
    if (!search.value) return appliances.value
    const term = search.value.toLowerCase()
    return appliances.value.filter(a =>
        a.name.toLowerCase().includes(term) ||
        a.id.toLowerCase().includes(term) ||
        a.location.toLowerCase().includes(term)
    )
})

const totalCount = computed(() => appliances.value.length)
const onlineCount = computed(() =>
    appliances.value.filter(a => a.status === 'online').length
)
const offlineCount = computed(() =>
    appliances.value.filter(a => a.status === 'offline').length
)

onMounted(() => {
    setTimeout(() => {
        isLoading.value = false
    }, 600)
    // TODO: replace with real fetch from backend
})
</script>

<template>
    <div class="px-5">
        <div class="flex flex-col gap-6">
            <!-- Header -->
            <div class="flex flex-col gap-4 md:flex-row md:items-center md:justify-between">
                <div>
                    <h1 class="text-2xl font-semibold tracking-tight">Appliances</h1>
                    <p class="text-sm text-muted-foreground">
                        View and manage all monitored appliances.
                    </p>
                </div>
                <Button size="sm">
                    <PlusIcon class="mr-2 h-4 w-4" />
                    Add appliance
                </Button>
            </div>

            <!-- Summary + search -->
            <div class="flex flex-col gap-3 md:flex-row md:items-center md:justify-between">
                <div class="flex flex-wrap gap-2 text-xs">
                    <Badge variant="outline">Total: {{ totalCount }}</Badge>
                    <Badge variant="outline" class="bg-emerald-500/10 text-emerald-400 border-emerald-500/40">
                        Online: {{ onlineCount }}
                    </Badge>
                    <Badge variant="outline" class="bg-red-500/10 text-red-400 border-red-500/40">
                        Offline: {{ offlineCount }}
                    </Badge>
                </div>
                <Input v-model="search" class="md:w-65" placeholder="Search by name, ID, or location" />
            </div>

            <!-- Appliance cards -->
            <div v-if="isLoading" class="grid gap-4 md:grid-cols-2 xl:grid-cols-3">
                <Skeleton v-for="i in 3" :key="i" class="h-40 w-full rounded-xl" />
            </div>

            <div v-else-if="!filteredAppliances.length" class="py-8 text-center text-sm text-muted-foreground">
                No appliances found. Try adjusting your search or add a new appliance.
            </div>

            <div v-else class="grid gap-4 md:grid-cols-2 xl:grid-cols-3">
                <Card v-for="appliance in filteredAppliances" :key="appliance.id" class="flex flex-col justify-between">
                    <CardHeader class="space-y-1">
                        <div class="flex items-center justify-between">
                            <CardTitle class="text-base font-semibold">
                                {{ appliance.name }}
                            </CardTitle>
                            <Badge :variant="appliance.status === 'online' ? 'default' : 'outline'" class="text-[11px]">
                                {{ appliance.status === 'online' ? 'Online' : 'Offline' }}
                            </Badge>
                        </div>
                        <CardDescription class="flex items-center gap-1 text-xs">
                            <MapPinIcon class="h-3 w-3" />
                            {{ appliance.location }}
                        </CardDescription>
                    </CardHeader>

                    <CardContent class="space-y-2 text-xs">
                        <div class="flex items-center justify-between">
                            <span class="flex items-center gap-1 text-muted-foreground">
                                <PowerIcon class="h-3 w-3" />
                                Current power
                            </span>
                            <span class="text-sm font-medium">
                                <span v-if="appliance.status === 'online'">
                                    {{ appliance.power }} W
                                </span>
                                <span v-else class="text-muted-foreground">—</span>
                            </span>
                        </div>
                        <div class="flex items-center justify-between text-[11px] text-muted-foreground">
                            <span>Appliance ID:</span>
                            <span>{{ appliance.id }}</span>
                        </div>
                        <div class="flex items-center justify-between text-[11px] text-muted-foreground">
                            <span>Last update:</span>
                            <span>{{ appliance.lastUpdate }}</span>
                        </div>
                    </CardContent>

                    <CardFooter class="flex items-center justify-end gap-2">
                        <Button size="sm" variant="outline" @click="$router.push(`/appliances/${appliance.id}`)">
                            View details
                        </Button>
                    </CardFooter>
                </Card>
            </div>
        </div>
    </div>
</template>