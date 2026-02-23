import { onMounted, onBeforeUnmount, type Ref } from 'vue'
import * as echarts from 'echarts'

export function useEChart(elRef: Ref<HTMLDivElement | null>) {
    let chart: echarts.ECharts | null = null

    const init = () => {
        if (!elRef.value) return
        chart = echarts.init(elRef.value)
    }

    const setOptions = (options: echarts.EChartsOption) => {
        chart?.setOption(options)
    }

    const resize = () => chart?.resize()

    onMounted(() => {
        init()
        window.addEventListener('resize', resize)
    })

    onBeforeUnmount(() => {
        window.removeEventListener('resize', resize)
        chart?.dispose()
    })

    return {
        setOptions
    }
}