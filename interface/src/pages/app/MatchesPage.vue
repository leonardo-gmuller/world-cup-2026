<template>
    <section class="app-page">
        <div class="flex items-start justify-between gap-3">
            <div>
                <h1 class="app-title">Jogos</h1>
                <p class="app-subtitle">
                    Veja os jogos e envie seus palpites.
                </p>
            </div>

            <Button icon="pi pi-refresh" rounded text :loading="matchStore.loading" @click="matchStore.importAll" />
        </div>

        <div class="rounded-3xl bg-white p-4 shadow-sm">
            <label class="text-xs font-semibold text-slate-500">
                Grupo selecionado
            </label>

            <Select v-model="selectedGroupId" :options="groupStore.groups" optionLabel="name" optionValue="id"
                placeholder="Selecione um grupo" class="w-full" :loading="groupStore.loading" />
        </div>

        <div class="flex gap-2 overflow-x-auto pb-1">
            <button v-for="stage in stages" :key="stage.value" type="button" class="
    shrink-0 rounded-full px-4 py-2 text-sm font-semibold cursor-pointer
    transform-gpu transition-all duration-150
    active:scale-95
  " style="touch-action: manipulation;" :class="selectedStage === stage.value
    ? 'bg-emerald-600 text-white shadow-md'
    : 'bg-white text-slate-600 shadow-sm'
    " @click="selectedStage = stage.value">
                {{ stage.label }}
            </button>
        </div>

        <div v-if="matchStore.loading" class="space-y-3">
            <article v-for="i in 6" :key="i" class="app-card p-4">
                <div class="flex items-center justify-between">
                    <div>
                        <Skeleton width="7rem" height="0.8rem" />
                        <Skeleton width="5rem" height="0.7rem" class="mt-2" />
                    </div>

                    <Skeleton width="4rem" height="1.5rem" borderRadius="999px" />
                </div>

                <div class="mt-5 flex items-center justify-between gap-3">
                    <div class="flex-1 text-center">
                        <Skeleton shape="circle" size="2.5rem" class="mx-auto" />
                        <Skeleton width="5rem" height="0.8rem" class="mx-auto mt-2" />
                    </div>

                    <Skeleton width="4rem" height="2.5rem" borderRadius="1rem" />

                    <div class="flex-1 text-center">
                        <Skeleton shape="circle" size="2.5rem" class="mx-auto" />
                        <Skeleton width="5rem" height="0.8rem" class="mx-auto mt-2" />
                    </div>
                </div>

                <div class="mt-5 grid grid-cols-[1fr_auto_1fr_auto] items-center gap-2">
                    <Skeleton height="2.75rem" borderRadius="0.75rem" />
                    <span class="font-bold text-slate-300">x</span>
                    <Skeleton height="2.75rem" borderRadius="0.75rem" />
                    <Skeleton shape="circle" size="2.75rem" />
                </div>
            </article>
        </div>

        <div v-else class="space-y-3">
            <MatchCard v-for="(match, index) in filteredMatches" :key="match.id" :match="match"
                :group-id="selectedGroupId" :delay="index < 12 ? index * 40 : 0" :animated="index < 12"
                :loading="predictionStore.loading" @save="savePrediction" :prediction="predictionsByMatch[match.id]" />
        </div>
    </section>
</template>

<script setup>
import { onMounted, ref, computed, watch, nextTick } from 'vue'
import Button from 'primevue/button'
import Select from 'primevue/select'
import MatchCard from '@/components/MatchCard.vue'
import { Skeleton } from 'primevue'
import { useMatchStore } from '@/stores/match'
import { useGroupStore } from '@/stores/group'
import { usePredictionStore } from '@/stores/prediction'
import { useAppToast } from '@/utils/toast'
import { useRoute } from 'vue-router'

const route = useRoute()

const matchStore = useMatchStore()
const groupStore = useGroupStore()
const predictionStore = usePredictionStore()
const toast = useAppToast()

const selectedGroupId = ref("")

const selectedStage = ref('all')

const stages = [
    { label: 'Todos', value: 'all' },
    { label: 'Grupos', value: 'group_stage' },
    { label: 'Fase eliminatória', value: 'round_of_32' },
    { label: 'Oitavas', value: 'round_of_16' },
    { label: 'Quartas', value: 'quarter_final' },
    { label: 'Semi', value: 'semi_final' },
    { label: 'Final', value: 'final' },
]

const filteredMatches = computed(() => {
    if (selectedStage.value === 'all') {
        return matchStore.matches
    }

    return matchStore.matches.filter((match) => {
        return match.stage === selectedStage.value
    })
})

onMounted(async () => {
    await Promise.all([
        matchStore.fetchMatches(),
        groupStore.fetchGroups(),     
    ])

     await scrollToMatch()

    if (groupStore.groups.length > 0) {
        selectedGroupId.value = groupStore.groups[0].internal_id || groupStore.groups[0].id
    }
})

async function savePrediction(payload) {
    if (!selectedGroupId.value) return

    try {
        await predictionStore.save({
            ...payload,
            group_id: selectedGroupId.value,
        })

        toast.success('Palpite salvo com sucesso.')
    } catch (error) {
        toast.error(error, 'Erro ao salvar palpite')
    }
}


const predictionsByMatch = computed(() => {
    const map = {}

    predictionStore.predictions.forEach((prediction) => {
        map[prediction.match_id] = prediction
    })

    return map
})

watch(selectedGroupId, async (groupId) => {
    if (groupId) {
        await predictionStore.fetchByGroup(groupId)
    }
})

async function scrollToMatch() {
    const matchId = route.query.match

    if (!matchId) return

    await nextTick()

    const element = document.getElementById(`match-${matchId}`)

    console.log('Scrolling to match', matchId, element)

    if (element) {
        element.scrollIntoView({
            behavior: 'smooth',
            block: 'start',
        })
    }
}
</script>