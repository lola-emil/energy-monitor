<template>
    <div ref="areachart" class="h-full"></div>

</template>

<script setup lang="ts">
import { onMounted, watch, useTemplateRef } from 'vue';
import { useDark } from '@vueuse/core';
import * as echarts from 'echarts';
import type { ECharts } from 'echarts';

const chartRef = useTemplateRef("areachart");
let chart: ECharts | null = null;

const isDark = useDark();

const option = {
    backgroundColor: 'transparent',
    animation: false,
    tooltip: {
        trigger: 'axis',
        backgroundColor: "oklch(0.21 0.006 285.885)",
        borderColor: "oklch(1 0 0 / 10%)",
        textStyle: { color: "oklch(0.985 0 0)" },
    },

    // legend: {
    //     bottom: 0,
    //     textStyle: { color: "oklch(0.274 0.006 286.033)" }
    // },

    grid: {
        left: '0',
        right: '0',
        top: 20,
        bottom: 60
    },

    xAxis: {
        type: 'category',
        boundaryGap: false,
        data: [
            'Apr 1', 'Apr 2', 'Apr 3', 'Apr 4', 'Apr 5', 'Apr 6',
            'Apr 7', 'Apr 8', 'Apr 9', 'Apr 10', 'Apr 11', 'Apr 12',
            'Apr 13', 'Apr 14', 'Apr 15', 'Apr 16', 'Apr 17', 'Apr 18',
            'Apr 19', 'Apr 20', 'Apr 21', 'Apr 22', 'Apr 23', 'Apr 24',
            'Apr 25', 'Apr 26', 'Apr 27', 'Apr 28', 'Apr 29', 'Apr 30',
            'May 1', 'May 2', 'May 3', 'May 4', 'May 5', 'May 6',
            'May 7', 'May 8', 'May 9', 'May 10', 'May 11', 'May 12',
            'May 13', 'May 14', 'May 15', 'May 16', 'May 17', 'May 18',
            'May 19', 'May 20', 'May 21', 'May 22', 'May 23', 'May 24',
            'May 25', 'May 26', 'May 27', 'May 28', 'May 29', 'May 30',
            'May 31', 'Jun 1', 'Jun 2', 'Jun 3', 'Jun 4', 'Jun 5',
            'Jun 6', 'Jun 7', 'Jun 8', 'Jun 9', 'Jun 10', 'Jun 11',
            'Jun 12', 'Jun 13', 'Jun 14', 'Jun 15', 'Jun 16', 'Jun 17',
            'Jun 18', 'Jun 19', 'Jun 20', 'Jun 21', 'Jun 22', 'Jun 23',
            'Jun 24', 'Jun 25', 'Jun 26', 'Jun 27', 'Jun 28', 'Jun 29'
        ],
        axisLine: { lineStyle: { color: "oklch(1 0 0 / 10%)" } },
        axisLabel: { color: "oklch(0.274 0.006 286.033)", show: false },

    },

    yAxis: {
        type: 'value',
        axisLabel: { color: "oklch(0.274 0.006 286.033)" },
        splitLine: {
            lineStyle: {
                color: "oklch(1 0 0 / 10%)",
                opacity: 0.3
            }
        }
    },

    series: [
        {
            name: 'Desktop',
            type: 'line',
            smooth: true,
            stack: 'total',
            data: [
                169, 184, 171, 153, 144, 159, 192, 191, 135, 138, 127,
                137, 120, 193, 180, 169, 158, 167, 147, 124, 193, 121,
                160, 187, 132, 175, 192, 138, 149, 163, 141, 128, 143,
                194, 145, 130, 126, 128, 145, 135, 175, 146, 172, 168,
                190, 191, 134, 170, 146, 151, 146, 164, 143, 188, 185,
                167, 192, 199, 177, 173, 139, 164, 173, 124, 167, 138,
                143, 173, 152, 150, 188, 166, 144, 130, 173, 126, 137,
                180, 157, 196, 163, 153, 122, 134, 198, 188, 188, 193,
                141, 133
            ],
            lineStyle: {
                color: "oklch(0.808 0.114 19.571)",
                width: 2
            },
            areaStyle: {
                color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                    { offset: 0, color: "oklch(0.808 0.114 19.571)" },
                    { offset: 1, color: 'transparent' }
                ])
            },
            showSymbol: false
        },
        {
            name: 'Mobile',
            type: 'line',
            smooth: true,
            stack: 'total',
            data: [
                145, 106, 93, 144, 148, 127, 99, 106, 84, 86, 130, 95,
                85, 136, 111, 100, 132, 126, 89, 100, 111, 80, 98, 126,
                89, 81, 145, 121, 87, 86, 85, 104, 103, 123, 114, 148,
                135, 136, 90, 128, 105, 130, 89, 98, 110, 131, 130, 81,
                136, 96, 103, 119, 81, 146, 96, 122, 101, 137, 95, 82,
                89, 144, 93, 134, 99, 88, 127, 120, 103, 142, 118, 130,
                101, 149, 114, 127, 149, 130, 105, 135, 94, 137, 142, 135,
                102, 135, 87, 131, 82, 147
            ],
            lineStyle: {
                color: "oklch(0.637 0.237 25.331)",
                width: 2
            },
            areaStyle: {
                color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                    { offset: 0, color: "#dc2626" }, // convert your oklch
                    { offset: 1, color: "rgba(249,115,22,0)" }
                ])
            },
            showSymbol: false
        }
    ]
};


const initChart = () => {
    if (!chartRef.value) return;

    console.log(option.xAxis.data.length)
    console.log(option.series[0]?.data.length)

    chart?.dispose();
    chart = echarts.init(chartRef.value);

    chart.setOption(option);

};

const resize = () => chart?.resize();

onMounted(() => {
    initChart();
    window.addEventListener('resize', resize);
});

watch(isDark, () => {
    initChart(); // re-read CSS vars
});
</script>