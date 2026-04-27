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
                                        <LayoutDashboard />
                                        <span>Overview</span>
                                    </RouterLink>
                                </SidebarMenuButton>
                            </SidebarMenuItem>

                            <SidebarMenuItem>
                                <SidebarMenuButton as-child>
                                    <RouterLink to="/analytics">
                                        <ChartLine />
                                        <span>Analytics</span>
                                    </RouterLink>
                                </SidebarMenuButton>
                            </SidebarMenuItem>

                            <SidebarMenuItem>
                                <SidebarMenuButton as-child>
                                    <RouterLink to="/devices">
                                        <MonitorSmartphone />
                                        <span>Devices / Appliances</span>
                                    </RouterLink>
                                </SidebarMenuButton>
                            </SidebarMenuItem>

                            <SidebarMenuItem>
                                <SidebarMenuButton as-child>
                                    <RouterLink to="/settings">
                                        <Settings />
                                        <span>Settings</span>
                                    </RouterLink>
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
        </SidebarInset>
    </SidebarProvider>
</template>


<script setup lang="ts">
import {
    Sidebar,
    SidebarContent,
    SidebarFooter,
    SidebarGroup,
    SidebarGroupContent, SidebarHeader,
    SidebarInset,
    SidebarMenu,
    SidebarMenuButton,
    SidebarMenuItem,
    SidebarProvider,
    SidebarRail,
    SidebarTrigger
} from '@/components/ui/sidebar';
import { RouterView } from 'vue-router';
import { MonitorSmartphone, LayoutDashboard, ChartLine, Settings } from "lucide-vue-next";
import NavUser from '@/components/NavUser.vue';
import { useModalStore } from '@/stores/modal';
import { GalleryVerticalEnd } from "lucide-vue-next";

const data = {
    user: {
        name: "shadcn",
        email: "m@example.com",
        avatar: "/avatars/shadcn.jpg",
    },
}

</script>
