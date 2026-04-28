<style>
#split-screen-image {
    background-image: url(/Boyshet.png);
    background-size: cover;
    background-position: right;
    background-repeat: no-repeat;
}
</style>

<template>
    <main class="min-h-screen grid grid-cols-1 lg:grid-cols-2">
        <div class="flex justify-center items-center">
            <form class="w-xs" @submit.prevent="handleLogin">
                <div class="w-full flex justify-center flex-col items-center gap-3">
                    <div class="w-16">
                        <Logo />
                    </div>
                    <span class="tracking-wider text-lg font-bold">VOLTRIS</span>
                </div>
                <br>
                <p class="text-lg font-semibold">Welcome Back</p>
                <span class="text-muted-foreground text-sm">Sign in to continue to your dashboard</span>
                <br>
                <br>
                <FieldSet>
                    <FieldGroup>
                        <Field>
                            <FieldLabel for="username">
                                Username
                            </FieldLabel>
                            <Input id="username" v-model="form.username" type="text" placeholder="" />

                        </Field>
                        <Field>
                            <FieldLabel for="password">
                                Password
                            </FieldLabel>

                            <Input id="password" type="password" v-model="form.password" placeholder="********" />
                        </Field>
                    </FieldGroup>
                </FieldSet>

                <p v-if="errorMessage" class="text-sm text-red-500 mt-3">
                    {{ errorMessage }}
                </p>
                <Button type="submit" class="w-full mt-5" :disabled="isLoading">
                    <span v-if="isLoading">
                        <span class="loading loading-spinner loading-xs"></span>
                        Please wait..
                    </span>
                    <span v-else>Sign In</span>
                </Button>
            </form>
        </div>
        <div class="p-5 hidden lg:flex">
            <div class="rounded-lg bg-accent h-full w-full" id="split-screen-image">
            </div>
        </div>
    </main>

</template>


<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'

import {
    Field,
    FieldGroup,
    FieldLabel,
    FieldSet
} from '@/components/ui/field'

import { Input } from '@/components/ui/input'
import Button from '@/components/ui/button/Button.vue'

import { useAuthStore } from '@/stores/auth'
import { authService } from '@/services/auth.service'
import { Card, CardContent } from '@/components/ui/card'
import Logo from '@/components/Logo.vue'

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
        errorMessage.value = 'Username and password are required'
        return
    }

    isLoading.value = true

    try {
        const response = await authService.login(
            form.username,
            form.password
        )

        authStore.setToken(response.token)

        router.push('/')
    } catch (error: any) {
        errorMessage.value =
            error?.response?.data?.message ||
            'Invalid username or password'
    } finally {
        isLoading.value = false
    }
}
</script>