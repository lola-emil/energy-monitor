<template>

    <main class="px-5 mt-3">
        <div class="my-3 flex justify-between items-end">
            <div class="flex items-center gap-1">
                <input type="datetime-local" class="input" v-model="dateTime" />

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
                            1,284 kWh
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
                                233 V
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
                            <p class="text-lg font-semibold">50 kWh</p>
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
                            <p class="text-lg font-semibold">233 A</p>
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
                            <p class="text-lg font-semibold">233 Hz</p>
                        </div>
                    </CardContent>
                </Card>
            </div>
        </section>

        <section class="mt-5 min-h-116 grid grid-cols-2 gap-5">
            <Card>
                <CardHeader>
                    <CardTitle>
                        Energy Usage (Monthly)
                    </CardTitle>
                </CardHeader>

                <CardContent class="p-0 flex-1">
                    <div class="h-96">
                        <EnergyUsageChart />
                    </div>
                </CardContent>
            </Card>

            <Card>
                <CardHeader>
                    <CardTitle>Devices</CardTitle>
                </CardHeader>

                <CardContent>

                </CardContent>
            </Card>
        </section>

    </main>

    <br>
</template>


<script setup lang="ts">
import { Leaf, Plug, Zap, Atom, Radio, RotateCw, Printer } from "lucide-vue-next";
import { ref } from "vue";

import { Button } from '@/components/ui/button';
import {
    Card, CardContent, CardHeader,
    CardTitle
} from '@/components/ui/card';
import EnergyUsageChart from "./components/EnergyUsageChart.vue";


function getNowDatetimeLocal() {
    const now = new Date()
    const year = now.getFullYear()
    const month = String(now.getMonth() + 1).padStart(2, '0')
    const day = String(now.getDate()).padStart(2, '0')
    const hours = String(now.getHours()).padStart(2, '0')
    const minutes = String(now.getMinutes()).padStart(2, '0')
    return `${year}-${month}-${day}T${hours}:${minutes}`
}

const dateTime = ref(getNowDatetimeLocal());


</script>
