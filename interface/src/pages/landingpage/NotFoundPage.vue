<template>
    <section class="not-found-page flex min-h-screen items-center justify-center px-6 text-slate-900">
        <div class="mx-auto w-full max-w-xl text-center">
            <RouterLink
                to="/"
                class="mx-auto mb-8 flex w-fit items-center gap-3"
            >
                <img :src="logo" alt="Bolão Copa" class="h-12 w-12 object-contain" />

                <div class="text-left">
                    <p class="text-xs font-semibold uppercase tracking-widest text-emerald-600">
                        Copa do Mundo 2026
                    </p>

                    <p class="text-lg font-black">
                        Bolão Copa 2026
                    </p>
                </div>
            </RouterLink>

            <div class="not-found-card rounded-[2rem] bg-white p-6 shadow-2xl">
                <p class="score-text text-7xl font-black text-emerald-600">
                    404
                </p>

                <h1 class="mt-4 text-3xl font-black">
                    Essa página foi eliminada
                </h1>

                <p class="mx-auto mt-3 max-w-md text-sm leading-6 text-slate-500">
                    O link que você tentou acessar não existe ou saiu de campo.
                </p>

                <div class="mt-6 rounded-3xl bg-slate-50 p-4">
                    <p class="text-xs font-semibold uppercase text-slate-500">
                        Placar final
                    </p>

                    <div class="mt-3 flex items-center justify-center gap-4">
                        <div class="text-center">
                            <p class="text-sm font-bold text-slate-900">
                                Página
                            </p>
                        </div>

                        <div class="rounded-2xl bg-white px-5 py-3 shadow-sm">
                            <p class="score-text text-2xl font-black text-slate-900">
                                0 x 4
                            </p>
                        </div>

                        <div class="text-center">
                            <p class="text-sm font-bold text-slate-900">
                                Você
                            </p>
                        </div>
                    </div>
                </div>

                <div class="mt-7 flex flex-col gap-3 sm:flex-row sm:justify-center">
                    <RouterLink
                        to="/"
                        class="rounded-2xl bg-emerald-600 px-6 py-4 text-sm font-black text-white shadow-sm transition hover:bg-emerald-700 active:scale-[0.98]"
                    >
                        Voltar para a página inicial
                    </RouterLink>

                    <RouterLink
                        to="/app/home"
                        class="rounded-2xl bg-slate-100 px-6 py-4 text-sm font-black text-slate-800 transition hover:bg-slate-200 active:scale-[0.98]"
                    >
                        Entrar no bolão
                    </RouterLink>
                </div>
            </div>
        </div>
    </section>
</template>

<script setup>
import logo from '@/assets/logo.png'
</script>

<style scoped>
.not-found-page {
    background:
        radial-gradient(circle at top left, rgba(16, 185, 129, .10), transparent 28%),
        radial-gradient(circle at top right, rgba(253, 224, 71, .10), transparent 24%),
        #f8fafc;
}

.not-found-card {
    animation: page-in 700ms cubic-bezier(.16, 1, .3, 1);
}

@keyframes page-in {
    from {
        opacity: 0;
        transform: translateY(22px) scale(.98);
    }

    to {
        opacity: 1;
        transform: translateY(0) scale(1);
    }
}
</style>