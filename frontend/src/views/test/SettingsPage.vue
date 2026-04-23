<script setup lang="ts">
import { ref, onMounted } from 'vue'
import {
    Card,
    CardHeader,
    CardTitle,
    CardDescription,
    CardContent,
} from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Switch } from '@/components/ui/switch'
import { Separator } from '@/components/ui/separator'
import { Skeleton } from '@/components/ui/skeleton'

// ----- State (mock, replace with API calls) -----
const isLoading = ref(true)
const isSaving = ref(false)

// Billing & cost
const currency = ref('PHP')
const ratePerKwh = ref(10.25)
const fixedMonthlyCharge = ref(0)

// Display & refresh
const defaultAnalyticsRange = ref<'today' | '7d' | 'month'>('month')
const refreshIntervalSeconds = ref(5)
const timeFormat = ref<'24h' | '12h'>('24h')

// Alerts
const enableVoltageAlerts = ref(true)
const overVoltageThreshold = ref(250)
const underVoltageThreshold = ref(190)

const enableCurrentAlerts = ref(true)
const overCurrentThreshold = ref(15) // amps

const enableOfflineAlerts = ref(true)

onMounted(() => {
    // TODO: load settings from backend
    setTimeout(() => {
        isLoading.value = false
    }, 500)
})

const handleSave = async () => {
    isSaving.value = true
    try {
        // TODO: send settings to backend via API
        await new Promise((r) => setTimeout(r, 700))
    } finally {
        isSaving.value = false
    }
}
</script>

