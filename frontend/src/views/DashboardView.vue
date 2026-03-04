<script setup lang="ts">
import { RotateCw, Sun, Settings } from "lucide-vue-next"
</script>

<template>
    <main>
        <nav class="w-full h-16 border-b flex items-center">
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

            <section class="grid grid-cols-2 gap-3 min-h-96">
                <div class="border flex items-center justify-center rounded">
                    Line chart
                </div>
                <div class="grid grid-cols-2 gap-3">
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
                </div>
            </section>

            <section class="grid grid-cols-2 gap-3 min-h-86">
                <div class="flex flex-col gap-3">
                    <div class="grid grid-cols-2 gap-3 flex-1">
                        <div class="border flex items-center justify-center rounded">KPI</div>
                        <div class="border flex items-center justify-center rounded">KPI</div>
                    </div>
                    <div class="border flex-1 flex items-center justify-center rounded">Bar chart</div>
                </div>
                <div class="border flex items-center justify-center rounded">Table</div>
            </section>
        </div>
    </main>
</template>