<script setup lang="ts">
import Navbar from "@/components/Navbar.vue";
import useEchart from "@/composables/useECharts";
import type { EChartsOption } from "echarts";
import { Leaf, Plug, Zap, Atom, Radio, RotateCw, Printer} from "lucide-vue-next"
import { onMounted, ref, useTemplateRef } from "vue";

const devices = [
    {
        device_name: "Office Router",
        device_id: "DEV-1001",
        mac_address: "00:1A:2B:3C:4D:01",
        status: "online"
    },
    {
        device_name: "Warehouse Sensor",
        device_id: "DEV-1002",
        mac_address: "00:1A:2B:3C:4D:02",
        status: "offline"
    },
    {
        device_name: "Lobby Camera",
        device_id: "DEV-1003",
        mac_address: "00:1A:2B:3C:4D:03",
        status: "online"
    },
    {
        device_name: "Conference Room Display",
        device_id: "DEV-1004",
        mac_address: "00:1A:2B:3C:4D:04",
        status: "offline"
    },
    {
        device_name: "Entrance Door Controller",
        device_id: "DEV-1005",
        mac_address: "00:1A:2B:3C:4D:05",
        status: "online"
    },
    {
        device_name: "Parking Lot Camera",
        device_id: "DEV-1006",
        mac_address: "00:1A:2B:3C:4D:06",
        status: "offline"
    },
    {
        device_name: "Server Room UPS",
        device_id: "DEV-1007",
        mac_address: "00:1A:2B:3C:4D:07",
        status: "online"
    },
    {
        device_name: "Smart Thermostat",
        device_id: "DEV-1008",
        mac_address: "00:1A:2B:3C:4D:08",
        status: "online"
    },
    {
        device_name: "Meeting Room Tablet",
        device_id: "DEV-1009",
        mac_address: "00:1A:2B:3C:4D:09",
        status: "offline"
    },
    {
        device_name: "Backup NAS",
        device_id: "DEV-1010",
        mac_address: "00:1A:2B:3C:4D:10",
        status: "online"
    }
];

const barChartEl = useTemplateRef<HTMLDivElement>("pie-chart");
const barChart = useEchart(barChartEl);


const lineChartOpt: EChartsOption = {
    backgroundColor: 'transparent',

    tooltip: {
        trigger: 'axis',
        backgroundColor: '#111827',
        textStyle: { color: '#fff' },
        formatter: (params: any) => {
            return params
                .map((p: any) => `${p.marker} ${p.seriesName}: ${p.value / 1000}K`)
                .join('<br/>');
        }
    },

    legend: {
        top: 10,
        right: 20,
        icon: 'roundRect',
        itemWidth: 18,
        itemHeight: 3,
        textStyle: {
            color: '#374151',
            fontSize: 13
        },
        data: ['Actual Value', 'Projected Value']
    },

    grid: {
        top: 60,
        left: 60,
        right: 30,
        bottom: 40
    },

    xAxis: {
        type: 'category',
        boundaryGap: false,
        data: [
            'Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun',
            'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'
        ],
        axisLine: { show: false },
        axisTick: { show: false },
        axisLabel: {
            color: '#6B7280'
        }
    },

    yAxis: {
        type: 'value',
        min: 0,
        max: 80000,
        interval: 20000,
        axisLine: { show: false },
        axisTick: { show: false },
        axisLabel: {
            color: '#6B7280',
            formatter: value => value === 0 ? '0' : `${value / 1000}K`
        },
        splitLine: {
            lineStyle: {
                color: '#E5E7EB',
                type: 'solid'
            }
        }
    },

    series: [
        {
            name: 'Actual Value',
            type: 'line',
            // smooth: true,
            data: [48000, 59000, 53000, 68000, 34000, 45000, 30000, 33000, 58000, 46000, 39000, 51000],
            symbol: 'circle',
            symbolSize: 6,
            lineStyle: {
                width: 2,
                color: '#3B82F6'
            },
            itemStyle: {
                color: '#3B82F6'
            },
            areaStyle: {
                color: {
                    type: 'linear',
                    x: 0,
                    y: 0,
                    x2: 0,
                    y2: 1,
                    colorStops: [
                        { offset: 0, color: 'rgba(59,130,246,0.25)' },
                        { offset: 1, color: 'rgba(59,130,246,0.05)' }
                    ]
                }
            }
        },
        {
            name: 'Projected Value',
            type: 'line',
            // smooth: true,
            data: [60000, 76000, 62000, 78000, 55000, 55000, 41000, 70000, 30000, 63000, 45000, 75000],
            symbol: 'none',
            lineStyle: {
                width: 2,
                type: 'dashed',
                color: '#10B981'
            }
        }
    ]
};

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

