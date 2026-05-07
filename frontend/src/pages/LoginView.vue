<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import {
    Field,
    FieldGroup,
    FieldLabel,
    FieldSet,
} from '@/components/ui/field'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { useAuthStore } from '@/stores/auth'
import { authService } from '@/services/auth.service'
import Logo from '@/components/Logo.vue'
import { LockKeyholeIcon, UserIcon, Loader2Icon } from 'lucide-vue-next'

const router = useRouter()
const authStore = useAuthStore()

const form = reactive({
    username: '',
    password: '',
})

const isLoading = ref(false)
const errorMessage = ref('')

const handleLogin = async () => {
    errorMessage.value = ''

    if (!form.username || !form.password) {
        errorMessage.value = 'Username and password are required.'
        return
    }

    isLoading.value = true

    try {
        const response = await authService.login(form.username, form.password)
        authStore.setToken(response.token)
        router.push('/')
    } catch (error: any) {
        errorMessage.value =
            error?.response?.data?.message ||
            'Invalid username or password.'
    } finally {
        isLoading.value = false
    }
}
</script>

<template>
    <main class="min-h-screen bg-muted/20">
        <div class="grid min-h-screen lg:grid-cols-2">
            <!-- Left branding panel -->
            <section class="hidden lg:flex lg:flex-col lg:justify-between lg:border-r lg:bg-background/80 lg:p-10">
                <div class="flex items-center gap-3">
                    <div class="w-10">
                        <Logo />
                    </div>
                    <div>
                        <p class="text-sm font-semibold tracking-[0.2em]">VOLTRIS</p>
                        <p class="text-xs text-muted-foreground">
                            Smart Energy Monitoring
                        </p>
                    </div>
                </div>

                <div class="max-w-md space-y-5">
                    <div
                        class="inline-flex rounded-full border bg-background px-3 py-1 text-xs text-muted-foreground shadow-sm">
                        Real-time monitoring platform
                    </div>

                    <div class="space-y-3">
                        <h1 class="text-4xl font-semibold tracking-tight">
                            Monitor your energy system with clarity.
                        </h1>
                        <p class="text-sm leading-6 text-muted-foreground">
                            Access your dashboard to track appliance activity, analyze usage trends,
                            review alerts, and manage system settings in one place.
                        </p>
                    </div>

                    <div class="grid gap-3 pt-2">
                        <div class="rounded-2xl border bg-muted/20 p-4">
                            <p class="text-sm font-medium">Live monitoring</p>
                            <p class="mt-1 text-xs text-muted-foreground">
                                View real-time power, voltage, current, and device status instantly.
                            </p>
                        </div>

                        <div class="rounded-2xl border bg-muted/20 p-4">
                            <p class="text-sm font-medium">Alerts & analytics</p>
                            <p class="mt-1 text-xs text-muted-foreground">
                                Detect anomalies early and review performance trends over time.
                            </p>
                        </div>
                    </div>
                </div>

                <p class="text-xs text-muted-foreground">
                    © {{ new Date().getFullYear() }} Voltris. All rights reserved.
                </p>
            </section>

            <!-- Right form panel -->
            <section class="flex items-center justify-center px-4 py-10 sm:px-6 lg:px-8">
                <div class="w-full max-w-md space-y-6">
                    <!-- Mobile branding -->
                    <div class="flex flex-col items-center gap-3 lg:hidden">
                        <div class="w-16">
                            <Logo />
                        </div>
                        <div class="text-center">
                            <p class="text-lg font-bold tracking-[0.2em]">VOLTRIS</p>
                            <p class="text-sm text-muted-foreground">Smart Energy Monitoring</p>
                        </div>
                    </div>

                    <Card class="rounded-2xl border bg-background/95 shadow-xl backdrop-blur">
                        <CardHeader class="space-y-2 pb-4">
                            <CardTitle class="text-2xl tracking-tight">Welcome back</CardTitle>
                            <CardDescription>
                                Sign in to continue to your dashboard.
                            </CardDescription>
                        </CardHeader>

                        <CardContent>
                            <form class="space-y-5" @submit.prevent="handleLogin">
                                <FieldSet>
                                    <FieldGroup class="space-y-4">
                                        <Field class="space-y-2">
                                            <FieldLabel for="username">Username</FieldLabel>
                                            <div class="relative">
                                                <UserIcon
                                                    class="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
                                                <Input id="username" v-model="form.username" type="text"
                                                    placeholder="Enter your username" class="pl-9"
                                                    autocomplete="username" />
                                            </div>
                                        </Field>

                                        <Field class="space-y-2">
                                            <FieldLabel for="password">Password</FieldLabel>
                                            <div class="relative">
                                                <LockKeyholeIcon
                                                    class="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
                                                <Input id="password" v-model="form.password" type="password"
                                                    placeholder="Enter your password" class="pl-9"
                                                    autocomplete="current-password" />
                                            </div>
                                        </Field>
                                    </FieldGroup>
                                </FieldSet>

                                <div v-if="errorMessage"
                                    class="rounded-xl border border-destructive/30 bg-destructive/5 px-3 py-2 text-sm text-destructive">
                                    {{ errorMessage }}
                                </div>

                                <Button type="submit" class="h-11 w-full rounded-xl" :disabled="isLoading">
                                    <span v-if="isLoading" class="inline-flex items-center gap-2">
                                        <Loader2Icon class="h-4 w-4 animate-spin" />
                                        Signing in...
                                    </span>
                                    <span v-else>Sign In</span>
                                </Button>

                                <p class="text-center text-xs text-muted-foreground">
                                    Secure access for authorized users only.
                                </p>
                            </form>
                        </CardContent>
                    </Card>
                </div>
            </section>
        </div>
    </main>
</template>