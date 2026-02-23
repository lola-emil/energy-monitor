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

const props = defineProps<Props>()

const chartRef = ref<HTMLDivElement | null>(null)

const { setOptions } = useEChart(chartRef)

const options = computed<EChartsOption>(() => ({
    xAxis: {
        type: 'category',
        data: props.categories
    },
    yAxis: {
        type: 'value'
    },
    series: props.series.map(s => ({
        name: s.name,
        type: 'bar',
        smooth: true,
        data: s.data
    }))
}))

onMounted(() => {
    setOptions(options.value)
})

watch(
    options,
    (newOptions) => {
        setOptions(newOptions)
    },
    { immediate: true }
)
</script>

<style scoped>
.chart {
    width: 100%;
    height: 400px;
}
</style>