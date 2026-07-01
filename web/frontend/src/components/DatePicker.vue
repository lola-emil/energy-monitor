<script setup lang="ts">
import { computed, ref } from 'vue'
import { Button } from '@/components/ui/button'
import {
    Popover,
    PopoverContent,
    PopoverTrigger,
} from '@/components/ui/popover'
import {
    Select, SelectTrigger, SelectValue,
    SelectContent,
    SelectItem,
} from '@/components/ui/select'

const selectedYear = ref(new Date().getFullYear())
const selectedMonth = ref<number | null>(null)

const months = [
    'Jan', 'Feb', 'Mar',
    'Apr', 'May', 'Jun',
    'Jul', 'Aug', 'Sep',
    'Oct', 'Nov', 'Dec',
]

const displayValue = computed(() => {
    if (selectedMonth.value === null) return 'Select month'

    return `${months[selectedMonth.value]} ${selectedYear.value}`
})

function selectMonth(index: number) {
    selectedMonth.value = index
}

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
                <Button v-for="(month, index) in months" :key="month" variant="outline" @click="selectMonth(index)">
                    {{ month }}
                </Button>
            </div>
        </PopoverContent>
    </Popover>
</template>
