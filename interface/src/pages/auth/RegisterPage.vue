<template>
    <main class="relative flex h-screen items-center justify-center overflow-hidden bg-slate-950 px-4">
        <div class="absolute inset-0 overflow-hidden">
            <div class="absolute left-1/2 top-1/3 h-80 w-80 -translate-x-1/2 rounded-full bg-emerald-500/30 blur-3xl" />
            <div class="absolute bottom-10 left-8 h-56 w-56 rounded-full bg-yellow-400/20 blur-3xl" />
            <div class="absolute right-8 top-10 h-56 w-56 rounded-full bg-emerald-300/20 blur-3xl" />
        </div>

        <section v-motion :initial="{ opacity: 0, y: 24, scale: 0.96 }"
            :enter="{ opacity: 1, y: 0, scale: 1, transition: { duration: 350 } }"
            class="relative z-10 w-full max-w-md rounded-[2rem] border border-white/10 bg-white/10 p-6 shadow-2xl backdrop-blur-xl">
            <div class="mb-6 text-center">
                <img :src="logo" alt="Bolão Copa 2026" class="mx-auto mb-4 w-36 drop-shadow-2xl" />

                <h1 class="text-2xl font-bold text-white">Criar conta</h1>
                <p class="mt-1 text-sm text-white/70">
                    Entre no bolão e comece a mandar seus palpites.
                </p>
            </div>

            <form class="space-y-4" @submit.prevent="submit">
                <InputText v-model="form.name" placeholder="Nome" class="w-full" />

                <InputText v-model="form.email" type="email" placeholder="E-mail" class="w-full" />

                <Password v-model="form.password" placeholder="Senha" toggleMask fluid :feedback="false" />

                <Button type="submit" label="Criar conta" icon="pi pi-user-plus" class="w-full" :loading="loading" />
            </form>

            <p class="mt-5 text-center text-sm text-white/70">
                Já tem conta?
                <RouterLink to="/login" class="font-semibold text-yellow-300">
                    Entrar
                </RouterLink>
            </p>
        </section>
    </main>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import Button from 'primevue/button'
import InputText from 'primevue/inputtext'
import Password from 'primevue/password'
import { useAuthStore } from '@/stores/auth'
import logo from '@/assets/logo.png';
const router = useRouter()
const auth = useAuthStore()

const loading = ref(false)

const form = reactive({
    name: '',
    email: '',
    password: '',
})

async function submit() {
    loading.value = true

    try {
        await auth.register(form)
        router.push('/app/home')
    } finally {
        loading.value = false
    }
}
</script>