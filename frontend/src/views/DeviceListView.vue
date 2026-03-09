<script setup lang="ts">
import Navbar from "@/components/Navbar.vue";
import { Wifi, Plus } from "lucide-vue-next";
import { ref } from "vue";

const items = ref([1, 2, 3, 4, 5, 6, 7]);
</script>

<template>
    <main>
        <Navbar title="Your Devices"/>

        <div class="px-5 mt-3">
            <section>
                <div class="w-full flex justify-between my-3 items-end">
                    <input type="text" class="input" placeholder="Search">

                    <button class="btn btn-primary">
                        <Plus :size="16" class="p-0"/>
                        Add Device
                    </button>
                </div>

                <div class="divider"></div>

                <div class="mt-3 grid grid-cols-2 lg:grid-cols-4 gap-3">
                    <div class="card bg-base-100" v-for="_ in items">
                        <div class="card-body">
                            <p class="text-lg">Refrigerator</p>
                            <p class="text-sm">Last seen: 1 min ago</p>
                            <div class="divider"></div>
                            <div>
                                <div class="badge badge-success">
                                    <Wifi :size="18" />
                                    Online
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </section>
        </div>
    </main>
</template>