<template>

    <main class="px-5 mt-3">
        <div class="my-3 flex justify-between items-end">
            <div class="flex items-center gap-1">
                <!-- <NativeSelect class="w-max">
                    <NativeSelectOption>Daily</NativeSelectOption>
                    <NativeSelectOption>Weekly</NativeSelectOption>
                    <NativeSelectOption>Montly</NativeSelectOption>
                    <NativeSelectOption>Yearly</NativeSelectOption>
                </NativeSelect> -->

                <Input type="month" v-model="month" />

                <Button class="btn btn-ghost btn-square">
                    <RotateCw :size="20" />
                </Button>
            </div>
            <div>
                <Button class="btn btn-primary">
                    <Printer :size="19" /> Print Report
                </Button>
            </div>
        </div>
        <section class="grid grid-cols-3 min-h-96 gap-5">
            <Card class="col-span-3 lg:col-span-1">
                <CardContent class="p-6">
                    <div class="p-3 rounded-lg bg-primary text-white w-max">
                        <leaf />
                    </div>

                    <p class="text-xl mt-3">
                        Total Energy Consumed
                    </p>

                    <div>
                        <p class="text-3xl font-semibold">
                            {{ overview?.total_consumed.toFixed(2) }} kWh
                        </p>
                    </div>
                </CardContent>
            </Card>
            <div class="col-span-3 lg:col-span-2 grid grid-cols-2 gap-5">
                <Card>
                    <CardContent class="p-6">
                        <div class="p-3 bg-primary rounded-lg text-white w-max">
                            <Plug />
                        </div>

                        <p class="text-lg mt-3">
                            Average Voltage
                        </p>

                        <div>
                            <p class="text-lg font-semibold">
                                {{ overview?.avg_volt.toFixed(2) }} V
                            </p>
                        </div>
                    </CardContent>
                </Card>

                <Card>
                    <CardContent class="p-6">
                        <div class="p-3 bg-primary rounded-lg text-white w-max">
                            <Zap />
                        </div>

                        <p class="text-lg mt-3">Average Power Draw</p>


                        <div>
                            <p class="text-lg font-semibold">
                                {{ overview?.avg_power.toFixed(2) }} kWh</p>
                        </div>
                    </CardContent>
                </Card>

                <Card>
                    <CardContent class="p-6">
                        <div class="p-3 bg-primary rounded-lg text-white w-max">
                            <Atom />
                        </div>

                        <p class="text-lg mt-3">Average Current</p>


                        <div>
                            <p class="text-lg font-semibold">23 A</p>
                        </div>
                    </CardContent>
                </Card>



                <Card>
                    <CardContent class="p-6">
                        <div class="p-3 bg-primary rounded-lg text-white w-max">
                            <Radio />
                        </div>

                        <p class="text-lg mt-3">Frequency</p>


                        <div>
                            <p class="text-lg font-semibold">50 Hz</p>
                        </div>
                    </CardContent>
                </Card>
            </div>
        </section>

        <section class="mt-5 min-h-116 grid grid-cols-1 gap-5">
            <Card>
                <CardHeader>
                    <CardTitle>
                        Energy Usage (Monthly)
                    </CardTitle>
                </CardHeader>

                <CardContent class="p-0 flex-1">
                    <div class="h-96">
                        <EnergyUsageChart :data="graphData"/>
                    </div>
                </CardContent>
            </Card>

            <Card>
                <CardHeader>
                    <CardTitle>Alerts</CardTitle>
                </CardHeader>

                <CardContent>
                    <Empty>
                        <EmptyHeader>

                        </EmptyHeader>
                        <EmptyContent>
                            <EmptyMedia variant="icon">
                                <TriangleAlert />
                            </EmptyMedia>
                            <EmptyTitle>No Alerts Yet</EmptyTitle>
                            <EmptyDescription>
                                The system haven't detected any anomalies yet.
                            </EmptyDescription>
                        </EmptyContent>

                    </Empty>
                </CardContent>
            </Card>
        </section>

    </main>

    <br>
</template>


<script setup lang="ts">
import { Leaf, Plug, Zap, Atom, Radio, RotateCw, Printer } from "lucide-vue-next";
import { onMounted, ref, watch } from "vue";

import { Button } from '@/components/ui/button';
import {
    Card, CardContent, CardHeader,
    CardTitle
} from '@/components/ui/card';
import EnergyUsageChart from "./components/EnergyUsageChart.vue";
import { TriangleAlert } from 'lucide-vue-next';
import { Input } from "@/components/ui/input";
import {
    Empty,
    EmptyContent,
    EmptyDescription,
    EmptyHeader,
    EmptyMedia,
    EmptyTitle,
} from '@/components/ui/empty';
import { axiosInstance } from "@/api/axios";
import { useAuthStore } from "@/stores/auth";

interface Overview {
    total_consumed: number;
    avg_volt: number;
    avg_power: number;
    avg_curr: number;
    avg_freq: number;
}


type MonthlyEnergy = {
    month: string | Date;
    energy_kwh: number;
};

const auth = useAuthStore();

const month = ref();
const overview = ref<Overview | null>(null);
const monthlyEnergy = ref<MonthlyEnergy[]>([]);
const graphData = ref<number[]>([]);

function initData(month: string) {
    axiosInstance.get<Overview>("/api/dashboard/overview?month=" + month,
        {
            headers: {
                "Authorization": `Bearer ${auth.token}`
            }
        })
        .then(val => {
            overview.value = val.data;
        }).catch(err => {
            console.log(err)
        })
}

function initGraphData() {
    axiosInstance.get<MonthlyEnergy[]>("/api/dashboard/monthly-consumption", {
        headers: {
            "Authorization": `Bearer ${auth.token}`
        }
    }).then(val => {
        monthlyEnergy.value = val.data;
        graphData.value = mapMonthlyToChart(val.data).data;

        console.log(graphData.value)
    }).catch(err => {
        console.log(err)
    })
}


function mapMonthlyToChart(data: MonthlyEnergy[]) {
    const months = [
        'Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun',
        'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'
    ];

    const values = new Array(12).fill(0);

    data.forEach(d => {
        const date = new Date(d.month);
        const monthIndex = date.getMonth(); // 0 = Jan, 11 = Dec
        values[monthIndex] = d.energy_kwh;
    });

    return {
        labels: months,
        data: values
    };
}

watch([month], () => {
    console.log(month.value);
    initData(month.value + "-01");
})

onMounted(() => {
    const date = "2026-04-01"

    initData(date);
    
    initGraphData();


})

</script>
