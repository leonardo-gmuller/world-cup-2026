<template>
    <div class="h-screen overflow-hidden bg-slate-50">
        <main class="mx-auto h-full w-full max-w-md overflow-y-auto px-4 py-4 pb-24 md:max-w-2xl lg:max-w-4xl">
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
</script>