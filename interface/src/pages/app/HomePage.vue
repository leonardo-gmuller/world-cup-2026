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
                    <h2 class="text-xl font-bold md:text-3xl">Brasil x Alemanha</h2>
                    <p class="mt-1 text-sm opacity-90">Hoje, 16:00</p>
                </div>

                <RouterLink to="/app/matches"
                    class="app-card-button rounded-2xl bg-white/15 px-4 py-2 text-sm font-semibold">
                    Palpitar
                </RouterLink>
            </div>
        </div>

        <div class="grid grid-cols-3 gap-3">
            <div class="rounded-3xl bg-white p-4 text-center shadow-sm">
                <p class="text-xs text-slate-500">Posição</p>
                <p class="mt-1 text-xl font-bold text-slate-900">2º</p>
            </div>

            <div class="rounded-3xl bg-white p-4 text-center shadow-sm">
                <p class="text-xs text-slate-500">Pontos</p>
                <Transition name="fade-slide" mode="out-in">
                    <p :key="points" class="mt-1 text-xl font-bold text-slate-900">
                        {{ points }}
                    </p>
                </Transition>
            </div>

            <div class="rounded-3xl bg-white p-4 text-center shadow-sm">
                <p class="text-xs text-slate-500">Palpites</p>
                <p class="mt-1 text-xl font-bold text-slate-900">12</p>
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
import { ref } from 'vue';
import logo from '@/assets/logo.png';

const points = ref(120);
</script>