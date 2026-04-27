<script setup lang="ts">
import { reactive, ref } from "vue"

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

const emit = defineEmits<{
  submit: [
    payload: {
      name: string
      location: string
      device_code: string
    }
  ]
}>()

const open = ref(false)
const isLoading = ref(false)
const errorMessage = ref("")

const form = reactive({
  name: "",
  location: "",
  device_code: "",
})

const resetForm = () => {
  form.name = ""
  form.location = ""
  form.device_code = ""
  errorMessage.value = ""
}

const handleSubmit = async () => {
  errorMessage.value = ""

  if (!form.name || !form.location || !form.device_code) {
    errorMessage.value = "All fields are required"
    return
  }

  if (!form.device_code.startsWith("EMS-")) {
    errorMessage.value = "Invalid device code format"
    return
  }

  isLoading.value = true

  try {
    emit("submit", {
      name: form.name.trim(),
      location: form.location.trim(),
      device_code: form.device_code.trim().toUpperCase(),
    })

    resetForm()
    open.value = false
  } finally {
    isLoading.value = false
  }
}
</script>

<template>
  <Dialog v-model:open="open">
    <DialogTrigger as-child>
      <Button>
        Add Appliance
      </Button>
    </DialogTrigger>

    <DialogContent class="sm:max-w-md">
      <DialogHeader>
        <DialogTitle>
          Register New Appliance
        </DialogTitle>

        <DialogDescription>
          Enter the protected device code to register your appliance.
        </DialogDescription>
      </DialogHeader>

      <FieldSet>
        <FieldGroup>

          <Field>
            <FieldLabel for="device_code">
              Device Code
            </FieldLabel>

            <Input
              id="device_code"
              v-model="form.device_code"
              placeholder="e.g. EMS-8F29A7XQ"
            />
          </Field>

          <Field>
            <FieldLabel for="name">
              Appliance Name
            </FieldLabel>

            <Input
              id="name"
              v-model="form.name"
              placeholder="e.g. Refrigerator"
            />
          </Field>

          <Field>
            <FieldLabel for="location">
              Location
            </FieldLabel>

            <Input
              id="location"
              v-model="form.location"
              placeholder="e.g. Kitchen"
            />
          </Field>

        </FieldGroup>
      </FieldSet>

      <p
        v-if="errorMessage"
        class="text-sm text-red-500"
      >
        {{ errorMessage }}
      </p>

      <DialogFooter>
        <Button
          variant="outline"
          @click="open = false"
        >
          Cancel
        </Button>

        <Button
          @click="handleSubmit"
          :disabled="isLoading"
        >
          {{ isLoading ? "Registering..." : "Register Appliance" }}
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>