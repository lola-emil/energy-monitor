<template>
    <Dialog :open="modal.isOpen" @update:open="onOpenChange">
        <form>

            <DialogContent class="sm:max-w-106.25">

                <DialogHeader>
                    <DialogTitle>Add New Device</DialogTitle>
                    <DialogDescription>
                    </DialogDescription>
                </DialogHeader>

                <Alert v-if="formError" variant="destructive">
                    <AlertCircleIcon />
                    <AlertTitle>{{ formError }}</AlertTitle>
                </Alert>

                <div class="grid gap-4">
                    <div class="grid gap-3">
                        <Label for="name-1">Device Name</Label>
                        <Input v-model="form.deviceName" />
                    </div>
                    <div class="grid gap-3">
                        <Label for="username-1">Device Code</Label>
                        <Input v-model="form.deviceCode" />
                    </div>
                </div>
                <DialogFooter>
                    <DialogClose as-child>
                        <Button variant="outline">
                            Cancel
                        </Button>
                    </DialogClose>
                    <Button type="submit" @click="submit">
                        Save changes
                    </Button>
                </DialogFooter>
            </DialogContent>
        </form>
    </Dialog>
</template>


<script setup lang="ts">
import { Button } from '@/components/ui/button';
import {
    Dialog,
    DialogClose,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle
} from '@/components/ui/dialog';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { useModalStore } from '@/stores/modal';

import { AlertCircleIcon } from 'lucide-vue-next';
import {
    Alert, AlertTitle
} from '@/components/ui/alert';
import { reactive, ref } from 'vue';
import { axiosInstance } from '@/api/axios';
import { useAuthStore } from '@/stores/auth';

const modal = useModalStore();
const auth = useAuthStore();

const form = reactive({
    deviceName: "",
    deviceCode: ""
})

const formError = ref<string | null>(null);

const submit = () => {
    axiosInstance.post("/api/device-claims", {
        device_name: form.deviceName,
        device_code: form.deviceCode,
    }, {
        headers: {
            "Authorization": `Bearer ${auth.token}`
        }
    }).then(res => {
        console.log(res.data);
    }).catch(err => {
        console.log(err.status);

        if (err.status == 400) {
            formError.value = "Invalid device or input";
        }

        if (err.status == 409) {
            formError.value = "Device not available";
        }

        if (err.status >= 500) {
            formError.value = "Server error";
        }
    })
}

function onOpenChange(value: boolean) {
    if (!value) modal.close()
}
</script>