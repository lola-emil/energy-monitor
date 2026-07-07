<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
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
import { settingsService } from '@/services/settings.service'
import type { Settings } from '@/types/settings'
import {
    CircleDollarSignIcon,
    BellRingIcon,
    ShieldAlertIcon,
    WifiOffIcon,
    SaveIcon,
} from 'lucide-vue-next'

const errorMessage = ref('')
const successMessage = ref('')
const isLoading = ref(true)
const isSaving = ref(false)

const form = ref<Settings>({
    currency: 'PHP',
    rate_per_kwh: 12.5,
    fixed_monthly_charge: 150,
    default_analytics_range: 'month',
    refresh_interval_seconds: 30,
    time_format: '24h',
    enable_voltage_alerts: true,
    over_voltage_threshold: 240,
    under_voltage_threshold: 200,
    enable_current_alerts: true,
    over_current_threshold: 15,
    enable_offline_alerts: true,
})

const voltageAlertsEnabled = computed(() => form.value.enable_voltage_alerts)
const currentAlertsEnabled = computed(() => form.value.enable_current_alerts)

const fetchSettings = async () => {
    try {
        isLoading.value = true
        errorMessage.value = ''
        successMessage.value = ''

        const data = await settingsService.get()
        form.value = data
    } catch (error: any) {
        console.error(error)
        errorMessage.value =
            error?.response?.data?.message || 'Failed to load settings.'
    } finally {
        isLoading.value = false
    }
}

const handleSave = async () => {
    try {
        isSaving.value = true
        errorMessage.value = ''
        successMessage.value = ''

        await settingsService.update(form.value)
        successMessage.value = 'Settings updated successfully.'
    } catch (error: any) {
        console.error(error)
        errorMessage.value =
            error?.response?.data?.message || 'Failed to save settings.'
    } finally {
        isSaving.value = false
    }
}

onMounted(fetchSettings)
</script>

