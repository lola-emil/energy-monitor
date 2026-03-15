<script setup lang="ts">
import { MonitorSmartphone, LayoutDashboard, Settings, LogOut } from "lucide-vue-next";
import { RouterLink } from "vue-router";
</script>

<template>

    <div class="drawer lg:drawer-open">
        <input id="my-drawer-4" type="checkbox" class="drawer-toggle" />
        <div class="drawer-content h-screen">

            <slot />

        </div>

        <div class="drawer-side is-drawer-close:overflow-visible">
            <label for="my-drawer-4" aria-label="close sidebar" class="drawer-overlay"></label>
            <div class="flex min-h-full flex-col bg-base-100 is-drawer-close:w-14 is-drawer-open:w-64">
                <!-- Sidebar content here -->
                <ul class="menu w-full grow">
                    <!-- List item -->
                    <li>
                        <RouterLink to="/" class="is-drawer-close:tooltip is-drawer-close:tooltip-right"
                            data-tip="Dashboard & Analytics">
                            <LayoutDashboard :size="19" />
                            <span class="is-drawer-close:hidden">Dashboard & Analytics</span>
                        </RouterLink>
                    </li>

                    <li>
                        <RouterLink to="/devices" class="is-drawer-close:tooltip is-drawer-close:tooltip-right"
                            data-tip="Your Devices">
                            <MonitorSmartphone :size="19" />
                            <span class="is-drawer-close:hidden">Your Devices</span>
                        </RouterLink>
                    </li>

                </ul>


                <ul class="menu w-full">

                    <li>
                        <RouterLink to="/settings" class="is-drawer-close:tooltip is-drawer-close:tooltip-right"
                            data-tip="Settings">
                            <Settings :size="19" />
                            <span class="is-drawer-close:hidden">Settings</span>
                        </RouterLink>
                    </li>

                    <li>
                        <button class="is-drawer-close:tooltip is-drawer-close:tooltip-right" data-tip="Homepage">
                            <LogOut :size="19" />
                            <span class="is-drawer-close:hidden">Sign Out</span>
                        </button>
                    </li>

                </ul>
            </div>
        </div>
    </div>
</template>