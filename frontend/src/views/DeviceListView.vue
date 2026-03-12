<script setup lang="ts">
import Sidebar from '@/components/Sidebar.vue';
</script>

<template>
    <Sidebar>
        <div class="flex h-screen bg-base-200">

            <!-- Device List Panel -->
            <aside class="w-72 bg-base-100 border-r border-base-300 p-4 flex flex-col">

                <h2 class="text-sm opacity-70 mb-4">All Devices</h2>

                <div class="space-y-3">

                    <div class="card bg-base-200 p-3 flex justify-between items-center">
                        <span>Macbook Air 13</span>
                        <img class="w-8" src="https://via.placeholder.com/40" />
                    </div>

                    <div class="card bg-base-200 p-3 flex justify-between items-center">
                        <span>iPhone 7</span>
                        <img class="w-6" src="https://via.placeholder.com/30" />
                    </div>

                    <div class="card bg-base-200 p-3 flex justify-between items-center">
                        <span>iPad</span>
                        <img class="w-6" src="https://via.placeholder.com/30" />
                    </div>

                    <div class="card bg-base-200 p-3 flex justify-between items-center">
                        <span>AirPods</span>
                        <img class="w-5" src="https://via.placeholder.com/25" />
                    </div>

                    <button class="btn btn-outline w-full mt-3">
                        + Add Device
                    </button>

                </div>
            </aside>


            <!-- Main Dashboard -->
            <main class="flex-1 p-8 overflow-auto">

                <RouterView/>

            </main>

        </div>
    </Sidebar>
</template>