<template>
    <div class="min-h-screen bg-base-200 flex flex-col">
        <Navbar />

        <main class="flex-1 p-4 md:p-8">
            <div class="space-y-6">
                <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
                    <div>
                        <h2 class="text-2xl font-bold text-base-content">Device Management</h2>
                        <p class="text-base-content/60 text-sm mt-1">Monitor and control your connected appliances</p>
                    </div>
                    <button class="btn btn-primary gap-2">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24"
                            stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                        </svg>
                        Add Device
                    </button>
                </div>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="form-control flex-1">
                        <input type="text" v-model="searchQuery"
                            placeholder="Search by appliance name or device code..."
                            class="input input-bordered w-full" />
                    </div>
                    <select v-model="statusFilter" class="select select-bordered w-full md:w-48">
                        <option value="all">All Status</option>
                        <option value="active">Active</option>
                        <option value="offline">Offline</option>
                    </select>
                </div>

                <div class="card bg-base-100 shadow-sm border border-base-300 overflow-hidden">
                    <div class="overflow-x-auto">
                        <table class="table table-zebra">
                            <thead>
                                <tr>
                                    <th>Device Code</th>
                                    <th>Appliance Name</th>
                                    <th>Location</th>
                                    <th>Last Updated</th>
                                    <th>Status</th>
                                    <th class="text-right">Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr v-for="device in filteredDevices" :key="device.id">
                                    <td>
                                        <span
                                            class="font-mono text-xs bg-base-200 text-base-content/80 px-2 py-1 rounded border border-base-300">
                                            {{ device.device_code }}
                                        </span>
                                    </td>

                                    <td>
                                        <div class="font-bold text-base-content">{{ device.name }}</div>
                                    </td>

                                    <td>
                                        <div class="flex items-center gap-2 text-base-content/70 text-sm">
                                            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none"
                                                viewBox="0 0 24 24" stroke="currentColor">
                                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                                    d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                                    d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
                                            </svg>
                                            {{ device.location }}
                                        </div>
                                    </td>

                                    <td class="text-base-content/60 text-sm whitespace-nowrap">
                                        {{ device.updated_at }}
                                    </td>

                                    <td>
                                        <span class="badge gap-1.5"
                                            :class="device.status === 'online' ? 'badge-success text-success-content' : 'badge-ghost text-base-content/60'">
                                            <span class="w-1.5 h-1.5 rounded-full"
                                                :class="device.status === 'online' ? 'bg-success-content' : 'bg-base-content/40'"></span>
                                            {{ device.status === 'online' ? 'Active' : 'Offline' }}
                                        </span>
                                    </td>

                                    <td class="text-right">
                                        <div class="dropdown dropdown-end">
                                            <div tabindex="0" role="button" class="btn btn-ghost btn-xs">
                                                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none"
                                                    viewBox="0 0 24 24" stroke="currentColor">
                                                    <path stroke-linecap="round" stroke-linejoin="round"
                                                        stroke-width="2"
                                                        d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z" />
                                                </svg>
                                            </div>
                                            <ul tabindex="0"
                                                class="dropdown-content menu bg-base-100 rounded-box z-[1] w-32 p-2 shadow border border-base-300">
                                                <li><a>View Details</a></li>
                                                <li><a>Edit</a></li>
                                                <div class="divider my-1"></div>
                                                <li><a class="text-error">Remove</a></li>
                                            </ul>
                                        </div>
                                    </td>
                                </tr>

                                <tr v-if="filteredDevices.length === 0">
                                    <td colspan="6" class="text-center py-12 text-base-content/50">
                                        <p class="font-medium">No devices found matching your criteria.</p>
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </div>

                    <div class="card-footer border-t border-base-300 p-4 flex justify-between items-center bg-base-100">
                        <span class="text-sm text-base-content/60">Showing {{ filteredDevices.length }} of
                            {{ devices.length }}
                            devices</span>
                        <div class="join">
                            <button class="join-item btn btn-sm">«</button>
                            <button class="join-item btn btn-sm btn-active">1</button>
                            <button class="join-item btn btn-sm">»</button>
                        </div>
                    </div>
                </div>

            </div>
        </main>
    </div>
</template>

<script setup lang="ts">
import Navbar from '@/components/Navbar.vue';
import { applianceService } from '@/services/appliance.service';
import type { Appliance } from '@/types/appliance';
import { ref, computed, onMounted } from 'vue';

const searchQuery = ref('');
const statusFilter = ref('all');

// Dummy Data
const devices = ref<Appliance[]>([]);

// Computed property for filtering
const filteredDevices = computed(() => {
    return devices.value.filter(device => {
        const matchesSearch =
            device.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
            device.device_code.toLowerCase().includes(searchQuery.value.toLowerCase());
        const matchesStatus = statusFilter.value === 'all' || device.status === statusFilter.value;
        return matchesSearch && matchesStatus;
    });
});

onMounted(async () => {
    const data = await applianceService.getAll();
    console.log(data);
    devices.value = data;
});
</script>