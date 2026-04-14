<template>
    <div ref="barchart" class="h-full"></div>
</template>

<script setup lang="ts">
import { onMounted, onBeforeUnmount, watch, useTemplateRef } from 'vue';
import { useDark } from '@vueuse/core';
import * as echarts from 'echarts';
import type { ECharts, EChartsOption } from 'echarts';

const chartRef = useTemplateRef("barchart");
let chart: ECharts | null = null;

const isDark = useDark();

const getThemeColors = () => {
    const css = getComputedStyle(document.documentElement);

    return {
        background: css.getPropertyValue('--background').trim(),
        primary: css.getPropertyValue('--chart-2').trim(),
        foreground: css.getPropertyValue('--popover').trim(),
        popover: css.getPropertyValue('--popover').trim(),
        popoverForeground: css.getPropertyValue('--popover-foreground').trim(),
        muted: css.getPropertyValue('--muted-foreground').trim(),
        border: css.getPropertyValue('--border').trim(),
        chart1: css.getPropertyValue('--chart-1').trim(),
    };
};

const getOption = (): EChartsOption => {
    const c = getThemeColors();

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
            data: [
                'Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun',
                'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'
            ],
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
                data: [750, 720, 680, 600, 580, 550, 590, 610, 650, 720, 840, 980],
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