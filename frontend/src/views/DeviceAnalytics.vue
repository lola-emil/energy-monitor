<script setup lang="ts"></script>

<template>

    <!-- Stats Row -->
    <div class="flex justify-between items-center mb-8">

        <div class="flex gap-12">

            <div>
                <h1 class="text-4xl font-bold">40.2%</h1>
                <p class="text-sm opacity-70">Current Load</p>
            </div>

            <div>
                <h1 class="text-2xl font-semibold">24.8%</h1>
                <p class="text-sm opacity-70">Average Load</p>
            </div>

            <div>
                <h1 class="text-2xl font-semibold">36.2%</h1>
                <p class="text-sm opacity-70">Yesterday Load</p>
            </div>

        </div>

        <button class="btn btn-primary">
            Settings
        </button>

    </div>


    <!-- Chart Placeholder -->
    <div class="card bg-base-100 shadow p-6 h-72 mb-8 flex items-center justify-center">
        <span class="opacity-60">Chart Area</span>
    </div>


    <!-- Most Used Section -->
    <div class="card bg-base-100 shadow p-6">

        <div class="flex justify-between mb-6">
            <h3 class="font-semibold">Most Used</h3>

            <select class="select select-bordered select-sm">
                <option>All Time</option>
            </select>
        </div>

        <div class="space-y-4">

            <div class="flex justify-between items-center">
                <span>Figma</span>
                <span>24%</span>
                <span>20:05:42</span>
            </div>

            <div class="flex justify-between items-center">
                <span>Photoshop</span>
                <span>12.5%</span>
                <span>12:48:56</span>
            </div>

            <div class="flex justify-between items-center">
                <span>Chrome</span>
                <span>8.2%</span>
                <span>48:00:64</span>
            </div>

        </div>

    </div>
</template>