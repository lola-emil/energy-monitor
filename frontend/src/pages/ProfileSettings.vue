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

const isLoading = ref(false)
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

        form.username = 'admin'
    } catch (error) {
        errorMessage.value = 'Failed to load user settings'
    }
}

const handleSave = async () => {
    errorMessage.value = ''
    successMessage.value = ''

    if (!form.username.trim()) {
        errorMessage.value = 'Username is required'
        return
    }

    if (form.new_password && form.new_password !== form.confirm_password) {
        errorMessage.value = 'New password and confirm password do not match'
        return
    }

    try {
        isLoading.value = true

        successMessage.value = 'Credentials updated successfully'

        form.current_password = ''
        form.new_password = ''
        form.confirm_password = ''
    } catch (error: any) {
        errorMessage.value =
            error?.response?.data?.message ||
            'Failed to update credentials'
    } finally {
        isLoading.value = false
    }
}

onMounted(() => {
    fetchProfile()
})
</script>

<template>
    <div class="px-5 my-5">
        <div class="space-y-6">
            <div>
                <h1 class="text-2xl font-semibold">Account Settings</h1>
                <p class="text-sm text-muted-foreground mt-1">
                    Update your username and password credentials.
                </p>
            </div>

            <form @submit.prevent="handleSave" class="space-y-6">
                <FieldSet>
                    <FieldGroup>
                        <Field>
                            <FieldLabel for="username">
                                Username
                            </FieldLabel>
                            <Input id="username" v-model="form.username" placeholder="Enter username" />
                        </Field>

                        <Field>
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
                    </FieldGroup>
                </FieldSet>

                <p v-if="errorMessage" class="text-sm text-destructive">
                    {{ errorMessage }}
                </p>

                <p v-if="successMessage" class="text-sm text-green-600">
                    {{ successMessage }}
                </p>

                <Button type="submit" :disabled="isLoading">
                    {{ isLoading ? 'Saving...' : 'Save Changes' }}
                </Button>
            </form>
        </div>
    </div>
</template>
