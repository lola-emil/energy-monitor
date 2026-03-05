import { init as echartInit } from "echarts";
import { onBeforeUnmount, onMounted, useTemplateRef } from "vue";

const useEchart = (elRef: ReturnType<typeof useTemplateRef<HTMLDivElement>>) => {
    let chart: echarts.ECharts | null = null;

    const init = () => {
        if (!elRef.value) return;
        chart = echartInit(elRef.value);
    };

    const setOptions = (options: echarts.EChartsOption) => {
        chart?.setOption(options);
    };

    const resize = () => chart?.resize();

    onMounted(() => {
        init();
        window.addEventListener("resize", resize)
    })

    onBeforeUnmount(() => {
        window.removeEventListener("resize", resize);
        chart?.dispose();
    })

    return {
        setOptions
    }
}


export default useEchart;