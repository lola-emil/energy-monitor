<script setup lang="ts">
import { ref } from 'vue';
import { User } from "lucide-vue-next"

const upat = ref([1, 2, 1, 2]);
</script>

<template>
    <main class="container mx-auto px-5">
        <nav></nav>
        <br>
        <div class="flex flex-col">
            <!-- KPIs -->
            <div class="grid grid-cols-2 lg:grid-cols-4 gap-3">
                <div v-for="_ in upat"
                    class="flex flex-col border border-base-content/25 p-5 gap-3 rounded bg-base-100">
                    <p class="text-lg">Total Visitors</p>
                    <div class="p-2 rounded w-max bg-info text-info-content">
                        <User />
                    </div>
                    <p class="text-2xl">5.52M</p>

                    <p class="text-sm text-neutral-400">See in-depth source</p>
                </div>
            </div>

            <!-- Line chart -->
            <div></div>
        </div>

        <div></div>
    </main>
</template>