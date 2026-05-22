<template>
    <main class="relative flex h-screen items-center justify-center overflow-hidden bg-slate-950 px-4">
        <div class="absolute inset-0 overflow-hidden">
            <div class="absolute left-1/2 top-1/3 h-80 w-80 -translate-x-1/2 rounded-full bg-emerald-500/30 blur-3xl" />
            <div class="absolute bottom-10 left-8 h-56 w-56 rounded-full bg-yellow-400/20 blur-3xl" />
            <div class="absolute right-8 top-10 h-56 w-56 rounded-full bg-emerald-300/20 blur-3xl" />
        </div>
        <section
            class="relative z-10 w-full max-w-md rounded-[2rem] border border-white/10 bg-white/10 p-6 shadow-2xl backdrop-blur-xl">
            <div class="mb-6 text-center">
                <img :src="logo" class="mx-auto mb-4 w-36 drop-shadow-2xl" />
                <div v-if="isLogin">
                    <h1 class="text-2xl font-bold text-white">Entrar no Bolão</h1>
                    <p class="mt-1 text-sm text-white/70">
                        Acesse sua conta para enviar seus palpites.
                    </p>
                </div>
                <div v-else>
                    <h1 class="text-2xl font-bold text-white">Criar conta</h1>
                    <p class="mt-1 text-sm text-white/70">
                        Entre no bolão e comece a mandar seus palpites.
                    </p>
                </div>
            </div>

            <Transition name="fade-slide" mode="out-in">
                <AuthLoginForm v-if="isLogin" key="login" @change-mode="changeMode" />
                <AuthRegisterForm v-else key="register" @change-mode="changeMode" />
            </Transition>
        </section>
    </main>
</template>

<script setup>
import { ref } from 'vue'
import logo from '@/assets/logo.png'
import AuthLoginForm from '@/components/auth/AuthLoginForm.vue'
import AuthRegisterForm from '@/components/auth/AuthRegisterForm.vue'

import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()

const isLogin = computed(() => route.path === '/login')

function changeMode() {
    router.push(isLogin.value ? '/register' : '/login')
}

</script>