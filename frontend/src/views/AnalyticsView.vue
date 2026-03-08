<script setup lang="ts">
import Navbar from "@/components/Navbar.vue";
import { Leaf, Plug, Zap, Atom} from "lucide-vue-next"


</script>

<template>
    <Navbar />

    <main class="px-5 mt-3">
        <section class="grid grid-cols-3 min-h-96 gap-3">
            <div class="card bg-base-100 col-span-1">
                <div class="card-body">
                    <div class="p-3 rounded-lg bg-red-400 text-white w-max">
                        <leaf />
                    </div>

                    <p class="text-xl mt-3">Total Energy Consumption</p>


                    <div>
                        <p class="text-3xl">1,284 kWh</p>
                    </div>
                </div>
            </div>
            <div class="col-span-2 grid grid-cols-2 gap-3">
                <div class="card bg-base-100">
                    <div class="card-body">
                        <div class="p-3 bg-red-400 rounded-lg text-white w-max">
                            <Plug />
                        </div>

                        <p class="text-lg mt-3">Average Voltage</p>


                        <div>
                            <p class="text-lg">233 V</p>
                        </div>
                    </div>
                </div>
                <div class="card bg-base-100">
                    <div class="card-body">
                        <div class="p-3 bg-red-400 rounded-lg text-white w-max">
                            <Zap />
                        </div>

                        <p class="text-lg mt-3">Average Power Draw</p>


                        <div>
                            <p class="text-lg">50 kWh</p>
                        </div>
                    </div>
                </div>

                                <div class="card bg-base-100">
                    <div class="card-body">
                        <div class="p-3 bg-red-400 rounded-lg text-white w-max">
                            <Atom />
                        </div>

                        <p class="text-lg mt-3">Average Current</p>


                        <div>
                            <p class="text-lg">233 A</p>
                        </div>
                    </div>
                </div>

                                <div class="card bg-base-100">
                    <div class="card-body">
                        <div class="p-3 bg-red-400 rounded-lg text-white w-max">
                            <Plug />
                        </div>

                        <p class="text-lg mt-3">Frequency</p>


                        <div>
                            <p class="text-lg">233 Hz</p>
                        </div>
                    </div>
                </div>
                
            </div>
        </section>

        <section class="mt-3 min-h-96 grid">
            <div class="card bg-gray-600">
                <div class="card-body"></div>
            </div>
        </section>

        <section class="mt-3 min-h-125 grid">
            <div class="card bg-gray-600">
                <div class="card-body"></div>
            </div>
        </section>
    </main>

    <br>
</template>