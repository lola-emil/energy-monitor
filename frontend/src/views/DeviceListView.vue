<script setup lang="ts">
import { Sun, Wifi, Settings, Plus } from "lucide-vue-next"
import { ref } from "vue";

const items = ref([1, 2, 3, 4, 5, 6, 7]);
</script>

<template>
    <main>
        <nav class="w-full h-20 border-b flex items-center">
            <div class="container mx-auto flex justify-between">
                <p class="font-semibold lg:text:lg">Device List</p>


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


        <div class="container mx-auto mt-3">
            <section>
                <div class="w-full flex justify-between">
                    <input type="text" class="input" placeholder="Search">

                    <button class="btn btn-sm btn-primary">
                        <Plus :size="16" />
                        Add Device
                    </button>
                </div>
                
                <div class="divider"></div>

                <div class="mt-3 grid grid-cols-3 gap-3">
                    <div class="border p-3 rounded" v-for="value in items">
                        <p class="text-lg">Refrigerator</p>
                        <p class="text-sm">Appliance</p>
                        <br>
                        <p class="text-sm">Last seen: 1 min ago</p>
                        <div class="divider"></div>
                        <div>
                            <div class="badge badge-success">
                                <Wifi :size="18"/>
                                Online
                            </div>
                        </div>
                    </div>         
                </div>
            </section>
        </div>
    </main>
</template>