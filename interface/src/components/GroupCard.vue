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
            <div class="mt-1 flex items-center gap-2">
                <p class="font-mono text-sm font-semibold text-slate-800">
                    {{ group.invite_code }}
                </p>
                <Button class="ml-auto" icon="pi pi-copy" label="Copiar código" text rounded severity="secondary"
                    @click="copyInviteCode(group.invite_code)" size="small" />
            </div>

        </div>


    </article>
</template>
<script setup>
import Button from 'primevue/button'
import { useToast } from 'primevue/usetoast'

defineProps({
    group: {
        type: Object,
        required: true,
    },
})

const toast = useToast()

async function copyInviteCode(code) {
    try {
        if (navigator.clipboard && window.isSecureContext) {
            await navigator.clipboard.writeText(code)
        } else {
            const textarea = document.createElement('textarea')
            textarea.value = code
            textarea.style.position = 'fixed'
            textarea.style.opacity = '0'

            document.body.appendChild(textarea)
            textarea.focus()
            textarea.select()

            document.execCommand('copy')
            document.body.removeChild(textarea)
        }

        toast.add({
            severity: 'success',
            summary: 'Código copiado',
            detail: 'O código do convite foi copiado.',
            life: 3000,
        })
    } catch {
        toast.add({
            severity: 'error',
            summary: 'Erro ao copiar',
            detail: 'Não foi possível copiar o código.',
            life: 3000,
        })
    }
}
</script>