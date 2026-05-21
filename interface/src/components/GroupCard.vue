<template>
    <article v-motion :initial="{ opacity: 0, y: 20 }" :enter="{ opacity: 1, y: 0 }"
        class="rounded-3xl bg-white p-4 shadow-sm">
        <div class="flex items-start justify-between gap-3">
            <div>
                <h3 class="font-bold text-slate-900">{{ group.name }}</h3>
                <p v-if="group.description" class="mt-1 text-sm text-slate-500">
                    {{ group.description }}
                </p>
            </div>

            <span class="rounded-full bg-emerald-50 px-3 py-1 text-xs font-semibold text-emerald-600">
                Ativo
            </span>
        </div>

        <div class="mt-4 rounded-2xl bg-slate-50 p-3">
            <p class="text-xs text-slate-500">Código de convite</p>
            <p class="font-mono text-sm font-semibold text-slate-800">
                {{ group.invite_code }}
            </p>
        </div>
    </article>
</template>

<script setup>
defineProps({
    group: {
        type: Object,
        required: true,
    },
})
</script>