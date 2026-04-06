<script setup lang="ts">
import useEchart from '@/composables/useECharts';
import type { EChartsOption } from 'echarts';
import { onMounted, useTemplateRef } from 'vue';
import {
    Card, CardContent, CardTitle
} from '@/components/ui/card';
import { useRoute } from 'vue-router';
import { MonitorSmartphone } from "lucide-vue-next";


const route = useRoute();

const lineChartEl = useTemplateRef<HTMLDivElement>("line-chart");
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

const lineChart = useEchart(lineChartEl);

onMounted(() => {
    lineChart.setOptions(lineChartOpt);
    console.log(route.params.id);
})

</script>

<template>
    <main class="px-5 mt-3">
        <Card>
            <CardContent>
                <div class="w-full flex">
                    <div class="flex gap-5">
                        <div class="p-3 bg-primary rounded-lg text-white w-max">
                            <MonitorSmartphone />
                        </div>
                        <div class="flex flex-col gap-1">
                            <h1 class="text-xl font-semibold tracking-tight">PC</h1>
                            <p class="text-sm font-medium leading-none">ID: RL901</p>
                        </div>
                    </div>

                    <div></div>
                </div>
            </CardContent>
        </Card>

        <section class="mt-5 grid grid-cols-2 gap-5 h-[80vh]">
            <Card>
                <CardContent>
                    <CardTitle class="">
                        Energy Usage (Monthly)
                    </CardTitle>
                </CardContent>
            </Card>

            <Card>
                <CardContent>
                    <CardTitle class="">
                        Current and Power
                    </CardTitle>
                </CardContent>
            </Card>

            <Card>
                <CardContent></CardContent>
            </Card>

            <Card>
                <CardContent>
                    <CardTitle class="">
                        Voltage
                    </CardTitle>
                </CardContent>
            </Card>


        </section>
    </main>
</template>