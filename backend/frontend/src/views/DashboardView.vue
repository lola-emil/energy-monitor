<template>
    <div class="min-h-screen flex">

        <!-- Sidebar -->
        <aside class="hidden md:flex md:flex-col w-64 bg-white border-r">
            <div class="h-16 flex items-center px-6 font-semibold text-xl border-b">
                Aurora
            </div>

            <nav class="flex-1 p-4 space-y-2 text-sm">
                <a href="#" class="block px-4 py-2 rounded-lg bg-blue-50 text-blue-600 font-medium">Dashboard</a>
                <a href="#" class="block px-4 py-2 rounded-lg hover:bg-gray-100">Users</a>
                <a href="#" class="block px-4 py-2 rounded-lg hover:bg-gray-100">Account</a>
                <a href="#" class="block px-4 py-2 rounded-lg hover:bg-gray-100">Login</a>
                <a href="#" class="block px-4 py-2 rounded-lg hover:bg-gray-100">Documentation</a>
            </nav>

            <div class="p-4 border-t">
                <div class="bg-gradient-to-br from-green-100 to-blue-100 p-4 rounded-xl text-sm">
                    <p class="font-semibold mb-2">Get More Features</p>
                    <button class="w-full bg-black text-white py-2 rounded-lg mt-3">
                        Explore
                    </button>
                </div>
            </div>
        </aside>


        <!-- Main -->
        <div class="flex-1 flex flex-col">

            <!-- Topbar -->
            <header class="h-16 bg-white border-b flex items-center justify-between px-6">
                <div class="flex items-center gap-4">
                    <button class="md:hidden">☰</button>
                    <input type="text" placeholder="Search..."
                        class="hidden md:block w-80 px-4 py-2 rounded-full bg-gray-100 focus:outline-none" />
                </div>

                <div class="flex items-center gap-4">
                    <div class="w-8 h-8 bg-gray-200 rounded-full"></div>
                    <div class="w-8 h-8 bg-gray-200 rounded-full"></div>
                    <div class="w-8 h-8 bg-gray-300 rounded-full"></div>
                </div>
            </header>


            <!-- Content -->
            <main class="flex-1 p-6 space-y-6">

                <!-- Stats Cards -->
                <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
                    <div class="bg-white p-6 rounded-xl border">Card</div>
                    <div class="bg-white p-6 rounded-xl border">Card</div>
                    <div class="bg-white p-6 rounded-xl border">Card</div>
                    <div class="bg-white p-6 rounded-xl border">Card</div>
                </div>


                <!-- Charts Row -->
                <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
                    <div class="lg:col-span-2 bg-white p-6 rounded-xl border h-80">
                        Line Chart Area
                    </div>

                    <div class="bg-white p-6 rounded-xl border h-80">
                        Side Chart
                    </div>
                </div>


                <!-- Table + Bar Chart -->
                <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
                    <div class="bg-white p-6 rounded-xl border h-96">
                        Bar Chart
                    </div>

                    <div class="bg-white p-6 rounded-xl border h-96">
                        Table Area
                    </div>
                </div>


                <!-- Promo Section -->
                <div class="bg-gradient-to-r from-green-200 via-blue-200 to-indigo-200 rounded-2xl p-10 text-center">
                    <h2 class="text-2xl font-semibold mb-2">
                        Discover More with Our Pro License
                    </h2>
                    <p class="mb-6">Starts from only $59</p>
                    <div class="flex justify-center gap-4">
                        <button class="bg-black text-white px-6 py-2 rounded-lg">
                            See Now
                        </button>
                        <button class="bg-white px-6 py-2 rounded-lg">
                            Download Free
                        </button>
                    </div>
                </div>

            </main>


            <!-- Footer -->
            <footer class="bg-white border-t text-center py-4 text-sm text-gray-500">
                © 2026 Aurora Dashboard
            </footer>

        </div>

    </div>

</template>


<script setup lang="ts">
import { User } from "lucide-vue-next"

</script>