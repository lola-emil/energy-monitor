<script setup lang="ts">
import type { Alert } from '@/services/alert.service';
import { Table, TableHeader, TableRow, TableHead, TableCell, TableBody } from "@/components/ui/table"
import { formatTime } from '@/lib/time';

defineProps<{
    alerts: Alert[]
    loading: boolean
}>()

const severityClass = (s: string) => {
    switch (s) {
        case 'high':
            return 'bg-red-500 text-white px-2 py-1 rounded'
        case 'medium':
            return 'bg-yellow-500 text-black px-2 py-1 rounded'
        default:
            return 'bg-muted text-muted-foreground px-2 py-1 rounded'
    }
}
</script>

<template>
    <Table>
        <TableHeader>
            <TableRow>
                <TableHead>Message</TableHead>
                <TableHead>Severity</TableHead>
                <TableHead>Timestamp</TableHead>
            </TableRow>
        </TableHeader>

        <TableBody>
            <TableRow v-if="loading">
                <TableCell colspan="3">Loading...</TableCell>
            </TableRow>

            <TableRow v-for="a in alerts" :key="a.id">
                <TableCell>{{ a.message }}</TableCell>

                <TableCell>
                    <span :class="severityClass(a.severity)">
                        {{ a.severity.toUpperCase() }}
                    </span>
                </TableCell>

                <TableCell>{{ formatTime(a.triggered_at) }}</TableCell>
            </TableRow>

            <TableRow v-if="!alerts.length && !loading">
                <TableCell colspan="3" class="text-center">
                    No alerts found
                </TableCell>
            </TableRow>
        </TableBody>
    </Table>
</template>