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

            <Select v-model="selectedGroupId" :options="groupStore.groups" optionLabel="name" optionValue="internal_id"
                placeholder="Selecione um grupo" class="mt-2 w-full" />
        </div>

        <div class="space-y-3">
            <MatchCard v-for="(match, index) in matchStore.matches" :key="match.id" :match="match"
                :group-id="selectedGroupId" :loading="predictionStore.loading" :delay="index * 80"
                @save="savePrediction" />
        </div>
    </section>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import Button from 'primevue/button'
import Select from 'primevue/select'
import MatchCard from '@/components/MatchCard.vue'
import { useMatchStore } from '@/stores/match'
import { useGroupStore } from '@/stores/group'
import { usePredictionStore } from '@/stores/prediction'

const matchStore = useMatchStore()
const groupStore = useGroupStore()
const predictionStore = usePredictionStore()

const selectedGroupId = ref(null)

onMounted(async () => {
    await Promise.all([
        matchStore.fetchMatches(),
        groupStore.fetchGroups(),
    ])

    if (groupStore.groups.length > 0) {
        selectedGroupId.value = groupStore.groups[0].internal_id || groupStore.groups[0].id
    }
})

async function savePrediction(payload) {
    if (!selectedGroupId.value) return

    await predictionStore.save({
        ...payload,
        group_id: selectedGroupId.value,
    })
}
</script>