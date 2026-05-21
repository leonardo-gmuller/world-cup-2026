<template>
    <main class="relative flex h-screen items-center justify-center overflow-hidden bg-slate-950">
        <!-- glow background -->
        <div class="absolute inset-0 overflow-hidden">
            <div
                class="absolute left-1/2 top-1/2 h-72 w-72 -translate-x-1/2 -translate-y-1/2 rounded-full bg-emerald-500/30 blur-3xl" />

            <div class="absolute bottom-10 left-10 h-56 w-56 rounded-full bg-yellow-400/20 blur-3xl" />

            <div class="absolute right-10 top-10 h-56 w-56 rounded-full bg-emerald-300/20 blur-3xl" />
        </div>

        <!-- content -->
        <div v-motion :initial="{ opacity: 0, scale: 0.85, y: 12 }" :enter="{
            opacity: 1,
            scale: 1,
            y: 0,
            transition: {
                duration: 400
            }
        }" class="relative z-10 text-center">
            <div class="rounded-[2rem] border border-white/10 bg-white/5 p-8 backdrop-blur-xl">
                <img :src="logo" alt="Bolão Copa 2026" class="mx-auto w-44 drop-shadow-2xl" />

                <p class="mt-6 text-sm font-medium text-white/70">
                    Carregando bolão...
                </p>
            </div>
        </div>
    </main>
</template>
<script setup>
import logo from '@/assets/logo.png';
</script>