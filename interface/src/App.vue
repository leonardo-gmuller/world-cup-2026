<template>
    <AppSplash v-if="showSplash && loading" />

    <template v-else>
        <Toast position="top-center" class="app-toast" />
        <RouterView v-slot="{ Component, route }">
            <Transition v-if="!route.path.startsWith('/app')" name="page">
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

watch(
    () => route.path,
    (to, from) => {
        const enteringApp = to.startsWith('/app') && !from?.startsWith('/app')

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