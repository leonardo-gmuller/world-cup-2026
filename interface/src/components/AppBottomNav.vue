<template>
  <nav class="fixed bottom-0 left-0 right-0 z-50 border-t border-slate-200 bg-white">
    <div class="mx-auto grid max-w-md grid-cols-5 px-2 py-2 md:max-w-2xl lg:max-w-4xl">
      <RouterLink
        v-for="item in items"
        :key="item.to"
        :to="item.to"
        class="flex flex-col items-center justify-center gap-1 rounded-2xl px-2 py-2 text-xs text-slate-500"
        :class="{ 'bg-emerald-50 text-emerald-600 font-semibold': route.path === item.to }"
      >
        <i :class="item.icon" class="text-lg" />
        <span>{{ item.label }}</span>
      </RouterLink>
    </div>
  </nav>
</template>

<script setup>
import { useRoute } from 'vue-router'

const route = useRoute()

const items = [
  { label: 'Início', icon: 'pi pi-home', to: '/app/home' },
  { label: 'Grupos', icon: 'pi pi-users', to: '/app/groups' },
  { label: 'Jogos', icon: 'pi pi-calendar', to: '/app/matches' },
  { label: 'Ranking', icon: 'pi pi-chart-bar', to: '/app/ranking' },
  { label: 'Perfil', icon: 'pi pi-user', to: '/app/profile' },
]
</script>