onMounted(() => {
    barChart.setOptions(lineChartOpt);
})
</script>

<template>
    <Navbar title="Dashboard & Analytics" />

    <main class="px-5 mt-3">
        <div class="my-3 flex justify-between items-end">
            <div class="flex items-center gap-1">
                <input type="datetime-local" class="input" v-model="dateTime"/>

                <button class="btn btn-ghost btn-square">
                    <RotateCw :size="20" />
                </button>
            </div>
            <div>
                <button class="btn btn-primary"><Printer :size="19"/> Print Report</button>
            </div>
        </div>
        <section class="grid grid-cols-3 min-h-96 gap-3">
            <div class="card bg-base-100 col-span-3 lg:col-span-1">
                <div class="card-body">
                    <div class="p-3 rounded-lg bg-primary text-white w-max">
                        <leaf />
                    </div>

                    <p class="text-xl mt-3">Total Energy Consumed</p>


                    <div>
                        <p class="text-3xl font-semibold">1,284 kWh</p>
                    </div>
                </div>
            </div>
            <div class="col-span-3 lg:col-span-2 grid grid-cols-2 gap-3">
                <div class="card bg-base-100">
                    <div class="card-body">
                        <div class="p-3 bg-primary rounded-lg text-white w-max">
                            <Plug />
                        </div>

                        <p class="text-lg mt-3">Average Voltage</p>


                        <div>
                            <p class="text-lg font-semibold">233 V</p>
                        </div>
                    </div>
                </div>
                <div class="card bg-base-100">
                    <div class="card-body">
                        <div class="p-3 bg-primary rounded-lg text-white w-max">
                            <Zap />
                        </div>

                        <p class="text-lg mt-3">Average Power Draw</p>


                        <div>
                            <p class="text-lg font-semibold">50 kWh</p>
                        </div>
                    </div>
                </div>

                <div class="card bg-base-100">
                    <div class="card-body">
                        <div class="p-3 bg-primary rounded-lg text-white w-max">
                            <Atom />
                        </div>

                        <p class="text-lg mt-3">Average Current</p>


                        <div>
                            <p class="text-lg font-semibold">233 A</p>
                        </div>
                    </div>
                </div>

                <div class="card bg-base-100">
                    <div class="card-body">
                        <div class="p-3 bg-primary rounded-lg text-white w-max">
                            <Radio />
                        </div>

                        <p class="text-lg mt-3">Frequency</p>


                        <div>
                            <p class="text-lg font-semibold">233 Hz</p>
                        </div>
                    </div>
                </div>

            </div>
        </section>

        <section class="mt-3 min-h-96 grid">
            <div class="card bg-base-100">
                <div class="card-body p-3 px-0">
                    <div ref="pie-chart" class="h-full"></div>
                </div>
            </div>
        </section>

        <section class="mt-3 min-h-125 grid">
            <div class="card bg-base-100">
                <div class="card-body">
                    <table class="table table-sm">
                        <thead>
                            <tr>
                                <th>Device Name</th>
                                <th>ID</th>
                                <th>Status</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="value in devices">
                                <td>{{ value.device_name }}</td>
                                <td>{{ value.device_id }}</td>
                                <td>
                                    <div
                                        :class="value.status == 'online' ? 'badge badge-soft badge-success' : 'badge badge-soft badge-error'">
                                        {{ value.status }}
                                    </div>
                                </td>
                            </tr>

                        </tbody>
                    </table>
                </div>
            </div>
        </section>
    </main>

    <br>
</template>