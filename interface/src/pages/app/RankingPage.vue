<template>
    <section class="app-page">
        <div>
            <h1 class="app-title">Ranking</h1>
            <p class="app-subtitle">Veja quem está liderando o bolão.</p>
        </div>

        <div class="space-y-3">
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

                <p class="text-lg font-bold text-slate-900">{{ item.total_points }} pts</p>
            </article>
        </div>
    </section>
</template>

<script setup>
import { mockRanking } from '@/mocks/ranking'

const ranking = mockRanking
</script>