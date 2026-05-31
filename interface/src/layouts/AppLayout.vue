<template>
    <div class="h-dvh overflow-hidden bg-gradient-to-b from-slate-50 to-slate-100">
        <main ref="contentRef" class="mx-auto h-full w-full max-w-md overflow-y-auto px-4 py-4 md:max-w-2xl lg:max-w-4xl scroll-container"
            style="padding-bottom: calc(5.5rem + env(safe-area-inset-bottom));">
            <RouterView v-slot="{ Component }">
                <Transition name="page" mode="out-in">
                    <component :is="Component" />
                </Transition>
            </RouterView>
        </main>

        <AppBottomNav />
    </div>
</template>

<script setup>
import AppBottomNav from '@/components/AppBottomNav.vue'
import { ref, watch, nextTick } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const contentRef = ref(null)

watch(
    () => route.fullPath,
    async () => {
        await nextTick()

        contentRef.value?.scrollTo({
            top: 0,
            behavior: 'smooth',
        })
    }
)
</script>