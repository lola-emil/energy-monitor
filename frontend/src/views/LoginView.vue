<script setup lang="ts">
import { reactive, ref } from 'vue';
import { useAuthStore } from '@/stores/auth';
import { useRouter } from 'vue-router';
import {
    Field, FieldGroup,
    FieldLabel,
    FieldSet
} from '@/components/ui/field';
import { Input } from '@/components/ui/input';
import Button from '@/components/ui/button/Button.vue';


const auth = useAuthStore()
const router = useRouter()

const form = reactive({
    username: "",
    password: ""
});

const isLoading = ref(false);

const onSubmit = async () => {
    await auth.login({
        username: form.username,
        password: form.password
    })

    if (auth.isAuthenticated) {
        router.push('/')
    }
}

</script>

<template>
    <main class="min-h-screen flex justify-center items-center">
        <form class="w-xs" @submit.prevent="onSubmit">
            <p class="text-lg font-semibold">Sign In</p>
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

            <Button class="w-full mt-5">
                <span v-if="isLoading">
                    <span class="loading loading-spinner loading-xs"></span>
                    Please wait..
                </span>
                <span v-else>Sign In</span>
            </Button>
        </form>
    </main>

</template>