<template>
    <div class="min-h-screen bg-muted/20">
        <div class="mx-auto w-full space-y-6 px-4 py-6 sm:px-6 lg:px-8">
            <!-- Header -->
            <section class="rounded-2xl border bg-background px-5 py-5 shadow-sm">
                <div class="space-y-1">
                    <h1 class="text-2xl font-semibold tracking-tight">Settings</h1>
                    <p class="text-sm text-muted-foreground">
                        Configure billing preferences and alert thresholds for your energy monitoring system.
                    </p>
                </div>
            </section>

            <!-- Feedback -->
            <div v-if="errorMessage"
                class="rounded-xl border border-destructive/30 bg-destructive/5 px-4 py-3 text-sm text-destructive">
                {{ errorMessage }}
            </div>

            <div v-if="successMessage"
                class="rounded-xl border border-emerald-500/30 bg-emerald-500/5 px-4 py-3 text-sm text-emerald-700 dark:text-emerald-400">
                {{ successMessage }}
            </div>

            <!-- Loading state -->
            <section v-if="isLoading" class="space-y-6">
                <Card class="rounded-2xl shadow-sm">
                    <CardHeader class="space-y-2">
                        <Skeleton class="h-5 w-40" />
                        <Skeleton class="h-4 w-72" />
                    </CardHeader>
                    <CardContent class="grid gap-4 md:grid-cols-3">
                        <Skeleton class="h-20 w-full rounded-xl" />
                        <Skeleton class="h-20 w-full rounded-xl" />
                        <Skeleton class="h-20 w-full rounded-xl" />
                    </CardContent>
                </Card>

                <Card class="rounded-2xl shadow-sm">
                    <CardHeader class="space-y-2">
                        <Skeleton class="h-5 w-40" />
                        <Skeleton class="h-4 w-80" />
                    </CardHeader>
                    <CardContent class="space-y-5">
                        <Skeleton class="h-16 w-full rounded-xl" />
                        <Skeleton class="h-24 w-full rounded-xl" />
                        <Skeleton class="h-16 w-full rounded-xl" />
                    </CardContent>
                </Card>
            </section>

            <!-- Form -->
            <form v-else @submit.prevent="handleSave" class="space-y-6">
                <!-- Billing -->
                <Card class="rounded-2xl shadow-sm">
                    <CardHeader class="border-b pb-4">
                        <div class="flex items-start gap-3">
                            <div class="rounded-xl bg-primary/10 p-2 text-primary">
                                <CircleDollarSignIcon class="h-4 w-4" />
                            </div>
                            <div>
                                <CardTitle>Billing & Energy Cost</CardTitle>
                                <CardDescription>
                                    These values are used to estimate electricity costs throughout the dashboard.
                                </CardDescription>
                            </div>
                        </div>
                    </CardHeader>

                    <CardContent class="grid gap-5 p-5 md:grid-cols-3">
                        <div class="space-y-2">
                            <Label for="currency">Currency</Label>
                            <Input id="currency" v-model="form.currency" class="max-w-40" placeholder="PHP" />
                            <p class="text-xs text-muted-foreground">
                                Used when displaying estimated bills and charges.
                            </p>
                        </div>

                        <div class="space-y-2">
                            <Label for="rate">Rate per kWh</Label>
                            <div class="flex items-center gap-2">
                                <span class="text-sm text-muted-foreground">{{ form.currency }}</span>
                                <Input id="rate" v-model.number="form.rate_per_kwh" type="number" min="0" step="0.01"
                                    class="max-w-40" />
                            </div>
                            <p class="text-xs text-muted-foreground">
                                Multiplied by total energy consumption to estimate cost.
                            </p>
                        </div>

                        <div class="space-y-2">
                            <Label for="fixed-charge">Fixed monthly charge</Label>
                            <div class="flex items-center gap-2">
                                <span class="text-sm text-muted-foreground">{{ form.currency }}</span>
                                <Input id="fixed-charge" v-model.number="form.fixed_monthly_charge" type="number"
                                    min="0" step="0.01" class="max-w-40" />
                            </div>
                            <p class="text-xs text-muted-foreground">
                                Optional recurring charge added once per billing cycle.
                            </p>
                        </div>
                    </CardContent>
                </Card>

                <!-- Alerts -->
                <Card class="rounded-2xl shadow-sm">
                    <CardHeader class="border-b pb-4">
                        <div class="flex items-start gap-3">
                            <div class="rounded-xl bg-primary/10 p-2 text-primary">
                                <BellRingIcon class="h-4 w-4" />
                            </div>
                            <div>
                                <CardTitle>Alerts & Thresholds</CardTitle>
                                <CardDescription>
                                    Control which conditions generate alerts and define the acceptable limits.
                                </CardDescription>
                            </div>
                        </div>
                    </CardHeader>

                    <CardContent class="space-y-6 p-5">
                        <!-- Voltage -->
                        <div class="space-y-4">
                            <div class="flex items-start justify-between gap-4">
                                <div class="space-y-1">
                                    <div class="flex items-center gap-2">
                                        <ShieldAlertIcon class="h-4 w-4 text-primary" />
                                        <p class="text-sm font-medium">Voltage alerts</p>
                                    </div>
                                    <p class="text-xs text-muted-foreground">
                                        Notify when voltage rises above or falls below safe operating limits.
                                    </p>
                                </div>

                                <Switch v-model="form.enable_voltage_alerts" />
                            </div>

                            <div class="grid gap-4 rounded-xl border bg-muted/20 p-4 md:grid-cols-2"
                                :class="!voltageAlertsEnabled ? 'pointer-events-none opacity-50' : ''">
                                <div class="space-y-2">
                                    <Label for="over-voltage">Over-voltage threshold</Label>
                                    <div class="flex items-center gap-2">
                                        <Input id="over-voltage" v-model.number="form.over_voltage_threshold"
                                            type="number" min="0" class="max-w-35" />
                                        <span class="text-sm text-muted-foreground">V</span>
                                    </div>
                                </div>

                                <div class="space-y-2">
                                    <Label for="under-voltage">Under-voltage threshold</Label>
                                    <div class="flex items-center gap-2">
                                        <Input id="under-voltage" v-model.number="form.under_voltage_threshold"
                                            type="number" min="0" class="max-w-35" />
                                        <span class="text-sm text-muted-foreground">V</span>
                                    </div>
                                </div>
                            </div>
                        </div>

                        <Separator />

                        <!-- Current -->
                        <div class="space-y-4">
                            <div class="flex items-start justify-between gap-4">
                                <div class="space-y-1">
                                    <div class="flex items-center gap-2">
                                        <BellRingIcon class="h-4 w-4 text-primary" />
                                        <p class="text-sm font-medium">Current / load alerts</p>
                                    </div>
                                    <p class="text-xs text-muted-foreground">
                                        Notify when current draw exceeds the configured threshold.
                                    </p>
                                </div>

                                <Switch v-model="form.enable_current_alerts" />
                            </div>

                            <div class="rounded-xl border bg-muted/20 p-4"
                                :class="!currentAlertsEnabled ? 'pointer-events-none opacity-50' : ''">
                                <div class="space-y-2">
                                    <Label for="over-current">Over-current threshold</Label>
                                    <div class="flex items-center gap-2">
                                        <Input id="over-current" v-model.number="form.over_current_threshold"
                                            type="number" min="0" class="max-w-35" />
                                        <span class="text-sm text-muted-foreground">A</span>
                                    </div>
                                    <p class="text-xs text-muted-foreground">
                                        Useful for detecting overloads and abnormal appliance behavior.
                                    </p>
                                </div>
                            </div>
                        </div>

                        <Separator />

                        <!-- Offline -->
                        <!-- <div class="flex items-start justify-between gap-4">
                            <div class="space-y-1">
                                <div class="flex items-center gap-2">
                                    <WifiOffIcon class="h-4 w-4 text-primary" />
                                    <p class="text-sm font-medium">Offline appliance alerts</p>
                                </div>
                                <p class="text-xs text-muted-foreground">
                                    Notify when an appliance stops reporting data for an extended period.
                                </p>
                            </div>

                            <Switch v-model="form.enable_offline_alerts" />
                        </div> -->
                    </CardContent>
                </Card>

                <!-- Actions -->
                <div class="flex justify-end">
                    <Button type="submit" class="gap-2 rounded-xl" :disabled="isSaving">
                        <SaveIcon class="h-4 w-4" />
                        <span v-if="isSaving">Saving...</span>
                        <span v-else>Save settings</span>
                    </Button>
                </div>
            </form>
        </div>
    </div>
</template>