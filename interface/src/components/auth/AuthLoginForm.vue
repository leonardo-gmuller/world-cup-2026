<template>
    <Form v-slot="$form" :validateOnValueUpdate="false" :initialValues="initialValues" :resolver="resolver"
        class="space-y-4" @submit="submit">
        <div>
            <FormField name="email" v-slot="$field">
                <InputText v-bind="$field" type="email" placeholder="E-mail" fluid :invalid="$form.email?.invalid"
                    :disabled="loading" />

                <Message v-if="$field.invalid" severity="error" size="small" variant="simple">
                    {{ $field.error?.message }}
                </Message>
            </FormField>
        </div>

        <div>
            <FormField name="password" v-slot="$field">
                <Password v-bind="$field" placeholder="Senha" toggleMask fluid :feedback="false"
                    :invalid="$form.password?.invalid" :disabled="loading" @keyup.enter="handleEnter" />

                <Message v-if="$field.invalid" severity="error" size="small" variant="simple">
                    {{ $field.error?.message }}
                </Message>
            </FormField>
        </div>

        <Button type="submit" class="w-full" icon="pi pi-sign-in" :loading="loading" :disabled="loading"
            label="Entrar" />

        <button type="button"
            class="w-full text-center text-sm font-medium text-white/60 transition hover:text-white cursor-pointer"
            @click="router.push('/forgot-password')">
            Esqueci minha senha
        </button>

        <Transition name="fade-slide">
            <div v-if="errorMessage"
                class="rounded-2xl border border-red-400/20 bg-red-500/10 px-4 py-3 text-sm text-red-200">
                <div class="flex items-center gap-2">
                    <i class="pi pi-exclamation-circle" />
                    <span>{{ errorMessage }}</span>
                </div>
            </div>
        </Transition>

        <p class="text-center text-sm text-white/70">
            Ainda não tem conta?

            <button type="button" class="font-semibold text-yellow-300 cursor-pointer" @click="$emit('change-mode')">
                Criar conta
            </button>
        </p>
    </Form>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { z } from 'zod'
import { zodResolver } from '@primevue/forms/resolvers/zod'
import { Form, FormField } from '@primevue/forms'

import Button from 'primevue/button'
import InputText from 'primevue/inputtext'
import Password from 'primevue/password'
import Message from 'primevue/message'

import { useAuthStore } from '@/stores/auth'

defineEmits(['change-mode'])

const router = useRouter()
const auth = useAuthStore()

const loading = ref(false)
const errorMessage = ref('')

const initialValues = {
    email: '',
    password: '',
}

const resolver = zodResolver(
    z.object({
        email: z.string().min(1, 'Informe o e-mail.').email('Informe um e-mail válido.'),
        password: z.string().min(1, 'Informe sua senha.'),
    })
)

async function submit({ values, valid }) {
    errorMessage.value = ''

    if (!valid) return

    loading.value = true

    try {
        await auth.login(values)
        router.push('/app/home')
    } catch {
        errorMessage.value = 'E-mail ou senha inválidos.'
    } finally {
        loading.value = false
    }
}

function handleEnter(event) {
    const form = event.target.closest('form')

    if (form) {
        form.requestSubmit()
    }
}
</script>