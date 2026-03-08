<script setup lang="ts">
import useEchart from "@/composables/useECharts";
import type { EChartsOption } from "echarts";
import { RotateCw, Zap, PhilippinePeso } from "lucide-vue-next";
import { onMounted, useTemplateRef } from "vue";
import Navbar from "@/components/Navbar.vue";

const barChartEl = useTemplateRef<HTMLDivElement>("bar-chart");
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

const pieChartEl = useTemplateRef<HTMLDivElement>("pie-chart");
const pieChartOpt: EChartsOption = {
    
    tooltip: {
        trigger: 'axis',
        axisPointer: {
            type: 'shadow'
        }
    },

    grid: {
        left: '5%',
        right: '5%',
        bottom: '15%',
        containLabel: true
    },

    xAxis: {
        type: 'category',
        data: ['Search', 'Direct', 'Referral', 'Unassigned', 'Social', 'Newsletter'],
        axisTick: { show: false },
        axisLine: {
            lineStyle: { color: '#E5E7EB' }
        },
        axisLabel: {
            rotate: 40
        }
    },

    yAxis: {
        type: 'value',
        min: 0,
        max: 120000,
        interval: 20000,
        axisLine: { show: false },
        splitLine: {
            lineStyle: {
                color: '#E5E7EB'
            }
        },
        axisLabel: {
            formatter: value => (value / 1000) + 'k'
        }
    },

    series: [
        {
            name: 'Users',
            type: 'bar',
            data: [112000, 82000, 58000, 41000, 30000, 15000],
            barWidth: 40,
            itemStyle: {
                color: '#93B4E7',
                borderRadius: [8, 8, 0, 0]
            }
        }
    ]
};
const pieChart = useEchart(pieChartEl);

onMounted(() => {
    barChart.setOptions(lineChartOpt);
    pieChart.setOptions(pieChartOpt);
})


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

</script>

<template>
    <main class="bg-base-200">
        <Navbar />


        <div class=" px-5 mt-3 flex flex-col gap-3">

            <div class="my-3 flex justify-between items-end">
                <div class="flex items-center gap-1">
                    <input type="datetime-local" class="input" />

                    <button class="btn btn-ghost btn-square">
                        <RotateCw :size="20" />
                    </button>
                </div>
                <div></div>
            </div>

            <section class="grid grid-cols-2 gap-3 min-h-150">
                <div class="flex flex-col gap-3">
                    <div class="grid grid-cols-2 gap-3">
                        <div class="card bg-base-100">
                            <div class="card-body">
                                <div class="p-2 rounded-lg bg-primary text-primary-content w-max">
                                    <Zap />
                                </div>
                                <div>
                                    <p class="font-semibold">Power Consumption</p>
                                    <p class="text-lg">54 kWh</p>
                                </div>
                            </div>
                        </div>

                        <div class="card bg-base-100">
                            <div class="card-body">
                                <div class="p-2 rounded-lg bg-success text-success-content w-max">
                                    <PhilippinePeso />
                                </div>
                                <div>
                                    <p class="font-semibold">Estimated Cost</p>
                                    <p class="text-lg">632.25</p>
                                </div>
                            </div>
                        </div>

                    </div>
                    <div class="card bg-base-100 flex-1">
                        <div class="card-body p-0 flex items-center justify-center">
                            <div ref="pie-chart" class="h-full w-full flex-1">

                            </div>
                        </div>
                    </div>
                </div>
                <div class="card bg-base-100">
                    <div class="card-body flex items-center justify-center">
                        <div ref="bar-chart" class="h-full w-full">
                        </div>
                    </div>
                </div>
            </section>

            <section class="grid grid-cols-1 gap-3 min-h-86">



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
        </div>
    </main>

    <br>
    <br>
    <br>
</template>