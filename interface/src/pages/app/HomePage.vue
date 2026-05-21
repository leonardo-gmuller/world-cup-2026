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
                    <h2
                        class="flex flex-col items-center gap-3 text-lg font-bold md:flex-row md:justify-between md:text-3xl">
                        <div class="flex items-center gap-2">
                            <img v-if="nextMatch?.HomeTeamFlagURL" :src="nextMatch.HomeTeamFlagURL" loading="lazy"
                                decoding="async" class="h-6 w-6 object-contain md:h-10 md:w-10" />

                            <span class="text-center">
                                {{ nextMatch?.HomeTeamName || 'A definir' }}
                            </span>
                        </div>

                        <span class="rounded-full bg-white/15 px-3 py-1 text-sm font-black md:text-xl">
                            VS
                        </span>

                        <div class="flex items-center gap-2">
                            <span class="text-center">
                                {{ nextMatch?.AwayTeamName || 'A definir' }}
                            </span>

                            <img v-if="nextMatch?.AwayTeamFlagURL" :src="nextMatch.AwayTeamFlagURL" loading="lazy"
                                decoding="async" class="h-6 w-6 object-contain md:h-10 md:w-10" />
                        </div>
                    </h2>

                    <p class="mt-1 text-sm opacity-90">
                        {{ nextMatchDate }}
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
                <p class="text-sm opacity-90">Sua melhor posição</p>

                <div class="mt-2 flex items-end justify-between">
                    <div>
                        <h2 class="text-5xl font-black">#{{ rankingPosition }}</h2>
                        <p class="mt-1 text-sm opacity-90">
                            entre {{ totalPlayers }} participantes
                        </p>
                    </div>

                    <div class="rounded-2xl bg-white/15 px-4 py-3 backdrop-blur-sm">
                        <p class="text-xs opacity-80">Pontos</p>
                        <p class="text-2xl font-bold">{{ points }}</p>
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
import { useDashboardStore } from '@/stores/dashboard';

const dashboardStore = useDashboardStore()

onMounted(async () => {
    await dashboardStore.fetchHome()
})

const nextMatch = computed(() => dashboardStore.nextMatch)
const groupsCount = computed(() => dashboardStore.groupsCount)
const predictionsCount = computed(() => dashboardStore.predictionsCount)
const matchesCount = computed(() => dashboardStore.matchesCount)

const nextMatchDate = computed(() => {
    if (!nextMatch.value?.StartsAt) return 'Nenhum jogo encontrado'

    return new Date(nextMatch.value.StartsAt).toLocaleString('pt-BR', {
        dateStyle: 'short',
        timeStyle: 'short',
    })
})

const points = computed(() => {
    return dashboardStore.home?.total_points || 0
})

const rankingPosition = computed(() => {
    return dashboardStore.home?.user_position || '-'
})

const totalPlayers = computed(() => {
    return dashboardStore.home?.total_players || 0
})
</script>