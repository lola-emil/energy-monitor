<script setup lang="ts">
import KPICard from '@/components/KPICard.vue';
import Navbar from '@/components/Navbar.vue';
import BaseChart from '@/components/BaseChart.vue';


</script>

<template>
    <Navbar />

    <div class="container mx-auto">
        <br>
        <div class="grid grid-cols-2 gap-5">
            <KPICard />
            <KPICard />
        </div>

        <br>

        <div class="h-96">
            <BaseChart :categories="['Jan', 'Feb', 'Mar', 'Apr', 'May']" :series="[
                { name: 'Revenue', data: [120, 200, 150, 80, 70] },
                { name: 'Expenses', data: [90, 140, 110, 60, 50] }
            ]" />
        </div>
    </div>
</template>
