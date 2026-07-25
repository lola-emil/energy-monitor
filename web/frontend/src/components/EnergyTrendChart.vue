<template>
    <div class="w-full h-96 relative py-5">
        <Bar :data="chartData" :options="chartOptions" />
    </div>
</template>

<script setup lang="ts">
import { Bar } from 'vue-chartjs';
import {
    Chart as ChartJS,
    Title,
    Tooltip,
    Legend,
    BarElement,
    CategoryScale,
    LinearScale,
    type ChartOptions
} from 'chart.js';

ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale);

const labels = Array.from({ length: 25 }, (_, i) => `Jul ${i + 1}`);
const currentPeriodData = [12, 19, 15, 25, 22, 30, 28, 15, 20, 35, 40, 32, 28, 18, 22, 25, 30, 38, 42, 35, 28, 20, 15, 18, 22];
const previousPeriodData = [10, 15, 12, 20, 18, 25, 22, 12, 18, 30, 35, 28, 24, 15, 18, 20, 25, 32, 38, 30, 24, 18, 12, 15, 18];

const chartData = {
    labels: labels,
    datasets: [
        {
            label: 'Current Period',
            backgroundColor: '#6366f1',
            borderRadius: 4,
            borderSkipped: false,
            data: currentPeriodData,
            barPercentage: 0.7,
            categoryPercentage: 0.8
        },
        {
            label: 'Previous Period',
            backgroundColor: '#f97316', // Orange (Matches inspiration image)
            borderRadius: 4,
            borderSkipped: false,
            data: previousPeriodData,
            barPercentage: 0.7,
            categoryPercentage: 0.8
        }
    ]
};

const chartOptions: ChartOptions<'bar'> = {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
        legend: {
            position: 'bottom',
            labels: {
                usePointStyle: true,
                pointStyle: 'rectRounded',
                padding: 20,
                font: { size: 12, family: 'system-ui, sans-serif' },
            }
        },
        tooltip: {
            backgroundColor: 'rgba(17, 24, 39, 0.9)',
            titleFont: { size: 13 },
            bodyFont: { size: 13 },
            padding: 10,
            cornerRadius: 8,
            displayColors: true,
            callbacks: {
                label: function (context) {
                    return ` ${context.dataset.label}: ${context.raw} kWh`;
                }
            }
        }
    },
    scales: {
        x: {
            grid: { display: false },
            ticks: {
                font: { size: 11 },
                maxTicksLimit: 10
            }
        },
        y: {
            beginAtZero: true,
            border: { display: false },
            grid: {
            },
            ticks: {
                font: { size: 11 },
                stepSize: 10,
                callback: function (value) {
                    return value + ' kWh';
                }
            }
        }
    }
};
</script>