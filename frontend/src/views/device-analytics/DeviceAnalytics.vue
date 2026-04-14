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
                            <h1 class="text-xl font-semibold tracking-tight">{{ device?.device_name }}</h1>
                            <p class="text-sm font-medium leading-none">Device Code: {{ device?.device_code }}</p>
                        </div>
                    </div>

                    <div>
                        <div class="flex flex-col gap-1">
                            <span>Power Draw: <span>{{ realTimeReading?.value.PowerKwh ?? 0.00 }} kWh</span></span>
                            <span>Current: <span>{{ realTimeReading?.value.Current ?? 0.00 }} A</span></span>
                        </div>
                    </div>
                </div>
            </CardContent>
        </Card>

        <section class="mt-5 grid grid-cols-1 gap-5">
            <Card>
                <CardHeader>
                    <CardTitle class="">
                        Power and Current
                    </CardTitle>
                </CardHeader>
                <CardContent>
                    <div class="h-96">
                        <PowerHistoryChart/>
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
    Card, CardContent, CardTitle, CardHeader
} from '@/components/ui/card';
import { useRoute } from 'vue-router';
import { MonitorSmartphone } from "lucide-vue-next";
import { axiosInstance } from '@/api/axios';
import { useAuthStore } from '@/stores/auth';
import { useSSEStore } from '@/stores/sseEvent';
import PowerHistoryChart from './components/PowerHistoryChart.vue';

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

const device = ref<Device | null>(null);
const realTimeReading = ref<DeviceEvent | null>(null);

const fetchDevice = () => {
    let paramsId = route.params.id;

    sse.connect(auth.userId + "");

    axiosInstance.get<Device>(`/api/devices/${paramsId}`, {
        headers: {
            "Authorization": `Bearer ${auth.token}`
        }
    })
        .then(res => {
            device.value = res.data
        }).catch(err => {
            console.log(err);
        })


}

function subscribeToDevice(id: string) {
    const eventName = `read-${id}`
    currentEvent = eventName

    sse.subscribe(eventName, (evt) => {
        realTimeReading.value = JSON.parse(evt.data);
        console.log('SSE:', evt.data)
    })
}

function cleanup() {
    if (currentEvent) {
        realTimeReading.value = null;
        sse.unsubscribe(currentEvent)
    }
}

onMounted(() => {
    fetchDevice();
    subscribeToDevice(route.params.id + "");
})

watch(
    () => route.params.id,
    () => {
        fetchDevice()

        cleanup();
        subscribeToDevice(route.params.id + "")
    }
)
</script>
