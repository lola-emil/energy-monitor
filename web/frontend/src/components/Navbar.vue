<template>
    <div class="navbar bg-base-100 border-b border-base-300 px-4 md:px-8 sticky top-0 z-50 shadow-sm">

        <div class="navbar-start">
            <div class="dropdown lg:hidden">
                <div tabindex="0" role="button" class="btn btn-ghost btn-circle">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24"
                        stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M4 6h16M4 12h8m-8 6h16" />
                    </svg>
                </div>
                <ul tabindex="0"
                    class="menu menu-sm dropdown-content bg-base-100 rounded-box z-[1] mt-3 w-52 p-2 shadow-lg border border-base-300">
                    <li>
                        <RouterLink to="/">Analytics</RouterLink>
                    </li>
                    <li>
                        <RouterLink to="/devices">Devices</RouterLink>
                    </li>
                    <li>
                        <RouterLink to="/settings">Settings</RouterLink>
                    </li>
                </ul>
            </div>

            <a class=" btn btn-ghost text-xl gap-2 normal-case hover:bg-transparent">
            </a>
        </div>

        <div class="navbar-center hidden lg:flex">
            <ul class="menu menu-horizontal px-1 gap-1 font-medium">
                <li>
                    <RouterLink to="/">
                        Analytics
                    </RouterLink>
                </li>
                <li>
                    <RouterLink to="/devices">
                        Devices
                    </RouterLink>
                </li>
                <li>
                    <RouterLink to="/settings">
                        Settings
                    </RouterLink>
                </li>
            </ul>
        </div>

        <div class="navbar-end gap-2">
            <label class="swap swap-rotate btn btn-ghost btn-circle btn-sm">
                <input type="checkbox" class="theme-controller" value="dark" />
                <svg class="swap-off fill-current w-5 h-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                    <path
                        d="M5.64,15.76l-2.61,2.61a1,1,0,0,0,1.41,1.41l2.61-2.61A10,10,0,0,0,19.86,4.14,1,1,0,0,0,19,3H14A10,10,0,0,0,5.64,15.76Z" />
                </svg>
                <svg class="swap-on fill-current w-5 h-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                    <path
                        d="M21.64,13a1,1,0,0,0-1.05-.14,8.05,8.05,0,0,1-3.37.73A8.15,8.15,0,0,1,9.08,5.49a8.59,8.59,0,0,1,.25-2A1,1,0,0,0,8,2.36,10.14,10.14,0,1,0,22,14.05,1,1,0,0,0,21.64,13Z" />
                </svg>
            </label>

            <button class="btn btn-ghost btn-circle btn-sm indicator">
                <span class="indicator-item badge badge-secondary badge-xs"></span>
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24"
                    stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
                </svg>
            </button>

            <div class="dropdown dropdown-end">
                <div tabindex="0" role="button"
                    class="btn btn-ghost btn-circle avatar ring ring-primary ring-offset-base-100 ring-offset-2">
                    <div class="w-10 rounded-full">
                        <img src="https://ui-avatars.com/api/?name=Alex+M&background=random" alt="User Avatar" />
                    </div>
                </div>
                <ul tabindex="0"
                    class="menu menu-sm dropdown-content bg-base-100 rounded-box z-[1] mt-3 w-52 p-2 shadow-lg border border-base-300">
                    <li class="menu-title px-4 py-2">
                        <span class="font-bold">Alex Morgan</span>
                        <span class="text-xs opacity-70">alex@example.com</span>
                    </li>
                    <li><a>Profile</a></li>
                    <li><a>Billing</a></li>
                    <div class="divider my-1"></div>
                    <li><a @click="auth.logout" class="text-error">Logout</a></li>
                </ul>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';


const auth = useAuthStore();
</script>