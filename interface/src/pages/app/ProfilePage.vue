<template>
  <section class="app-page">
    <div>
      <h1 class="app-title">Perfil</h1>
      <p class="app-subtitle">Gerencie sua conta.</p>
    </div>

    <div class="rounded-3xl bg-white p-5 shadow-sm">
      <div class="flex items-center gap-3">
        <div class="flex h-14 w-14 items-center justify-center rounded-full bg-emerald-100 text-xl font-bold text-emerald-700">
          {{ initials }}
        </div>

        <div>
          <h2 class="font-bold text-slate-900">{{ user?.name || 'Usuário' }}</h2>
          <p class="text-sm text-slate-500">{{ user?.email || 'email@email.com' }}</p>
        </div>
      </div>

      <Button
        label="Sair"
        severity="danger"
        outlined
        class="mt-5 w-full"
        @click="logout"
      />
    </div>
  </section>
</template>

<script setup>
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import Button from 'primevue/button'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const auth = useAuthStore()

const user = computed(() => auth.user)

const initials = computed(() => {
  return user.value?.name?.slice(0, 1)?.toUpperCase() || 'U'
})

function logout() {
  auth.logout()
  router.push('/login')
}
</script>