<script setup lang="ts">
import { reactive, ref } from 'vue';
import { signIn } from "@/api/auth"


const form = reactive({
    email: "",
    password: ""
});

const isLoading = ref(false);

const onSubmit = () => {
    isLoading.value = true;
    signIn(form).then(res => {
        console.log(res);
    })
    .catch(err => {
        console.log(err);
    }).finally(() => {
        isLoading.value = false;
    });
}
</script>

<template>
    <main class="min-h-screen flex justify-center items-center">
        <form class="w-xs" @submit.prevent="onSubmit">
            <p class="text-lg font-semibold">Sign In</p>
            <br>
            <fieldset class="fieldset">
                <legend class="fieldset-legend">Email</legend>
                <input type="text" class="input" v-model="form.email"/>
                <p class="label"></p>
            </fieldset>

            <fieldset class="fieldset">
                <legend class="fieldset-legend">Password</legend>
                <input type="password" class="input" v-model="form.password"/>
                <p class="label"></p>
            </fieldset>

            <button class="btn btn-primary w-full" :disabled="isLoading">
                <span v-if="isLoading">
                    <span class="loading loading-spinner loading-xs"></span>
                    Please wait..
                </span>
                <span v-else>Sign In</span>
            </button>
        </form>
    </main>
</template>