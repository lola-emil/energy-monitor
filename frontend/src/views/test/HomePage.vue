<template>
    <div class="h-16 w-full">
    </div>

    <main class="px-5">
        <p class="text-2xl">My Home</p>
        <div class="mt-3">
            <div class="text-sm flex gap-3 text-accent  items-center">
                <ShieldCheck :size="16" />
                <span>No critical issues detected</span>
            </div>
        </div>
        <br>
        <Separator />
        <br>
        <section>
            <div>
                <p class="text-lg">My Devices</p>
                <span class="text-muted-foreground">4 devices</span>
            </div>
        </section>
    </main>
</template>

<script setup lang="ts">
import { ShieldCheck } from 'lucide-vue-next';
import {
    Card,
    CardAction,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from '@/components/ui/card'
import Separator from '@/components/ui/separator/Separator.vue';
</script>