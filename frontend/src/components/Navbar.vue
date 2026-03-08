<script setup lang="ts">
import { Sun, PanelLeftOpen } from "lucide-vue-next";

</script>

<template>
    <nav class="navbar w-full bg-base-100">
        <div class=" w-full flex items-center">
            <div class="flex-1 flex items-center">
                <label for="my-drawer-4" aria-label="open sidebar" class="btn btn-square btn-ghost">
                    <!-- Sidebar toggle icon -->
                    <PanelLeftOpen :size="19" />
                </label>
                <div class="px-4">Dashboard</div>
            </div>


            <div class="flex gap-3 flex-none">
                <label class="swap swap-rotate">
                    <!-- this hidden checkbox controls the state -->
                    <input type="checkbox" class="theme-controller" value="synthwave" />

                    <!-- sun icon -->
                    <Sun />

                    <!-- moon icon -->
                </label>

                <div class="avatar avatar-placeholder">
                    <div class="bg-neutral text-neutral-content w-8 rounded-full">
                        <span class="text-xs">UI</span>
                    </div>
                </div>
            </div>
        </div>
    </nav>
</template>