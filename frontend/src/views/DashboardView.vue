<script setup lang="ts">
import useEchart from "@/composables/useECharts";
import type { EChartsOption } from "echarts";
import { RotateCw, Sun, Settings, Zap, PhilippinePeso } from "lucide-vue-next";
import { onMounted, useTemplateRef } from "vue";

const barChartEl = useTemplateRef<HTMLDivElement>("bar-chart");
const barChart = useEchart(barChartEl);

const data = [
    { name: 'Search', value: 110000 },
    { name: 'Direct', value: 82000 },
    { name: 'Referral', value: 58000 },
    { name: 'Unassigned', value: 41000 },
    { name: 'Social', value: 30000 },
    { name: 'Newsletter', value: 15000 }
];

const maxValue = Math.max(...data.map(d => d.value));
// const roundedMax = Math.ceil(maxValue / 20000) * 20000;

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
            smooth: true,
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
            smooth: true,
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
  backgroundColor: 'transparent',

  title: {
    text: 'Top coupons',
    subtext: 'Last 7 days',
    left: 20,
    top: 15,
    textStyle: {
      fontSize: 16,
      fontWeight: 600,
      color: '#111827'
    },
    subtextStyle: {
      fontSize: 12,
      color: '#6B7280'
    }
  },

  tooltip: {
    trigger: 'item',
    backgroundColor: '#111827',
    textStyle: { color: '#fff' },
    formatter: '{b}<br/>{c}%'
  },

  legend: {
    orient: 'vertical',
    bottom: 20,
    left: 20,
    itemWidth: 14,
    itemHeight: 10,
    icon: 'roundRect',
    textStyle: {
      fontSize: 13,
      color: '#374151'
    },
    // formatter: function (name) {

    //   const item = (<any>lineChartOpt.series)[0].data.find((d: any) => d.name === name);
    //   return `${name} ${item.value}%`;
    // }
  },

  series: [
    {
      name: 'Coupons',
      type: 'pie',
      radius: ['70%', '82%'],
      center: ['50%', '48%'],
      avoidLabelOverlap: false,
      label: { show: false },
      labelLine: { show: false },

      emphasis: {
        scale: false
      },

      data: [
        {
          value: 72,
          name: 'Percentage discount',
          itemStyle: { color: '#3B82F6' }
        },
        {
          value: 18,
          name: 'Fixed card discount',
          itemStyle: { color: '#BFDBFE' }
        },
        {
          value: 10,
          name: 'Fixed product discount',
          itemStyle: { color: '#1D4ED8' }
        }
      ]
    }
  ],

  graphic: {
    type: 'text',
    left: 'center',
    top: '45%',
    style: {
      text: '72%',
    //   textAlign: 'center',
      fill: '#111827',
      fontSize: 28,
      fontWeight: 600
    }
  }
};
const pieChart = useEchart(pieChartEl);

onMounted(() => {
    barChart.setOptions(lineChartOpt);
    pieChart.setOptions(pieChartOpt);
})

</script>

<template>
    <main class="bg-base-200">
        <nav class="w-full h-16 flex items-center shadow">
            <div class="container mx-auto flex justify-between">
                <p class="font-semibold lg:text:lg">Dashboard</p>


                <div class="flex gap-3">
                    <label class="swap swap-rotate">
                        <!-- this hidden checkbox controls the state -->
                        <input type="checkbox" class="theme-controller" value="synthwave" />

                        <!-- sun icon -->
                        <Sun />

                        <!-- moon icon -->
                    </label>

                    <button class="btn btn-square btn-ghost btn-sm">
                        <Settings />
                    </button>

                    <div class="avatar avatar-placeholder">
                        <div class="bg-neutral text-neutral-content w-8 rounded-full">
                            <span class="text-xs">UI</span>
                        </div>
                    </div>
                </div>
            </div>


        </nav>


        <div class="container mx-auto mt-3 flex flex-col gap-3">

            <div class="my-3 flex justify-between items-end">
                <div class="flex items-center gap-1">
                    <input type="datetime-local" class="input" />

                    <button class="btn btn-ghost btn-square">
                        <RotateCw :size="20" />
                    </button>
                </div>
                <div></div>
            </div>

            <section class="grid grid-cols-1 gap-3 min-h-96">
                <div class="flex items-center justify-center rounded bg-base-100 shadow">
                    <div ref="bar-chart" class="h-full w-full">

                    </div>
                </div>
                <!-- <div class="grid grid-cols-2 gap-3">
                    <div class="border flex items-center justify-center rounded">
                        <p>KPI</p>
                    </div>
                    <div class="border flex items-center justify-center rounded">
                        <p>KPI</p>
                    </div>
                    <div class="border flex items-center justify-center rounded">
                        <p>KPI</p>
                    </div>
                    <div class="border flex items-center justify-center rounded">
                        <p>KPI</p>
                    </div>
                </div> -->
            </section>

            <section class="grid grid-cols-2 gap-3 min-h-86">
                <div class="flex flex-col gap-3">
                    <div class="grid grid-cols-2 gap-3">
                        <div class="bg-base-100 shadow p-3 flex items-center gap-4">
                            <div class="p-2 rounded-lg bg-primary text-primary-content">
                                <Zap />
                            </div>
                            <div>
                                <p class="font-semibold">Power Consumption</p>
                                <p class="text-lg">54 kWh</p>
                            </div>
                        </div>
                        <div class="bg-base-100 shadow p-3 flex items-center gap-4">
                            <div class="p-2 rounded-lg bg-success text-success-content">
                                <PhilippinePeso/>
                            </div>
                            <div>
                                <p class="font-semibold">Estimated Cost</p>
                                <p class="text-lg">632.25</p>
                            </div>
                        </div>
                    </div>
                    <div class="bg-base-100 shadow flex-1 flex items-center justify-center rounded">
                        <div ref="pie-chart" class="h-full w-full">

                        </div>
                    </div>
                </div>
                <div class="bg-base-100 shadow flex items-center justify-center rounded">Table</div>
            </section>
        </div>
    </main>
</template>