<template>
    <div class="flex flex-col gap-6">
        <div class="flex flex-col gap-2">
            <h1 class="text-2xl font-semibold tracking-tight">Settings</h1>
            <p class="text-sm text-muted-foreground">
                Configure billing, display preferences, and alert thresholds for your
                energy monitoring system.
            </p>
        </div>

        <form @submit.prevent="handleSave" class="space-y-6">
            <!-- Billing & Energy Cost -->
            <Card>
                <CardHeader>
                    <CardTitle>Billing & Energy Cost</CardTitle>
                    <CardDescription>
                        Used to estimate your monthly energy cost on the Overview page.
                    </CardDescription>
                </CardHeader>
                <CardContent class="space-y-4">
                    <div class="grid gap-4 md:grid-cols-3">
                        <div class="space-y-1">
                            <Label for="currency">Currency</Label>
                            <Input id="currency" v-model="currency" class="max-w-40" placeholder="PHP" />
                        </div>

                        <div class="space-y-1">
                            <Label for="rate">Rate per kWh</Label>
                            <div class="flex items-center gap-2">
                                <span class="text-sm text-muted-foreground">{{ currency }}</span>
                                <Input id="rate" v-model.number="ratePerKwh" type="number" min="0" step="0.01"
                                    class="max-w-40" />
                            </div>
                            <p class="text-xs text-muted-foreground">
                                Multiplied by total energy (kWh) to compute estimated bill.
                            </p>
                        </div>

                        <div class="space-y-1">
                            <Label for="fixed-charge">Fixed monthly charge (optional)</Label>
                            <div class="flex items-center gap-2">
                                <span class="text-sm text-muted-foreground">{{ currency }}</span>
                                <Input id="fixed-charge" v-model.number="fixedMonthlyCharge" type="number" min="0"
                                    step="0.01" class="max-w-40" />
                            </div>
                            <p class="text-xs text-muted-foreground">
                                Added once per month on top of energy charges.
                            </p>
                        </div>
                    </div>
                </CardContent>
            </Card>

            <!-- Display & Refresh -->
            <Card>
                <CardHeader>
                    <CardTitle>Display & Refresh</CardTitle>
                    <CardDescription>
                        Control how data is presented in the Analytics and Appliance pages.
                    </CardDescription>
                </CardHeader>
                <CardContent class="space-y-4">
                    <div class="grid gap-4 md:grid-cols-3">
                        <div class="space-y-1">
                            <Label>Default analytics range</Label>
                            <div class="inline-flex rounded-md border bg-background p-0.5">
                                <Button type="button" size="sm" variant="ghost" :class="defaultAnalyticsRange === 'today'
                                    ? 'bg-primary text-primary-foreground'
                                    : ''" @click="defaultAnalyticsRange = 'today'">
                                    Today
                                </Button>
                                <Button type="button" size="sm" variant="ghost" :class="defaultAnalyticsRange === '7d'
                                    ? 'bg-primary text-primary-foreground'
                                    : ''" @click="defaultAnalyticsRange = '7d'">
                                    Last 7 days
                                </Button>
                                <Button type="button" size="sm" variant="ghost" :class="defaultAnalyticsRange === 'month'
                                    ? 'bg-primary text-primary-foreground'
                                    : ''" @click="defaultAnalyticsRange = 'month'">
                                    This month
                                </Button>
                            </div>
                            <p class="text-xs text-muted-foreground">
                                Initial time range applied when opening the Analytics page.
                            </p>
                        </div>

                        <div class="space-y-1">
                            <Label for="refresh-interval">Real-time refresh interval</Label>
                            <div class="flex items-center gap-2">
                                <Input id="refresh-interval" v-model.number="refreshIntervalSeconds" type="number"
                                    min="1" step="1" class="max-w-30" />
                                <span class="text-sm text-muted-foreground">seconds</span>
                            </div>
                            <p class="text-xs text-muted-foreground">
                                How often the dashboard requests new readings for live charts.
                            </p>
                        </div>

                        <div class="space-y-1">
                            <Label>Time format</Label>
                            <div class="inline-flex rounded-md border bg-background p-0.5">
                                <Button type="button" size="sm" variant="ghost" :class="timeFormat === '24h'
                                    ? 'bg-primary text-primary-foreground'
                                    : ''" @click="timeFormat = '24h'">
                                    24-hour
                                </Button>
                                <Button type="button" size="sm" variant="ghost" :class="timeFormat === '12h'
                                    ? 'bg-primary text-primary-foreground'
                                    : ''" @click="timeFormat = '12h'">
                                    12-hour
                                </Button>
                            </div>
                            <p class="text-xs text-muted-foreground">
                                Used when displaying timestamps on charts and tables.
                            </p>
                        </div>
                    </div>
                </CardContent>
            </Card>

            <!-- Alerts & Thresholds -->
            <Card>
                <CardHeader>
                    <CardTitle>Alerts & Thresholds</CardTitle>
                    <CardDescription>
                        Configure when the system should raise alerts for abnormal readings.
                    </CardDescription>
                </CardHeader>
                <CardContent class="space-y-6">
                    <!-- Voltage alerts -->
                    <div class="space-y-3">
                        <div class="flex items-center justify-between">
                            <div>
                                <p class="text-sm font-medium">Voltage alerts</p>
                                <p class="text-xs text-muted-foreground">
                                    Trigger alerts when voltage goes beyond safe limits.
                                </p>
                            </div>
                            <Switch v-model:checked="enableVoltageAlerts" />
                        </div>
                        <div class="grid gap-4 md:grid-cols-2"
                            :class="!enableVoltageAlerts ? 'opacity-50 pointer-events-none' : ''">
                            <div class="space-y-1">
                                <Label for="over-voltage">Over-voltage threshold</Label>
                                <div class="flex items-center gap-2">
                                    <Input id="over-voltage" v-model.number="overVoltageThreshold" type="number" min="0"
                                        class="max-w-35" />
                                    <span class="text-sm text-muted-foreground">V</span>
                                </div>
                            </div>
                            <div class="space-y-1">
                                <Label for="under-voltage">Under-voltage threshold</Label>
                                <div class="flex items-center gap-2">
                                    <Input id="under-voltage" v-model.number="underVoltageThreshold" type="number"
                                        min="0" class="max-w-35" />
                                    <span class="text-sm text-muted-foreground">V</span>
                                </div>
                            </div>
                        </div>
                    </div>

                    <Separator />

                    <!-- Current / power alerts -->
                    <div class="space-y-3">
                        <div class="flex items-center justify-between">
                            <div>
                                <p class="text-sm font-medium">Current / load alerts</p>
                                <p class="text-xs text-muted-foreground">
                                    Trigger alerts when an appliance draws more current than
                                    expected.
                                </p>
                            </div>
                            <Switch v-model:checked="enableCurrentAlerts" />
                        </div>
                        <div class="grid gap-4 md:grid-cols-2"
                            :class="!enableCurrentAlerts ? 'opacity-50 pointer-events-none' : ''">
                            <div class="space-y-1">
                                <Label for="over-current">Over-current threshold</Label>
                                <div class="flex items-center gap-2">
                                    <Input id="over-current" v-model.number="overCurrentThreshold" type="number" min="0"
                                        class="max-w-35" />
                                    <span class="text-sm text-muted-foreground">A</span>
                                </div>
                                <p class="text-xs text-muted-foreground">
                                    You can later convert this to power if needed (P = V × I).
                                </p>
                            </div>
                        </div>
                    </div>

                    <Separator />

                    <!-- Offline alerts -->
                    <div class="flex items-center justify-between">
                        <div>
                            <p class="text-sm font-medium">Offline appliance alerts</p>
                            <p class="text-xs text-muted-foreground">
                                Notify when an appliance stops sending data for a set period.
                            </p>
                        </div>
                        <Switch v-model:checked="enableOfflineAlerts" />
                    </div>
                </CardContent>
            </Card>

            <!-- Save button -->
            <div class="flex justify-end">
                <Button type="submit" :disabled="isSaving || isLoading">
                    <span v-if="isSaving">Saving…</span>
                    <span v-else>Save settings</span>
                </Button>
            </div>
        </form>

        <!-- Optional loading overlay -->
        <div v-if="isLoading"
            class="pointer-events-none fixed inset-0 z-10 flex items-start justify-center bg-background/40 backdrop-blur-sm">
            <div class="mt-24 rounded-lg border bg-background px-6 py-4 shadow">
                <div class="flex items-center gap-3">
                    <Skeleton class="h-8 w-8 rounded-full" />
                    <div class="space-y-2">
                        <Skeleton class="h-3 w-40" />
                        <Skeleton class="h-3 w-24" />
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>