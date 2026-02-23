<script setup lang="ts">
import { UsersRound, TrendingUp } from "lucide-vue-next"


</script>

<template>
    <div class="card dark:border-gray-800 dark:bg-white/3">
        <div class="card-body p-5 md:p-6">
            <div class="flex items-center justify-center w-12 h-12 bg-gray-100 rounded-xl dark:bg-gray-800">
                <UsersRound />
            </div>
            <div class="flex items-end justify-between mt-5">
                <div><span class="text-sm text-gray-500 dark:text-gray-400">Customers</span>
                    <h4 class="mt-2 font-bold text-gray-800 text-3xl dark:text-white/90">3,782</h4>
                </div>
                <span
                    class="badge badge-soft badge-success">
                    <TrendingUp />
                    11.01%
                </span>
            </div>
        </div>
    </div>
</template>