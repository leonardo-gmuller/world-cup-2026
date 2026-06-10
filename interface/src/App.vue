<template>
    <AppSplash v-if="showSplash && loading" />

    <template v-else>
        <Toast position="top-center" class="app-toast" />
        <RouterView v-slot="{ Component, route }">
            <Transition v-if="usePageTransition" name="page">
                <component :is="Component" :key="route.fullPath" />
            </Transition>

            <component v-else :is="Component" class="scroll-container" />
        </RouterView>
    </template>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { useRoute } from 'vue-router'
import Toast from 'primevue/toast'
import AppSplash from '@/components/AppSplash.vue'

const route = useRoute()

const loading = ref(false)

const showSplash = computed(() => {
    return route.path.startsWith('/app')
})

const noTransitionPrefixes = [
    '/app',
    '/login',
    '/register',
    '/forgot-password',
    '/reset-password',
]

const usePageTransition = computed(() => {
    return !noTransitionPrefixes.some((prefix) => {
        return route.path.startsWith(prefix)
    })
})

watch(
    () => route.path,
    (to, from) => {
        const cameFromOutsideApp = !from || !from.startsWith('/app')
        const enteringApp = to.startsWith('/app') && cameFromOutsideApp

        if (!enteringApp) {
            loading.value = false
            return
        }

        loading.value = true

        setTimeout(() => {
            loading.value = false
        }, 1200)
    },
    { immediate: true }
)
</script>