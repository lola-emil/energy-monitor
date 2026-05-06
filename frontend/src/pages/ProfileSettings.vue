<script setup lang="ts">
import { reactive, ref, onMounted } from 'vue'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import {
    Field,
    FieldGroup,
    FieldLabel,
    FieldSet,
} from '@/components/ui/field'
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from '@/components/ui/card'
import { Skeleton } from '@/components/ui/skeleton'
import {
    UserIcon,
    KeyRoundIcon,
    SaveIcon,
    ShieldCheckIcon,
} from 'lucide-vue-next'

const isLoading = ref(false)
const isFetching = ref(true)
const successMessage = ref('')
const errorMessage = ref('')

const form = reactive({
    username: '',
    current_password: '',
    new_password: '',
    confirm_password: '',
})

const fetchProfile = async () => {
    try {
        isFetching.value = true
        errorMessage.value = ''
        form.username = 'admin'
    } catch (error) {
        errorMessage.value = 'Failed to load user settings.'
    } finally {
        isFetching.value = false
    }
}

const handleSave = async () => {
    errorMessage.value = ''
    successMessage.value = ''

    if (!form.username.trim()) {
        errorMessage.value = 'Username is required.'
        return
    }

    if (form.new_password && !form.current_password) {
        errorMessage.value = 'Current password is required to set a new password.'
        return
    }

    if (form.new_password && form.new_password !== form.confirm_password) {
        errorMessage.value = 'New password and confirm password do not match.'
        return
    }

    try {
        isLoading.value = true

        // TODO: Replace with real API call
        await new Promise(resolve => setTimeout(resolve, 700))

        successMessage.value = 'Credentials updated successfully.'
        form.current_password = ''
        form.new_password = ''
        form.confirm_password = ''
    } catch (error: any) {
        errorMessage.value =
            error?.response?.data?.message ||
            'Failed to update credentials.'
    } finally {
        isLoading.value = false
    }
}

onMounted(fetchProfile)
</script>

<template>
    <div class="min-h-screen bg-muted/20">
        <div class="mx-auto w-full max-w-4xl space-y-6 px-4 py-6 sm:px-6 lg:px-8">
            <!-- Header -->
            <section class="rounded-2xl border bg-background px-5 py-5 shadow-sm">
                <div class="space-y-1">
                    <h1 class="text-2xl font-semibold tracking-tight">Account Settings</h1>
                    <p class="text-sm text-muted-foreground">
                        Update your username and password credentials for this account.
                    </p>
                </div>
            </section>

            <!-- Feedback -->
            <div v-if="errorMessage"
                class="rounded-xl border border-destructive/30 bg-destructive/5 px-4 py-3 text-sm text-destructive">
                {{ errorMessage }}
            </div>

            <div v-if="successMessage"
                class="rounded-xl border border-emerald-500/30 bg-emerald-500/5 px-4 py-3 text-sm text-emerald-700 dark:text-emerald-400">
                {{ successMessage }}
            </div>

            <!-- Loading -->
            <section v-if="isFetching">
                <Card class="rounded-2xl shadow-sm">
                    <CardHeader class="space-y-2">
                        <Skeleton class="h-5 w-40" />
                        <Skeleton class="h-4 w-72" />
                    </CardHeader>
                    <CardContent class="space-y-4">
                        <Skeleton class="h-16 w-full rounded-xl" />
                        <Skeleton class="h-16 w-full rounded-xl" />
                        <Skeleton class="h-16 w-full rounded-xl" />
                        <Skeleton class="h-16 w-full rounded-xl" />
                    </CardContent>
                </Card>
            </section>

            <!-- Form -->
            <form v-else @submit.prevent="handleSave" class="space-y-6">
                <Card class="rounded-2xl shadow-sm">
                    <CardHeader class="border-b pb-4">
                        <div class="flex items-start gap-3">
                            <div class="rounded-xl bg-primary/10 p-2 text-primary">
                                <ShieldCheckIcon class="h-4 w-4" />
                            </div>
                            <div>
                                <CardTitle>Login Credentials</CardTitle>
                                <CardDescription>
                                    Manage the credentials used to access the dashboard.
                                </CardDescription>
                            </div>
                        </div>
                    </CardHeader>

                    <CardContent class="space-y-6 p-5">
                        <FieldSet>
                            <FieldGroup class="space-y-6">
                                <!-- Username -->
                                <div class="rounded-xl border bg-muted/20 p-4">
                                    <div class="mb-4 flex items-center gap-2">
                                        <UserIcon class="h-4 w-4 text-primary" />
                                        <p class="text-sm font-medium">Username</p>
                                    </div>

                                    <Field>
                                        <FieldLabel for="username">Username</FieldLabel>
                                        <Input id="username" v-model="form.username" placeholder="Enter username" />
                                    </Field>
                                </div>

                                <!-- Password -->
                                <div class="rounded-xl border bg-muted/20 p-4">
                                    <div class="mb-4 flex items-center gap-2">
                                        <KeyRoundIcon class="h-4 w-4 text-primary" />
                                        <p class="text-sm font-medium">Password</p>
                                    </div>

                                    <div class="grid gap-4 md:grid-cols-2">
                                        <Field class="md:col-span-2">
                                            <FieldLabel for="current_password">
                                                Current Password
                                            </FieldLabel>
                                            <Input id="current_password" v-model="form.current_password" type="password"
                                                placeholder="Enter current password" />
                                        </Field>

                                        <Field>
                                            <FieldLabel for="new_password">
                                                New Password
                                            </FieldLabel>
                                            <Input id="new_password" v-model="form.new_password" type="password"
                                                placeholder="Enter new password" />
                                        </Field>

                                        <Field>
                                            <FieldLabel for="confirm_password">
                                                Confirm New Password
                                            </FieldLabel>
                                            <Input id="confirm_password" v-model="form.confirm_password" type="password"
                                                placeholder="Confirm new password" />
                                        </Field>
                                    </div>

                                    <p class="mt-3 text-xs text-muted-foreground">
                                        Leave the password fields blank if you only want to update the username.
                                    </p>
                                </div>
                            </FieldGroup>
                        </FieldSet>
                    </CardContent>
                </Card>

                <!-- Actions -->
                <div class="flex justify-end">
                    <Button type="submit" class="gap-2 rounded-xl" :disabled="isLoading">
                        <SaveIcon class="h-4 w-4" />
                        {{ isLoading ? 'Saving...' : 'Save Changes' }}
                    </Button>
                </div>
            </form>
        </div>
    </div>
</template>