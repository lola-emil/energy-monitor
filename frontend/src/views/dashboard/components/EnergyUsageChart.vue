<template>
    <div ref="barchart" class="h-full"></div>
</template>

<script setup lang="ts">
import { onMounted, onBeforeUnmount, watch, useTemplateRef } from 'vue';
import { useDark } from '@vueuse/core';
import * as echarts from 'echarts';
import type { ECharts, EChartsOption } from 'echarts';
import { useThemeColors } from '@/composables/useThemeColors';

const chartRef = useTemplateRef("barchart");
let chart: ECharts | null = null;

const isDark = useDark();
const themeColors = useThemeColors();

const months = [
    'Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun',
    'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'
];

const props = defineProps<{
    months?: string[];
    data: number[]
}>();


watch(() => props.data, (newVal) => {
    initChart(); 
});

const getOption = (): EChartsOption => {
    const c = themeColors.colors.value!;

    return {
        backgroundColor: 'transparent',

        tooltip: {
            trigger: 'axis',
            backgroundColor: c.popover,
            textStyle: {
                color: c.popoverForeground ?? '#fff'
            }
        },

        grid: {
            left: '5%',
            right: '5%',
            bottom: '8%',
            top: 70
        },

        xAxis: {
            type: 'category',
            data: months,
            axisLine: {
                lineStyle: { color: c.border }
            },
            axisLabel: {
                color: c.muted
            }
        },

        yAxis: {
            type: 'value',
            axisLabel: {
                color: c.muted
            },
            splitLine: {
                lineStyle: {
                    color: c.border,
                    type: 'dashed'
                }
            }
        },

        series: [
            {
                type: 'bar',
                data: props.data,
                barWidth: 28,

                itemStyle: {
                    borderRadius: [6, 6, 0, 0],

                    color: c.primary
                },
                emphasis: {
                    itemStyle: {
                        color: c.primary,
                        opacity: 0.85
                    }
                }
            }
        ]
    };
};

const initChart = () => {
    if (!chartRef.value) return;

    chart?.dispose();
    chart = echarts.init(chartRef.value);

    chart.setOption(getOption());
};

onMounted(() => {
    console.log(props.data);
    initChart();
    window.addEventListener('resize', resize);
});

onBeforeUnmount(() => {
    chart?.dispose();
    window.removeEventListener('resize', resize);
});

const resize = () => chart?.resize();


watch(isDark, () => {
    initChart(); // re-read CSS vars
});
</script>