<template>
    <AppSplash v-if="loading" />

    <template v-else>
        <Toast position="top-center" class="app-toast" />
        <RouterView v-slot="{ Component }">
            <Transition name="page" mode="out-in">
                <component :is="Component" />
            </Transition>
        </RouterView>
    </template>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import Toast from 'primevue/toast'
import AppSplash from '@/components/AppSplash.vue'

const loading = ref(true)

onMounted(() => {
    setTimeout(() => {
        loading.value = false
    }, 1200)
})
</script>