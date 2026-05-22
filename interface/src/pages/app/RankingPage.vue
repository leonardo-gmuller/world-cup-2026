<template>
    <section class="app-page">
        <div>
            <h1 class="app-title">Ranking</h1>
            <p class="app-subtitle">Veja quem está liderando o bolão.</p>
        </div>
        <div class="app-card p-4">
            <label class="text-xs font-semibold text-slate-500">
                Grupo selecionado
            </label>

            <Select v-model="selectedGroupId" :options="groupStore.groups" optionLabel="name" optionValue="id"
                placeholder="Selecione um grupo" class="mt-2 w-full" :loading="groupStore.loading"
                :disabled="groupStore.loading" />
        </div>
        <div v-if="rankingStore.loading" class="space-y-3">
            <article v-for="i in 5" :key="i" class="app-card flex items-center gap-3 p-4">
                <Skeleton shape="circle" size="2.5rem" />
                <div class="flex-1">
                    <Skeleton width="8rem" height="1rem" />
                    <Skeleton width="5rem" height="0.8rem" class="mt-2" />
                </div>
                <Skeleton width="4rem" height="1.2rem" />
            </article>
        </div>
        <div v-else class="space-y-3">
            <div v-if="!rankingStore.loading && ranking.length === 0"
                class="rounded-3xl bg-white p-6 text-center shadow-sm">
                <div class="mx-auto flex h-16 w-16 items-center justify-center rounded-full bg-emerald-50">
                    <i class="pi pi-chart-bar text-2xl text-emerald-600" />
                </div>

                <h3 class="mt-4 text-lg font-bold text-slate-900">
                    Ranking ainda não disponível
                </h3>

                <p class="mt-2 text-sm leading-relaxed text-slate-500">
                    O ranking será calculado assim que os primeiros jogos forem finalizados
                    e os pontos começarem a ser distribuídos.
                </p>

                <RouterLink to="/app/matches"
                    class="mt-5 inline-flex items-center justify-center rounded-2xl bg-emerald-600 px-5 py-3 text-sm font-semibold text-white shadow-sm transition-transform duration-150 active:scale-95">
                    Fazer palpites
                </RouterLink>
            </div>
            <article v-for="(item, index) in ranking" v-motion :initial="{ opacity: 0, y: 18, scale: 0.98 }" :enter="{
                opacity: 1,
                y: 0,
                scale: 1,
                transition: {
                    delay: index * 80,
                    duration: 300
                }
            }" :key="item.user_id" class="flex items-center gap-3 rounded-3xl bg-white p-4 shadow-sm">
                <div class="flex h-10 w-10 items-center justify-center rounded-full bg-emerald-50 font-bold text-emerald-600"
                    :class="{
                        'bg-yellow-100 text-yellow-700': item.position === 1,
                        'bg-slate-100 text-slate-600': item.position === 2,
                        'bg-orange-100 text-orange-700': item.position === 3,
                        'bg-emerald-50 text-emerald-600': item.position > 3
                    }">
                    {{ item.position }}
                </div>

                <div class="flex-1">
                    <h3 class="font-semibold text-slate-900">{{ item.name }}</h3>
                    <p class="text-xs text-slate-500">{{ item.predictions_count }} palpites</p>
                </div>

                <p class="text-lg font-bold text-slate-900 score-text">{{ item.total_points }} pts</p>
            </article>
        </div>
    </section>
</template>

<script setup>
import { onMounted, ref, watch, computed } from 'vue'
import Select from 'primevue/select'
import Skeleton from 'primevue/skeleton'
import { useGroupStore } from '@/stores/group'
import { useRankingStore } from '@/stores/ranking'

const groupStore = useGroupStore()
const rankingStore = useRankingStore()

const selectedGroupId = ref(null)

const ranking = computed(() => rankingStore.ranking)

onMounted(async () => {
    await groupStore.fetchGroups()

    if (groupStore.groups.length > 0) {
        selectedGroupId.value = groupStore.groups[0].id
    }
})

watch(selectedGroupId, async (groupId) => {
    if (groupId) {
        await rankingStore.fetchByGroup(groupId)
    }
})
</script>