import { axiosInstance } from '@/api/axios'
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

interface User {
    id: number
    username: string
}

interface LoginCredentials {
    username: string
    password: string
}

export const useAuthStore = defineStore('auth', () => {
    const user = ref<User | null>(null)
    const token = ref<string | null>(localStorage.getItem('token'))
    const loading = ref(false)
    const error = ref<string | null>(null)

    const isAuthenticated = computed(() => !!token.value)

    const login = async (credentials: LoginCredentials): Promise<void> => {
        loading.value = true
        error.value = null

        try {
            const response = await axiosInstance.post('/auth/login', credentials)

            token.value = response.data.token
            user.value = response.data.user

            localStorage.setItem('token', token.value!)

            axiosInstance.defaults.headers.common['Authorization'] = `Bearer ${token.value}`
        } catch (err: any) {
            error.value = err.response?.data?.message || 'Login failed'
        } finally {
            loading.value = false
        }
    }

    const logout = (): void => {
        user.value = null
        token.value = null
        error.value = null

        localStorage.removeItem('token')
        delete axiosInstance.defaults.headers.common['Authorization']
    }

    return {
        user,
        token,
        loading,
        error,

        isAuthenticated,

        login,
        logout
    }
})