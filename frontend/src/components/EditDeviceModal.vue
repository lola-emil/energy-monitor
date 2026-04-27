<script setup lang="ts">
import { reactive, ref, watch } from "vue"
import type { Appliance } from "@/types/appliance"

import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog"

import {
    Field,
    FieldGroup,
    FieldLabel,
    FieldSet
} from "@/components/ui/field"

import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"

const props = defineProps<{
    appliance: Appliance
}>()

const emit = defineEmits<{
    submit: [
        payload: {
            id: number
            name: string
            location: string
        }
    ]
}>()

const open = ref(false)
const isLoading = ref(false)
const errorMessage = ref("")

const form = reactive({
    name: "",
    location: "",
})

watch(
    () => props.appliance,
    (value) => {
        if (value) {
            form.name = value.name
            form.location = value.location
        }
    },
    { immediate: true }
)

const handleSubmit = async () => {
    errorMessage.value = ""

    if (!form.name || !form.location) {
        errorMessage.value = "All fields are required"
        return
    }

    isLoading.value = true

    try {
        emit("submit", {
            id: props.appliance.id,
            name: form.name.trim(),
            location: form.location.trim(),
        })

        open.value = false
    } finally {
        isLoading.value = false
    }
}
</script>

<template>
    <Dialog v-model:open="open">
        <DialogTrigger as-child>
            <Button variant="outline" size="sm">
                Edit
            </Button>
        </DialogTrigger>

        <DialogContent class="sm:max-w-md">
            <DialogHeader>
                <DialogTitle>
                    Edit Appliance
                </DialogTitle>

                <DialogDescription>
                    Update appliance information.
                </DialogDescription>
            </DialogHeader>

            <FieldSet>
                <FieldGroup>

                    <Field>
                        <FieldLabel>
                            Appliance Name
                        </FieldLabel>

                        <Input v-model="form.name" />
                    </Field>

                    <Field>
                        <FieldLabel>
                            Location
                        </FieldLabel>

                        <Input v-model="form.location" />
                    </Field>

                </FieldGroup>
            </FieldSet>

            <p v-if="errorMessage" class="text-sm text-red-500">
                {{ errorMessage }}
            </p>

            <DialogFooter>
                <Button variant="outline" @click="open = false">
                    Cancel
                </Button>

                <Button @click="handleSubmit" :disabled="isLoading">
                    {{ isLoading ? "Saving..." : "Save Changes" }}
                </Button>
            </DialogFooter>
        </DialogContent>
    </Dialog>
</template>