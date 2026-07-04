<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { Button } from '@/components/ui/button'
import {
    Popover,
    PopoverContent,
    PopoverTrigger,
} from '@/components/ui/popover'
import {
    Select,
    SelectTrigger,
    SelectValue,
    SelectContent,
    SelectItem,
} from '@/components/ui/select'

export interface MonthPickerValue {
    year: number
    month: number | null
}

const model = defineModel<MonthPickerValue>({
    default: () => ({
        year: new Date().getFullYear(),
        month: null,
    }),
})

const emit = defineEmits<{
    (e: 'onSelect', value: MonthPickerValue): void
}>()

const selectedYear = ref(model.value.year)
const selectedMonth = ref<number | null>(model.value.month)

watch(model, (value) => {
    selectedYear.value = value.year
    selectedMonth.value = value.month
})
const months = [
    'Jan', 'Feb', 'Mar',
    'Apr', 'May', 'Jun',
    'Jul', 'Aug', 'Sep',
    'Oct', 'Nov', 'Dec',
]

const displayValue = computed(() => {
    if (selectedMonth.value == null) {
        return 'Select month'
    }

    return `${months[selectedMonth.value - 1]} ${selectedYear.value}`
})

function updateModel() {
    const value = {
        year: selectedYear.value,
        month: selectedMonth.value,
    }

    model.value = value
    emit('onSelect', value)
}

function selectMonth(index: number) {
    selectedMonth.value = index;
    updateModel()
}

watch(selectedYear, () => {
    updateModel()
})
</script>

<template>
    <Popover>
        <PopoverTrigger as-child>
            <Button variant="outline">
                {{ displayValue }}
            </Button>
        </PopoverTrigger>

        <PopoverContent class="w-72 space-y-4">
            <Select v-model="selectedYear">
                <SelectTrigger>
                    <SelectValue />
                </SelectTrigger>

                <SelectContent>
                    <SelectItem v-for="year in 20" :key="year" :value="2020 + year">
                        {{ 2020 + year }}
                    </SelectItem>
                </SelectContent>
            </Select>

            <div class="grid grid-cols-3 gap-2">
                <Button v-for="(month, index) in months" :key="month" variant="outline" @click="selectMonth(index + 1)">
                    {{ month }}
                </Button>
            </div>
        </PopoverContent>
    </Popover>
</template>