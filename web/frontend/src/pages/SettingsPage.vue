<template>
    <div class="min-h-screen bg-base-200 flex flex-col">
        <Navbar />
        <main class="flex-1 p-4 md:p-8">
            <div class="space-y-6">

                <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
                    <div>
                        <h2 class="text-2xl font-bold text-base-content">Settings</h2>
                        <p class="text-base-content/60 text-sm mt-1">Configure billing, alerts, and energy
                            monitoring preferences</p>
                    </div>
                    <div class="flex gap-3">
                        <button class="btn btn-ghost">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24"
                                stroke="currentColor">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                    d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
                            </svg>
                            Export Settings
                        </button>
                        <button class="btn btn-primary" @click="saveSettings">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24"
                                stroke="currentColor">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                    d="M8 7H5a2 2 0 00-2 2v9a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-3m-1 4l-4 4m0 0l-4-4m4 4V3" />
                            </svg>
                            Save Settings
                        </button>
                    </div>
                </div>

                <div class="card bg-base-100 shadow-sm border border-base-300">
                    <div class="card-body p-6">
                        <h3 class="card-title text-xl text-base-content mb-6">Billing & Energy Cost</h3>

                        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">

                            <div class="form-control">
                                <label class="label">
                                    <span class="label-text font-medium">Currency</span>
                                </label>
                                <select v-model="settings.currency" class="select select-bordered w-full">
                                    <option value="usd">US Dollar (USD)</option>
                                    <option value="eur">Euro (EUR)</option>
                                    <option value="gbp">British Pound (GBP)</option>
                                    <option value="cad">Canadian Dollar (CAD)</option>
                                    <option value="aud">Australian Dollar (AUD)</option>
                                </select>
                            </div>

                            <div class="form-control">
                                <label class="label">
                                    <span class="label-text font-medium">Rate per kWh</span>
                                    <span class="label-text-alt text-base-content/50">Your current utility
                                        rate</span>
                                </label>
                                <div class="flex">
                                    <label class="input w-full">
                                        <input type="number" v-model="settings.ratePerKwh" step="0.01" min="0"
                                            class="input input-bordered w-full" placeholder="0.15" />
                                        <span class="badge badge-neutral badge-xs">Per kWh</span>
                                    </label>
                                </div>
                            </div>

                            <div class="form-control">
                                <label class="label">
                                    <span class="label-text font-medium">Fixed Monthly Charge</span>
                                    <span class="label-text-alt text-base-content/50">Service fees or base
                                        charges</span>
                                </label>
                                <div class="flex">

                                    <input type="number" v-model="settings.fixedMonthlyCharge" step="0.01" min="0"
                                        class="input input-bordered w-full " placeholder="12.50" />
                                </div>
                            </div>
                        </div>

                        <div class="mt-8 p-5 bg-base-200 rounded-lg border border-base-300">
                            <h4 class="font-medium text-base-content mb-3">Current Billing Summary</h4>
                            <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
                                <div class="bg-base-100 p-3 rounded-lg">
                                    <p class="text-sm text-base-content/60">Estimated Monthly Cost</p>
                                    <p class="text-2xl font-bold text-base-content">${{ monthlyCost }}</p>
                                </div>
                                <div class="bg-base-100 p-3 rounded-lg">
                                    <p class="text-sm text-base-content/60">Current Rate</p>
                                    <p class="text-2xl font-bold text-base-content">${{ settings.ratePerKwh }} / kWh
                                    </p>
                                </div>
                                <div class="bg-base-100 p-3 rounded-lg">
                                    <p class="text-sm text-base-content/60">Fixed Charge</p>
                                    <p class="text-2xl font-bold text-base-content">${{ settings.fixedMonthlyCharge
                                        }}</p>
                                </div>
                                <div class="bg-base-100 p-3 rounded-lg">
                                    <p class="text-sm text-base-content/60">Total Devices</p>
                                    <p class="text-2xl font-bold text-base-content">{{ totalDevices }}</p>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>

                <div class="card bg-base-100 shadow-sm border border-base-300">
                    <div class="card-body p-6">
                        <h3 class="card-title text-xl text-base-content mb-6">Alerts & Thresholds</h3>

                        <div class="grid grid-cols-1 md:grid-cols-2 gap-8">

                            <div>
                                <h4 class="font-medium text-base-content mb-4">Voltage Alerts</h4>

                                <div class="form-control mb-4">
                                    <label class="label">
                                        <span class="label-text font-medium">Over Voltage Threshold</span>
                                        <span class="label-text-alt text-base-content/50">Trigger alert when voltage
                                            exceeds this value</span>
                                    </label>
                                    <div class="flex">
                                        <input type="number" v-model="settings.overVoltageThreshold"
                                            class="input input-bordered w-full" min="100" max="300" placeholder="250" />
                                    </div>
                                </div>

                                <div class="form-control">
                                    <label class="label">
                                        <span class="label-text font-medium">Under Voltage Threshold</span>
                                        <span class="label-text-alt text-base-content/50">Trigger alert when voltage
                                            falls below this value</span>
                                    </label>
                                    <div class="flex">
                                        <input type="number" v-model="settings.underVoltageThreshold"
                                            class="input input-bordered w-full" min="100" max="300" placeholder="100" />
                                    </div>
                                </div>
                            </div>

                            <div>
                                <h4 class="font-medium text-base-content mb-4">Device Alerts</h4>

                                <div class="card bg-base-200 border border-base-300 rounded-lg p-4 mb-4">
                                    <div class="flex items-center justify-between">
                                        <div>
                                            <p class="font-medium text-base-content">Device Offline Alert</p>
                                            <p class="text-sm text-base-content/60">Notify when a device stops
                                                reporting</p>
                                        </div>
                                        <label class="swap swap-rotate">
                                            <input type="checkbox" v-model="settings.deviceOfflineAlert"
                                                class="theme-controller" />
                                            <svg class="swap-off fill-current w-8 h-8 text-error"
                                                xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                                                <path
                                                    d="M12,22A10,10,0,1,0,2,12,10,10,0,0,12,22Zm0-2a8,8,0,1,1,8-8A8,8,0,0,1,12,20Z" />
                                            </svg>
                                            <svg class="swap-on fill-current w-8 h-8 text-success"
                                                xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                                                <path
                                                    d="M12,22A10,10,0,1,0,2,12,10,10,0,0,12,22Zm0-2a8,8,0,1,1,8-8A8,8,0,0,1,12,20Z" />
                                            </svg>
                                        </label>
                                    </div>
                                </div>

                                <div class="card bg-base-200 border border-base-300 rounded-lg p-4">
                                    <div class="flex items-center justify-between">
                                        <div>
                                            <p class="font-medium text-base-content">Offline Duration Threshold</p>
                                            <p class="text-sm text-base-content/60">How long before triggering alert
                                            </p>
                                        </div>
                                        <select v-model="settings.offlineDuration" class="select select-bordered w-32">
                                            <option value="5">5 minutes</option>
                                            <option value="15">15 minutes</option>
                                            <option value="30">30 minutes</option>
                                            <option value="60">1 hour</option>
                                        </select>
                                    </div>
                                </div>
                            </div>

                        </div>

                        <div class="mt-8">
                            <h4 class="font-medium text-base-content mb-3">Recent Alerts</h4>
                            <div class="space-y-3">
                                <div
                                    class="flex items-center justify-between p-3 bg-base-200 rounded-lg border border-base-300">
                                    <div class="flex items-center gap-3">
                                        <div class="bg-error/20 text-error p-2 rounded-lg">
                                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none"
                                                viewBox="0 0 24 24" stroke="currentColor">
                                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                                    d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                                            </svg>
                                        </div>
                                        <div>
                                            <p class="font-medium text-base-content">Over Voltage Alert</p>
                                            <p class="text-xs text-base-content/60">Kitchen Fridge (DEV-9932) - 255V
                                            </p>
                                        </div>
                                    </div>
                                    <span class="text-xs text-base-content/50">2 hours ago</span>
                                </div>

                                <div
                                    class="flex items-center justify-between p-3 bg-base-200 rounded-lg border border-base-300">
                                    <div class="flex items-center gap-3">
                                        <div class="bg-warning/20 text-warning p-2 rounded-lg">
                                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none"
                                                viewBox="0 0 24 24" stroke="currentColor">
                                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                                    d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                                            </svg>
                                        </div>
                                        <div>
                                            <p class="font-medium text-base-content">Device Offline Alert</p>
                                            <p class="text-xs text-base-content/60">Office Lights (DEV-1043) - 35
                                                minutes</p>
                                        </div>
                                    </div>
                                    <span class="text-xs text-base-content/50">1 day ago</span>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>

            </div>
        </main>

    </div>
</template>

<script setup lang="ts">
import Navbar from '@/components/Navbar.vue';
import { ref, computed } from 'vue';

const settings = ref({
    currency: 'usd',
    ratePerKwh: 0.15,
    fixedMonthlyCharge: 12.50,
    overVoltageThreshold: 250,
    underVoltageThreshold: 100,
    deviceOfflineAlert: true,
    offlineDuration: '15'
});

const monthlyCost = computed(() => {
    return ((32.5 * 30) * settings.value.ratePerKwh + settings.value.fixedMonthlyCharge).toFixed(2);
});

const totalDevices = computed(() => {
    return 14;
});

const saveSettings = () => {
    console.log('Saving settings:', settings.value);
    alert('Settings saved successfully!');
};
</script>