<template>
    <section class="app-page">
        <div class="flex items-center gap-3">
            <img :src="logo" alt="Bolão Copa" class="h-14 w-14 object-contain" />

            <div>
                <p class="app-subtitle">
                    Copa do Mundo 2026
                </p>

                <h1 class="app-title">
                    Bolão Copa 2026
                </h1>
            </div>
        </div>

        <div class="rounded-3xl bg-emerald-600 p-5 text-white shadow-sm md:p-8">
            <p class="text-sm opacity-90">Próximo jogo</p>
            <div class="mt-3 flex items-center justify-between gap-4">
                <div>
                    <h2 class="text-xl font-bold md:text-3xl">
                        {{ nextMatch?.home_team_name || 'A definir' }}
                        x
                        {{ nextMatch?.away_team_name || 'A definir' }}
                    </h2>

                    <p class="mt-1 text-sm opacity-90">
                        {{ nextMatch ? new Date(nextMatch.starts_at).toLocaleString('pt-BR') : 'Nenhum jogo encontrado'
                        }}
                    </p>
                </div>

                <RouterLink to="/app/matches"
                    class="app-card-button rounded-2xl bg-white/15 px-4 py-2 text-sm font-semibold">
                    Palpitar
                </RouterLink>
            </div>
        </div>

        <div class="grid grid-cols-3 gap-3">
            <div
                class="col-span-3 rounded-3xl bg-gradient-to-br from-emerald-500 to-emerald-600 p-5 text-white shadow-sm">
                <p class="text-sm opacity-90">Sua posição</p>

                <div class="mt-2 flex items-end justify-between">
                    <div>
                        <h2 class="text-5xl font-black">#3</h2>
                        <p class="mt-1 text-sm opacity-90">
                            entre 24 participantes
                        </p>
                    </div>

                    <div class="rounded-2xl bg-white/15 px-4 py-3 backdrop-blur-sm">
                        <p class="text-xs opacity-80">Pontos</p>
                        <p class="text-2xl font-bold">128</p>
                    </div>
                </div>
            </div>

            <div class="rounded-3xl bg-white p-4 text-center shadow-sm">
                <p class="text-xs text-slate-500">Palpites</p>
                <p class="mt-1 text-xl font-bold text-slate-900">
                    {{ predictionsCount }}
                </p>
            </div>

            <div class="rounded-3xl bg-white p-4 text-center shadow-sm">
                <p class="text-xs text-slate-500">Jogos</p>
                <p class="mt-1 text-xl font-bold text-slate-900">
                    {{ matchesCount }}
                </p>
            </div>

            <div class="rounded-3xl bg-white p-4 text-center shadow-sm">
                <p class="text-xs text-slate-500">Grupos</p>
                <p class="mt-1 text-xl font-bold text-slate-900">
                    {{ groupsCount }}
                </p>
            </div>
        </div>

        <div class="grid grid-cols-2 gap-3 md:grid-cols-4">
            <RouterLink to="/app/groups" class="app-card-button">
                <i class="pi pi-users text-xl text-emerald-600" />
                <p class="mt-3 font-semibold text-slate-900">Grupos</p>
            </RouterLink>

            <RouterLink to="/app/matches" class="app-card-button">
                <i class="pi pi-calendar text-xl text-emerald-600" />
                <p class="mt-3 font-semibold text-slate-900">Jogos</p>
            </RouterLink>

            <RouterLink to="/app/ranking" class="app-card-button">
                <i class="pi pi-chart-bar text-xl text-emerald-600" />
                <p class="mt-3 font-semibold text-slate-900">Ranking</p>
            </RouterLink>

            <RouterLink to="/app/profile" class="app-card-button">
                <i class="pi pi-user text-xl text-emerald-600" />
                <p class="mt-3 font-semibold text-slate-900">Perfil</p>
            </RouterLink>

            <RouterLink to="/app/rules" class="app-card-button">
                <i class="pi pi-info-circle text-xl text-emerald-600" />
                <p class="mt-3 font-semibold text-slate-900">Regras</p>
            </RouterLink>
        </div>
    </section>
</template>
<script setup>
import logo from '@/assets/logo.png';
import { computed, onMounted } from 'vue'
import { useGroupStore } from '@/stores/group'
import { useMatchStore } from '@/stores/match'
import { usePredictionStore } from '@/stores/prediction'

const groupStore = useGroupStore()
const matchStore = useMatchStore()
const predictionStore = usePredictionStore()

onMounted(async () => {
    await Promise.all([
        groupStore.fetchGroups(),
        matchStore.fetchMatches(),
    ])

    if (groupStore.groups.length > 0) {
        await predictionStore.fetchByGroup(groupStore.groups[0].id)
    }
})

const nextMatch = computed(() => {
    return matchStore.matches.find((match) => {
        return match.status === 'scheduled'
    })
})

const predictionsCount = computed(() => {
    return predictionStore.predictions.length
})

const groupsCount = computed(() => {
    return groupStore.groups.length
})
</script>