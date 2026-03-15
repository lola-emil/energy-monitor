<script setup lang="ts">
import useEchart from '@/composables/useECharts';
import type { EChartsOption } from 'echarts';
import { onMounted, useTemplateRef } from 'vue';

import { RotateCw, Zap, PhilippinePeso } from "lucide-vue-next";

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
})
</script>

<template>

    <!-- Stats Row -->

    <div class="grid grid-cols-3 gap-12">
        <div class="card bg-base-100">
            <div class="card-body">
                <div class="p-2 rounded-lg bg-primary text-primary-content w-max">
                    <PhilippinePeso />
                </div>
                <div>
                    <p class="font-semibold">Estimated Bill</p>
                    <p class="text-lg">864.90</p>
                </div>
            </div>
        </div>
        <div class="card bg-base-100">
            <div class="card-body">
                <div class="p-2 rounded-lg bg-primary text-primary-content w-max">
                    <Zap />
                </div>
                <div>
                    <p class="font-semibold">Power Draw</p>
                    <p class="text-lg">54 kWh</p>
                </div>
            </div>
        </div>

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
    </div>
    <br>
    <!-- Chart Placeholder -->
    <div class="card bg-base-100 shadow p-6 h-125 mb-8 flex items-center justify-center">
        <div ref="line-chart" class="h-full w-full flex-1">
        </div>
    </div>


    <!-- Most Used Section -->
    <!-- <div class="card bg-base-100 shadow p-6">

        <div class="flex justify-between mb-6">
            <h3 class="font-semibold">Most Used</h3>

            <select class="select select-bordered select-sm">
                <option>All Time</option>
            </select>
        </div>

        <div class="space-y-4">

            <div class="flex justify-between items-center">
                <span>Figma</span>
                <span>24%</span>
                <span>20:05:42</span>
            </div>

            <div class="flex justify-between items-center">
                <span>Photoshop</span>
                <span>12.5%</span>
                <span>12:48:56</span>
            </div>

            <div class="flex justify-between items-center">
                <span>Chrome</span>
                <span>8.2%</span>
                <span>48:00:64</span>
            </div>

        </div>

    </div> -->
</template>