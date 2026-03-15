<script setup lang="ts">
import Navbar from '@/components/Navbar.vue';
import Sidebar from '@/components/Sidebar.vue';
</script>

<template>
    <Sidebar>
        <Navbar title="Device List"/>
        <div class="flex h-screen bg-base-200">

            <!-- Device List Panel -->
            <aside class="w-72 bg-base-100 border-r border-base-300 p-4 flex flex-col">

                <div class="space-y-3">

                    <div class="card bg-base-200 p-3 flex justify-between items-center">
                        <span>Macbook Air 13</span>
                    </div>

                    <div class="card bg-base-200 p-3 flex justify-between items-center">
                        <span>iPhone 7</span>
                    </div>

                    <div class="card bg-base-200 p-3 flex justify-between items-center">
                        <span>iPad</span>
                    </div>

                    <div class="card bg-base-200 p-3 flex justify-between items-center">
                        <span>AirPods</span>
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