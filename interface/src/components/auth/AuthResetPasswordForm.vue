<template>
    <form class="space-y-4" @submit.prevent="submit">
        <div>
            <label class="mb-1 block text-sm font-medium text-white/80">
                Nova senha
            </label>

            <Password v-model="form.password" placeholder="Digite sua nova senha" toggleMask :feedback="false"
                inputClass="w-full" class="w-full" />
        </div>

        <div>
            <label class="mb-1 block text-sm font-medium text-white/80">
                Confirmar senha
            </label>

            <Password v-model="form.password_confirmation" placeholder="Confirme sua nova senha" toggleMask
                :feedback="false" inputClass="w-full" class="w-full" />
        </div>

        <Button type="submit" label="Redefinir senha" icon="pi pi-check" :loading="loading" class="w-full" />

        <button type="button" class="w-full text-center text-sm font-semibold text-white/70 hover:text-white"
            @click="router.push('/login')">
            Voltar para o login
        </button>
    </form>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'

import Button from 'primevue/button'
import Password from 'primevue/password'

import { useAuthStore } from '@/stores/auth'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()

const loading = ref(false)

const form = reactive({
    password: '',
    password_confirmation: '',
})

async function submit() {
    if (!route.query.token) return
    if (!form.password || !form.password_confirmation) return

    loading.value = true

    try {
        await authStore.resetPassword({
            token: route.query.token,
            password: form.password,
            password_confirmation: form.password_confirmation,
        })

        router.push('/login')
    } finally {
        loading.value = false
    }
}
</script>