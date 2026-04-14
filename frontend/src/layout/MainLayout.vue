<template>

    <SidebarProvider>
        <Sidebar>
            <SidebarHeader>
                <SidebarMenu>
                    <SidebarMenuItem>
                        <SidebarMenuButton size="lg">
                            <div
                                class="flex aspect-square size-8 items-center justify-center rounded-lg bg-sidebar-primary text-sidebar-primary-foreground">
                                <GalleryVerticalEnd class="size-4" />
                            </div>
                            <div class="grid flex-1 text-left text-sm leading-tight">
                                <span class="truncate font-semibold">Energy Consumption Monitor</span>
                                <span class="truncate text-xs"></span>
                            </div>
                        </SidebarMenuButton>
                    </SidebarMenuItem>
                </SidebarMenu>
            </SidebarHeader>
            <SidebarContent>
                <SidebarGroup>
                    <SidebarGroupContent>
                        <SidebarMenu>
                            <SidebarMenuItem>
                                <SidebarMenuButton as-child>
                                    <RouterLink to="/">
                                        <Home />
                                        <span>Dashboard</span>
                                    </RouterLink>
                                </SidebarMenuButton>
                            </SidebarMenuItem>
                        </SidebarMenu>
                    </SidebarGroupContent>
                </SidebarGroup>

                <SidebarGroup>
                    <SidebarGroupLabel>Your Devices</SidebarGroupLabel>
                    <SidebarGroupContent>
                        <SidebarMenu>
                            <SidebarMenuItem v-for="device in devices">
                                <SidebarMenuButton as-child>
                                    <RouterLink :to="'/devices/' + device.device_id">
                                        <MonitorSmartphone />
                                        <span>{{ device.device_name }}</span>
                                    </RouterLink>

                                </SidebarMenuButton>
                            </SidebarMenuItem>


                            <SidebarMenuItem @click="openDialog">
                                <SidebarMenuButton>
                                    <Plus />
                                    <span>Add New Device</span>
                                </SidebarMenuButton>
                            </SidebarMenuItem>
                        </SidebarMenu>
                    </SidebarGroupContent>
                </SidebarGroup>
            </SidebarContent>
            <SidebarFooter>
                <NavUser :user="data.user" />
            </SidebarFooter>
            <SidebarRail />
        </Sidebar>
        <SidebarInset>
            <header
                class="flex h-16 shrink-0 items-center gap-2 transition-[width,height] ease-linear group-has-data-[collapsible=icon]/sidebar-wrapper:h-12 border-b">
                <div class="flex items-center gap-2 px-4">
                    <SidebarTrigger class="-ml-1" />
                </div>
            </header>
            <RouterView />

            <AddDeviceModal></AddDeviceModal>

        </SidebarInset>
    </SidebarProvider>
</template>


<script setup lang="ts">
import {
    Sidebar,
    SidebarContent,
    SidebarFooter,
    SidebarGroup,
    SidebarGroupContent,
    SidebarGroupLabel,
    SidebarHeader,
    SidebarInset,
    SidebarMenu,
    SidebarMenuButton,
    SidebarMenuItem,
    SidebarProvider,
    SidebarRail,
    SidebarTrigger,
} from '@/components/ui/sidebar'
import { RouterView } from 'vue-router';
import { Home, Plus, MonitorSmartphone } from "lucide-vue-next"
import NavUser from '@/components/NavUser.vue';
import { useAPI } from '@/composables/useAPI';
import { axiosInstance } from '@/api/axios';
import { useAuthStore } from '@/stores/auth';
import { onMounted, ref } from 'vue';
import AddDeviceModal from '@/components/AddDeviceModal.vue';
import { useModalStore } from '@/stores/modal';
import { GalleryVerticalEnd } from "lucide-vue-next"

interface Device {
    Devicecode: string
    created_at: string
    device_id: number
    device_name: string
    id: number, is_active: boolean
    last_active: string | null
    user_id: number
}
const devices = ref<Device[]>([]);

const auth = useAuthStore();
const modal = useModalStore();

const data = {
    user: {
        name: "shadcn",
        email: "m@example.com",
        avatar: "/avatars/shadcn.jpg",
    },
}


const fetchDevices = () => {
    axiosInstance.get<Device[]>("/api/devices", {
        headers: {
            "Authorization": `Bearer ${auth.token}`,
        }
    }).then(res => {
        devices.value = res.data;
    }).catch(err => {
        console.log(err);
    })
}

const openDialog = () => {
    modal.open();
}

onMounted(() => {
    fetchDevices();
})

</script>
