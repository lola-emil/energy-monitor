<script setup lang="ts">
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
import { reactive, onMounted, ref } from "vue"
import { settingsService } from "@/services/settings.service"
import type { Settings } from "@/types/settings"

const errorMessage = ref("")
const successMessage = ref("")

// ----- State (mock, replace with API calls) -----
const isLoading = ref(false)
const isSaving = ref(false)

// Billing & cost
const currency = ref('PHP')

// Alerts
const enableVoltageAlerts = ref(true)

const enableCurrentAlerts = ref(true)


const form = ref<Settings>({
    currency: "PHP",
    rate_per_kwh: 12.5,
    fixed_monthly_charge: 150,

    default_analytics_range: "month",
    refresh_interval_seconds: 30,
    time_format: "24h",

    enable_voltage_alerts: true,
    over_voltage_threshold: 240,
    under_voltage_threshold: 200,

    enable_current_alerts: true,
    over_current_threshold: 15,

    enable_offline_alerts: true,
})

const fetchSettings = async () => {
    try {
        isLoading.value = true
        errorMessage.value = ""

        const data = await settingsService.get()

        form.value = data;

        console.log(form.value);

    } catch (error: any) {
        errorMessage.value =
            error?.response?.data?.message ||
            "Failed to load settings"
    } finally {
        isLoading.value = false
    }
}

onMounted(() => {
    fetchSettings()


})

const handleSave = async () => {
    try {
        isLoading.value = true
        errorMessage.value = ""
        successMessage.value = ""

        await settingsService.update(form.value)

        successMessage.value = "Settings updated successfully"
    } catch (error: any) {
        errorMessage.value =
            error?.response?.data?.message ||
            "Failed to save settings"
    } finally {
        isLoading.value = false
    }
}
</script>

<template>
    <div class="px-5 my-5">
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
                                <Input id="currency" v-model="form.currency" class="max-w-40" placeholder="PHP" />
                            </div>

                            <div class="space-y-1">
                                <Label for="rate">Rate per kWh</Label>
                                <div class="flex items-center gap-2">
                                    <span class="text-sm text-muted-foreground">{{ form.currency }}</span>
                                    <Input id="rate" v-model.number="form.rate_per_kwh" type="number" min="0"
                                        step="0.01" class="max-w-40" />
                                </div>
                                <p class="text-xs text-muted-foreground">
                                    Multiplied by total energy (kWh) to compute estimated bill.
                                </p>
                            </div>

                            <div class="space-y-1">
                                <Label for="fixed-charge">Fixed monthly charge (optional)</Label>
                                <div class="flex items-center gap-2">
                                    <span class="text-sm text-muted-foreground">{{ form.currency }}</span>
                                    <Input id="fixed-charge" v-model.number="form.fixed_monthly_charge" type="number"
                                        min="0" step="0.01" class="max-w-40" />
                                </div>
                                <p class="text-xs text-muted-foreground">
                                    Added once per month on top of energy charges.
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
                                <Switch v-model="form.enable_voltage_alerts" :checked="form.enable_voltage_alerts" />
                            </div>
                            <div class="grid gap-4 md:grid-cols-2"
                                :class="!enableVoltageAlerts ? 'opacity-50 pointer-events-none' : ''">
                                <div class="space-y-1">
                                    <Label for="over-voltage">Over-voltage threshold</Label>
                                    <div class="flex items-center gap-2">
                                        <Input id="over-voltage" v-model.number="form.over_voltage_threshold"
                                            type="number" min="0" class="max-w-35" />
                                        <span class="text-sm text-muted-foreground">V</span>
                                    </div>
                                </div>
                                <div class="space-y-1">
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
                                <Switch v-model="form.enable_current_alerts" :checked="form.enable_current_alerts" />
                            </div>
                            <div class="grid gap-4 md:grid-cols-2"
                                :class="!enableCurrentAlerts ? 'opacity-50 pointer-events-none' : ''">
                                <div class="space-y-1">
                                    <Label for="over-current">Over-current threshold</Label>
                                    <div class="flex items-center gap-2">
                                        <Input id="over-current" v-model.number="form.over_current_threshold"
                                            type="number" min="0" class="max-w-35" />
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
                            <Switch v-model="form.enable_offline_alerts" :checked="form.enable_offline_alerts" />
                        </div>
                    </CardContent>
                </Card>

                <!-- Save button -->
                <div class="flex justify-end">
                    <Button @click="handleSave" :disabled="isLoading">
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
    </div>
</template>