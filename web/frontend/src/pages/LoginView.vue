<template>
    <main class="min-h-screen flex items-center justify-center bg-base-200 p-4">
        <div class="card w-full max-w-md bg-base-100 shadow-xl">
            <div class="card-body">

                <h2 class="card-title text-2xl font-bold justify-center mb-1">Welcome Back</h2>
                <p class="text-center text-base-content/70 mb-6">Please login to your account</p>

                <form @submit.prevent="handleLogin">

                    <fieldset class="fieldset">
                        <legend class="fieldset-legend">Email address</legend>
                        <input type="text" v-model="form.username" class="input w-full" placeholder="you@example.com"
                            required />
                    </fieldset>

                    <fieldset class="fieldset">
                        <legend class="fieldset-legend">Password</legend>
                        <input type="password" v-model="form.password" class="input w-full" placeholder="••••••••"
                            required />
                    </fieldset>

                    <div class="form-control mt-6">
                        <button type="submit" class="btn btn-primary w-full">
                            Login
                        </button>
                    </div>
                </form>
            </div>
        </div>
    </main>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'

import { useAuthStore } from '@/stores/auth'
import { authService } from '@/services/auth.service'

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