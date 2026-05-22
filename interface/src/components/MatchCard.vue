<template>
    <article :id="`match-${match.id}`"
     v-motion="animated ? undefined : false"
        :initial="animated ? { opacity: 0, y: 14, scale: 0.98 } : undefined"
        :enter="animated ? { opacity: 1, y: 0, scale: 1, transition: { delay, duration: 220 } } : undefined"
        class="app-card p-4">
        <div class="flex items-center justify-between gap-3">
            <div>
                <p class="text-xs font-medium uppercase text-emerald-600">
                    {{ stageLabel(match.stage) }}
                </p>
                <p class="mt-1 text-xs text-slate-500">
                    {{ formattedDate }}
                </p>
            </div>

            <span class="rounded-full bg-slate-100 px-3 py-1 text-xs font-semibold text-slate-600">
                {{ statusLabel(match.status) }}
            </span>
        </div>

        <div class="mt-5 flex items-center justify-between gap-3 mr-12">
            <div class="flex-1 text-center">
                <img v-if="match.home_team_flag_url" :src="match.home_team_flag_url" loading="lazy" decoding="async"
                    class="mx-auto mb-2 h-10 w-10 object-contain" />
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
                <img v-if="match.away_team_flag_url" :src="match.away_team_flag_url" loading="lazy" decoding="async"
                    class="mx-auto mb-2 h-10 w-10 object-contain" />
                <p class="text-sm font-semibold text-slate-900">
                    {{ match.away_team_name || 'A definir' }}
                </p>
            </div>
        </div>

        <form v-if="canPredict" class="mt-5 grid grid-cols-[1fr_auto_1fr_auto] items-center gap-2"
            @submit.prevent="submit">
            <InputNumber v-model="form.home_score" inputClass="text-center" :min="0" fluid />

            <span class="font-bold text-slate-400">x</span>

            <InputNumber v-model="form.away_score" inputClass="text-center" :min="0" fluid />

            <Button type="submit" icon="pi pi-check" :loading="loading" rounded class="!h-11 !w-11 shrink-0" />
        </form>

        <div v-else class="mt-5 grid grid-cols-2 gap-3">
            <div class="rounded-2xl bg-slate-50 p-3 text-center">
                <p class="text-xs text-slate-500">Seu palpite</p>
                <p class="mt-1 text-lg font-bold text-slate-900">
                    {{ predictionScore }}
                </p>
            </div>

            <div class="rounded-2xl bg-emerald-50 p-3 text-center">
                <p class="text-xs text-emerald-700">Resultado</p>
                <p class="mt-1 text-lg font-bold text-emerald-700">
                    {{ realScore }}
                </p>
            </div>
        </div>
    </article>
</template>

<script setup>
import { computed, reactive, watch } from 'vue'
import Button from 'primevue/button'
import InputNumber from 'primevue/inputnumber'
import { stageLabel, statusLabel } from '@/utils/matchLabels'

const props = defineProps({
    match: {
        type: Object,
        required: true,
    },
    groupId: {
        type: String,
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
    animated: {
        type: Boolean,
        default: true,
    },
    prediction: {
        type: Object,
        default: null,
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

watch(
    () => props.prediction,
    (prediction) => {
        if (prediction) {
            form.home_score = prediction.home_score
            form.away_score = prediction.away_score
        }
    },
    { immediate: true }
)

const canPredict = computed(() => {
    return props.match.status === 'scheduled'
})

const predictionScore = computed(() => {
    if (!props.prediction) return 'Sem palpite'

    return `${props.prediction.home_score} x ${props.prediction.away_score}`
})

const realScore = computed(() => {
    if (
        props.match.home_score === null ||
        props.match.away_score === null
    ) {
        return 'Aguardando'
    }

    return `${props.match.home_score} x ${props.match.away_score}`
})
</script>