<script setup lang="ts"></script>

<template>
    <main class="min-h-screen flex justify-center items-center">
        <form class="w-xs">
            <p class="text-lg font-semibold">Sign In</p>
            <br>
            <fieldset class="fieldset">
                <legend class="fieldset-legend">Email</legend>
                <input type="text" class="input" />
                <p class="label"></p>
            </fieldset>

            <fieldset class="fieldset">
                <legend class="fieldset-legend">Password</legend>
                <input type="password" class="input" />
                <p class="label"></p>
            </fieldset>

            <button class="btn btn-primary w-full">Sign In</button>
        </form>
    </main>
</template>