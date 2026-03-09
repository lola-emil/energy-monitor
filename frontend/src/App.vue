<script setup lang="ts">
import { RouterView } from 'vue-router';
import Sidebar from './components/Sidebar.vue';
</script>

<template>
  <Sidebar>
        <RouterView />
  </Sidebar>
</template>

<style scoped></style>
