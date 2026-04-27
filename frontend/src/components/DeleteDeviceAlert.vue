<script setup lang="ts">
import type { Appliance } from "@/types/appliance"

import {
    AlertDialog,
    AlertDialogAction,
    AlertDialogCancel,
    AlertDialogContent,
    AlertDialogDescription,
    AlertDialogFooter,
    AlertDialogHeader,
    AlertDialogTitle,
    AlertDialogTrigger,
} from "@/components/ui/alert-dialog"

import { Button } from "@/components/ui/button"

defineProps<{
    appliance: Appliance
}>()

const emit = defineEmits<{
    submit: [id: number]
}>()
</script>

<template>
    <AlertDialog>
        <AlertDialogTrigger as-child>
            <Button variant="destructive" size="sm">
                Delete
            </Button>
        </AlertDialogTrigger>

        <AlertDialogContent>
            <AlertDialogHeader>
                <AlertDialogTitle>
                    Delete Appliance?
                </AlertDialogTitle>

                <AlertDialogDescription>
                    This will permanently remove
                    <strong>{{ appliance.name }}</strong>
                    from monitoring.

                    This action cannot be undone.
                </AlertDialogDescription>
            </AlertDialogHeader>

            <AlertDialogFooter>
                <AlertDialogCancel>
                    Cancel
                </AlertDialogCancel>

                <AlertDialogAction @click="emit('submit', appliance.id)">
                    Confirm Delete
                </AlertDialogAction>
            </AlertDialogFooter>
        </AlertDialogContent>
    </AlertDialog>
</template>