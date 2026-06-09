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
                <div>
                    <h1 class="text-2xl font-bold text-white"> {{ title }}</h1>
                    <p class="mt-1 text-sm text-white/70">
                        {{ subtitle }}
                    </p>
                </div>
            </div>

            <Transition name="fade-slide" mode="out-in">
                <AuthLoginForm v-if="mode === 'login'" key="login" @change-mode="changeMode" />

                <AuthRegisterForm v-else-if="mode === 'register'" key="register" @change-mode="changeMode" />

                <AuthForgotPasswordForm v-else-if="mode === 'forgot'" key="forgot" />

                <AuthResetPasswordForm v-else key="reset" />
            </Transition>
        </section>
    </main>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'

import logo from '@/assets/logo.png'

import AuthLoginForm from '@/components/auth/AuthLoginForm.vue'
import AuthRegisterForm from '@/components/auth/AuthRegisterForm.vue'
import AuthForgotPasswordForm from '@/components/auth/AuthForgotPasswordForm.vue'
import AuthResetPasswordForm from '@/components/auth/AuthResetPasswordForm.vue'

const route = useRoute()
const router = useRouter()

const mode = computed(() => {
    if (route.path === '/register') return 'register'
    if (route.path === '/forgot-password') return 'forgot'
    if (route.path === '/reset-password') return 'reset'

    return 'login'
})

const title = computed(() => {
    if (mode.value === 'register') return 'Criar conta'
    if (mode.value === 'forgot') return 'Recuperar senha'
    if (mode.value === 'reset') return 'Nova senha'

    return 'Entrar no Bolão'
})

const subtitle = computed(() => {
    if (mode.value === 'register') {
        return 'Entre no bolão e comece a mandar seus palpites.'
    }

    if (mode.value === 'forgot') {
        return 'Informe seu e-mail para gerar um link de recuperação.'
    }

    if (mode.value === 'reset') {
        return 'Defina uma nova senha para acessar sua conta.'
    }

    return 'Acesse sua conta para enviar seus palpites.'
})

function changeMode() {
    router.push(mode.value === 'login' ? '/register' : '/login')
}
</script>