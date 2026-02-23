<template>
    <div ref="chartRef" class="chart"></div>
</template>

<script setup lang="ts">
import { ref, watch, computed, onMounted } from 'vue'
import { useEChart } from '../composables/useEchart'
import type { EChartsOption } from 'echarts';

interface Props {
    categories: string[]
    series: { name: string; data: number[] }[]
}


const cssVar = ((
    name: string,
    fallback: string = '',
    element: HTMLElement = document.documentElement
): string => getComputedStyle(element).getPropertyValue(name).trim() || fallback);


const props = defineProps<Props>()

const chartRef = ref<HTMLDivElement | null>(null)

const { setOptions } = useEChart(chartRef)


const buildOptions = (): EChartsOption => ({
    backgroundColor: cssVar('--color-base-100'),

    color: [
        cssVar('--color-primary'),
        cssVar('--color-secondary'),
        cssVar('--color-accent'),
        cssVar('--color-info'),
        cssVar('--color-success'),
        cssVar('--color-warning'),
        cssVar('--color-error')
    ],

    textStyle: {
        color: cssVar('--color-base-content'),
        fontFamily: 'Inter, sans-serif'
    },

    grid: {
        left: 20,
        right: 20,
        top: 50,
        bottom: 30,
        containLabel: true
    },

    tooltip: { trigger: 'axis' },

    legend: {
        top: 10
    },

    xAxis: {
        type: 'category',
        data: props.categories,
    },

    yAxis: {
        type: 'value'
    },

    series: props.series.map(s => ({
        ...s,
        type: 'line',
        smooth: true,
        showSymbol: false,
        lineStyle: { width: 3 },
        areaStyle: { opacity: 0.15 },

        emphasis: {
            focus: 'none', // prevents dimming others
            lineStyle: {
                width: 3
            },
            areaStyle: {
                opacity: 0.15
            }
        }
    })),

    animationDuration: 600
})


const options = computed<EChartsOption>(() => buildOptions())

onMounted(() => {
    console.log(options.value);
    setOptions(options.value)
})

watch(
    () => [props.categories, props.series],
    () => {
        setOptions(buildOptions())
    },
    { deep: true, immediate: true }
)
</script>

<style scoped>
.chart {
    width: 100%;
    height: 400px;
}
</style>