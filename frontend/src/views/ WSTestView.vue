<script setup lang="ts">
import { useWebSocket } from '@/composables/useWebSocket';
import { onMounted } from 'vue';

const socket = useWebSocket();

onMounted(() => {
    socket.connect();
})

</script>

<template>

    <div>
        <ul>
            <li v-for="value in socket.messages.value">{{ value }}</li>
        </ul>
    </div>
</template>