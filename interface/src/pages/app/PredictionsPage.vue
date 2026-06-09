<template>
    <section class="app-page space-y-4">
        <div>
            <h1 class="text-2xl font-bold text-slate-900">
                Meus palpites
            </h1>

            <p class="mt-1 text-sm text-slate-500">
                Veja todos os palpites que você já salvou por grupo.
            </p>
        </div>

        <div class="app-card p-4">
            <label class="mb-2 block text-sm font-semibold text-slate-700">
                Grupo
            </label>

            <Dropdown v-model="selectedGroupId" :options="groups" optionLabel="name" optionValue="id"
                placeholder="Selecione um grupo" class="w-full" />
        </div>

        <Transition name="fade-slide" mode="out-in">
            <div v-if="loading" key="loading" class="space-y-3">
                <Skeleton v-for="item in 4" :key="item" height="110px" borderRadius="24px" />
            </div>

            <div v-else-if="predictionsWithMatch.length === 0" key="empty" class="app-card p-6 text-center">
                <p class="font-semibold text-slate-900">
                    Você ainda não salvou nenhum palpite.
                </p>

                <p class="mt-1 text-sm text-slate-500">
                    Vá para os jogos e comece a palpitar.
                </p>
            </div>
            <TransitionGroup v-else appear name="prediction-card" tag="div" class="space-y-4">
                <article v-for="prediction in predictionsWithMatch" :key="prediction.id" class="app-card p-4">
                    <div class="flex items-start justify-between gap-3">
                        <div>
                            <p class="text-xs font-medium uppercase text-emerald-600">
                                {{ stageLabel(prediction.match.stage) }}
                            </p>

                            <p class="mt-1 text-xs text-slate-500">
                                {{ formatDate(prediction.match.starts_at) }}
                            </p>
                        </div>

                        <span class="rounded-full bg-slate-100 px-3 py-1 text-xs font-semibold text-slate-600">
                            {{ statusLabel(prediction.match.status) }}
                        </span>
                    </div>

                    <div class="mt-4 grid grid-cols-[1fr_auto_1fr] items-center gap-3">
                        <div class="min-w-0 text-center">
                            <img v-if="prediction.match.home_team_flag_url" :src="prediction.match.home_team_flag_url"
                                class="mx-auto mb-2 h-8 w-8 object-contain" />

                            <p class="text-sm font-semibold text-slate-900">
                                {{ prediction.match.home_team_name || 'A definir' }}
                            </p>
                        </div>

                        <div class="rounded-2xl bg-slate-50 px-4 py-2 text-center">
                            <p class="text-lg font-bold text-slate-900 score-text">
                                {{ prediction.home_score }} x {{ prediction.away_score }}
                            </p>

                            <p class="mt-1 text-xs font-semibold" :class="prediction.points > 0
                                ? 'text-emerald-600'
                                : 'text-slate-500'
                                ">
                                {{ prediction.points || 0 }} pts
                            </p>
                        </div>

                        <div class="min-w-0 text-center">
                            <img v-if="prediction.match.away_team_flag_url" :src="prediction.match.away_team_flag_url"
                                class="mx-auto mb-2 h-8 w-8 object-contain" />

                            <p class="text-sm font-semibold text-slate-900">
                                {{ prediction.match.away_team_name || 'A definir' }}
                            </p>
                        </div>
                    </div>

                    <Button v-if="prediction.match.status === 'scheduled'" label="Editar palpite" icon="pi pi-pencil"
                        outlined class="mt-4 w-full" @click="goToMatch(prediction.match.id)" />
                </article>
            </TransitionGroup>
        </Transition>
    </section>
</template>

<script setup>
import { computed, onMounted, ref, watch } from 'vue'
import { useRouter } from 'vue-router'

import Button from 'primevue/button'
import Dropdown from 'primevue/dropdown'
import Skeleton from 'primevue/skeleton'

import api from '@/services/api'

import { useGroupStore } from '@/stores/group'
import { usePredictionStore } from '@/stores/prediction'
import { useMatchStore } from '@/stores/match'

import { stageLabel, statusLabel } from '@/utils/matchLabels'

const router = useRouter()

const groupStore = useGroupStore()
const predictionStore = usePredictionStore()
const matchStore = useMatchStore()

const selectedGroupId = ref(null)

const matches = computed(() => matchStore.matches)

const groups = computed(() => groupStore.groups)
const predictions = computed(() => predictionStore.myPredictions)
const loading = computed(() => predictionStore.loading)

const predictionsWithMatch = computed(() => {
    return predictions.value.map((prediction) => {
        const match = matches.value.find((match) => {
            return (
                match.id === prediction.match_id ||
                match.uuid === prediction.match_id ||
                match.match_id === prediction.match_id
            )
        })

        return {
            ...prediction,
            match,
        }
    })
})

async function loadGroups() {
    await groupStore.fetchGroups()

    if (groups.value.length > 0) {
        selectedGroupId.value = groups.value[0].uuid
    }
}

async function loadMatches() {
    await matchStore.fetchMatches()
}

async function loadPredictions() {
    if (!selectedGroupId.value) return

    await predictionStore.fetchMyPredictions(selectedGroupId.value)
}

function formatDate(value) {
    if (!value) return ''

    return new Date(value).toLocaleString('pt-BR', {
        dateStyle: 'short',
        timeStyle: 'short',
    })
}

function goToMatch(matchId) {
    router.push({
        path: '/app/matches',
        query: { match: matchId },
    })
}

watch(selectedGroupId, () => {
    loadPredictions()
})

onMounted(async () => {
    await Promise.all([
        loadGroups(),
        loadMatches(),
    ])
})
</script>
<style scoped>
.prediction-card-enter-active {
    transition: all 0.45s cubic-bezier(0.22, 1, 0.36, 1);
}

.prediction-card-leave-active {
    transition: all 0.25s ease;
}

.prediction-card-enter-from,
.prediction-card-leave-to {
    opacity: 0;
    transform: translateY(18px) scale(0.96);
}

.prediction-card-enter-to,
.prediction-card-leave-from {
    opacity: 1;
    transform: translateY(0) scale(1);
}

.prediction-card-move {
    transition: transform 0.45s cubic-bezier(0.22, 1, 0.36, 1);
}
</style>