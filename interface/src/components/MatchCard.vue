<template>
    <article v-motion :initial="{
        opacity: 0,
        y: 20,
        scale: 0.98
    }" :enter="{
    opacity: 1,
    y: 0,
    scale: 1,
    transition: {
        delay: props.delay,
        duration: 300
    }
}" class="rounded-3xl bg-white p-4 shadow-sm">
        <div class="flex items-center justify-between gap-3">
            <div>
                <p class="text-xs font-medium uppercase text-emerald-600">
                    {{ match.stage }}
                </p>
                <p class="mt-1 text-xs text-slate-500">
                    {{ formattedDate }}
                </p>
            </div>

            <span class="rounded-full bg-slate-100 px-3 py-1 text-xs font-semibold text-slate-600">
                {{ match.status }}
            </span>
        </div>

        <div class="mt-5 flex items-center justify-between gap-3">
            <div class="flex-1 text-center">
                <p class="text-sm font-semibold text-slate-900">
                    {{ match.home_team_name || 'A definir' }}
                </p>
            </div>

            <div class="rounded-2xl bg-slate-50 px-4 py-2 text-center">
                <p class="text-lg font-bold text-slate-900">
                    {{ scoreText }}
                </p>
            </div>

            <div class="flex-1 text-center">
                <p class="text-sm font-semibold text-slate-900">
                    {{ match.away_team_name || 'A definir' }}
                </p>
            </div>
        </div>

        <form v-if="match.status === 'scheduled'" class="mt-5 flex items-center gap-2" @submit.prevent="submit">
            <InputNumber v-model="form.home_score" inputClass="text-center" :min="0" fluid />

            <span class="font-bold text-slate-400">x</span>

            <InputNumber v-model="form.away_score" inputClass="text-center" :min="0" fluid />

            <Button type="submit" icon="pi pi-check" :loading="loading" rounded class="!h-11 !w-11 shrink-0" />
        </form>
    </article>
</template>

<script setup>
import { computed, reactive } from 'vue'
import Button from 'primevue/button'
import InputNumber from 'primevue/inputnumber'

const props = defineProps({
    match: {
        type: Object,
        required: true,
    },
    groupId: {
        type: Number,
        required: true,
    },
    loading: {
        type: Boolean,
        default: false,
    },
    delay: {
        type: Number,
        default: 0,
    },
})

const emit = defineEmits(['save'])

const form = reactive({
    home_score: 0,
    away_score: 0,
})

const formattedDate = computed(() => {
    if (!props.match.starts_at) return ''

    return new Date(props.match.starts_at).toLocaleString('pt-BR', {
        dateStyle: 'short',
        timeStyle: 'short',
    })
})

const scoreText = computed(() => {
    if (props.match.home_score === null || props.match.away_score === null) {
        return 'x'
    }

    return `${props.match.home_score} x ${props.match.away_score}`
})

function submit() {
    emit('save', {
        group_id: props.groupId,
        match_id: props.match.internal_id || props.match.match_id || props.match.id,
        home_score: form.home_score,
        away_score: form.away_score,
    })
}
</script>