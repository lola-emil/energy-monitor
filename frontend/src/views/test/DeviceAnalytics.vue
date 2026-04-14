<template>
    <main class="px-5 mt-3">
        <Card>
            <CardContent>
                <div class="w-full flex items-center justify-between">
                    <div class="flex gap-5">
                        <div class="p-3 bg-primary rounded-lg text-white w-max">
                            <MonitorSmartphone />
                        </div>
                        <div class="flex flex-col gap-1">
                            <h1 class="text-xl font-semibold tracking-tight">PC</h1>
                            <p class="text-sm font-medium leading-none">Device Code: ESP-MADAFAK</p>
                        </div>
                    </div>

                    <div>
                        <div class="flex flex-col gap-1">
                            <span>Power Draw: <span>25 kWh</span></span>
                            <span>Current: <span>50 A</span></span>
                        </div>
                    </div>
                </div>
            </CardContent>
        </Card>

        <section class="mt-5 grid grid-cols-1 gap-5">
            <Card>
                <CardHeader>
                    <CardTitle class="">
                        Power Monitor
                    </CardTitle>
                </CardHeader>
                <CardContent>
                    <div class="h-96">
                        <AnalyticsEnergyUsage />
                    </div>
                </CardContent>
            </Card>

            <div class="grid grid-cols-2 gap-5 h-52">
                <Card>
                    <CardHeader>
                        <CardTitle class="">
                            Energy Usage (KWh)
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                    </CardContent>
                </Card>


                <Card>
                    <CardContent>
                        <CardTitle class="">
                            Voltage
                        </CardTitle>
                    </CardContent>
                </Card>
            </div>

        </section>
        <br>
    </main>
</template>

<script setup lang="ts">
import { onMounted, ref, watch } from 'vue';
import {
    Card, CardContent, CardTitle
} from '@/components/ui/card';
import { useRoute } from 'vue-router';
import { MonitorSmartphone } from "lucide-vue-next";
import { axiosInstance } from '@/api/axios';
import { useAuthStore } from '@/stores/auth';
import { useSSEStore } from '@/stores/sseEvent';
import AnalyticsEnergyUsage from './components/AnalyticsEnergyUsage.vue';
import CardHeader from '@/components/ui/card/CardHeader.vue';

interface Device {
    device_code: string
    created_at: string
    device_id: number
    device_name: string
    id: number, is_active: boolean
    last_active: string | null
    user_id: number
}



export interface DeviceValue {
    DeviceId: number;
    Voltage: string;
    Current: string;
    PowerKwh: string;
}

export interface DeviceEvent {
    user_id: number;
    device_id: number;
    value: DeviceValue;
}


const route = useRoute();
const auth = useAuthStore();

const sse = useSSEStore();
let currentEvent: string | null = null


</script>
