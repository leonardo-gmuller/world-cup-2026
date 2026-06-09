<template>
    <form class="space-y-4" @submit.prevent="submit">
        <div>
            <label class="mb-1 block text-sm font-medium text-white/80">
                E-mail
            </label>

            <InputText v-model="form.email" type="email" placeholder="Digite seu e-mail" class="w-full" />
        </div>

        <Button type="submit" label="Gerar link de recuperação" icon="pi pi-link" :loading="loading" class="w-full" />

        <div v-if="success"
            class="rounded-2xl border border-emerald-400/30 bg-emerald-400/10 p-3 text-sm text-emerald-100">
            {{ message }}
        </div>

        <button type="button" class="w-full text-center text-sm font-semibold text-white/70 hover:text-white"
            @click="router.push('/login')">
            Voltar para o login
        </button>
    </form>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'

import Button from 'primevue/button'
import InputText from 'primevue/inputtext'

import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const loading = ref(false)
const success = ref(false)
const message = ref('')

const form = reactive({
    email: '',
})

async function submit() {
    if (!form.email) return

    loading.value = true

    try {
        const response = await authStore.forgotPassword({
            email: form.email,
        })

        success.value = true
        message.value = response.message || 'Se o e-mail existir, um link de recuperação será gerado.'
    } finally {
        loading.value = false
    }
}
</script>