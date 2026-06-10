<template>
    <section class="hero-background min-h-screen bg-slate-50 text-slate-900">
        <header class="landing-header">
            <div v-motion :initial="{ opacity: 0, y: -14 }" :enter="motionEnter(0)"
                class="mx-auto flex max-w-[1400px] items-center justify-between px-8 py-5">
                <div class="flex items-center gap-3">
                    <img :src="logo" alt="Bolão Copa" class="h-12 w-12 object-contain" />

                    <div>
                        <p class="text-xs font-semibold uppercase tracking-widest text-emerald-600">
                            Copa do Mundo 2026
                        </p>

                        <h1 class="text-lg font-black">
                            Bolão Copa 2026
                        </h1>
                    </div>
                </div>

                <RouterLink to="/login"
                    class="landing-button rounded-full bg-white px-4 py-2 text-sm font-bold text-slate-800 shadow-sm">
                    Entrar
                </RouterLink>
            </div>
        </header>

        <main>
            <section class="mx-auto grid max-w-[1400px] items-center gap-12 px-8 py-10 md:grid-cols-2 md:py-24">
                <div>
                    <div v-motion :initial="{ opacity: 0, y: 22 }" :enter="motionEnter(120)"
                        class="inline-flex rounded-full bg-emerald-100 px-4 py-2 text-sm font-bold text-emerald-700">
                        O bolão da Copa para jogar com os amigos
                    </div>

                    <h2 v-motion :initial="{ opacity: 0, y: 28 }" :enter="motionEnter(220)"
                        class="mt-6 text-4xl font-black leading-tight text-slate-950 md:text-6xl">
                        Crie seu grupo, faça seus palpites e dispute o ranking.
                    </h2>

                    <p v-motion :initial="{ opacity: 0, y: 22 }" :enter="motionEnter(340)"
                        class="mt-5 max-w-xl text-base leading-7 text-slate-600 md:text-lg">
                        Participe do Bolão Copa 2026, convide seus amigos e acompanhe a pontuação durante os jogos da
                        Copa do Mundo.
                    </p>

                    <div v-motion :initial="{ opacity: 0, y: 18 }" :enter="motionEnter(460)"
                        class="mt-8 flex flex-col gap-3 sm:flex-row">
                        <RouterLink to="/register"
                            class="landing-button rounded-2xl bg-emerald-600 px-6 py-4 text-center text-sm font-black text-white">
                            Começar agora
                        </RouterLink>

                        <button @click="scrollToHowItWorks"
                            class="landing-button rounded-2xl bg-white px-6 py-4 text-center text-sm font-black text-slate-800 shadow-sm">
                            Ver como funciona
                        </button>
                    </div>
                </div>

                <div v-motion :initial="{ opacity: 0, x: 42, scale: 0.96 }" :enter="motionEnter(520)"
                    class="landing-mock relative mx-auto w-full max-w-md">
                    <article class="landing-card rounded-[2rem] bg-white p-4 shadow-2xl">
                        <div class="flex items-center justify-between gap-3">
                            <div>
                                <p class="text-xs font-medium uppercase text-emerald-600">
                                    Fase de grupos
                                </p>

                                <p class="mt-1 text-xs text-slate-500">
                                    15/06/2026, 21:00
                                </p>
                            </div>

                            <span class="rounded-full bg-slate-100 px-3 py-1 text-xs font-semibold text-slate-600">
                                Agendado
                            </span>
                        </div>

                        <div class="mt-5 grid grid-cols-[1fr_auto_1fr] items-center gap-3">
                            <div class="text-center">
                                <img src="https://flagcdn.com/w320/br.png"
                                    class="mx-auto mb-2 h-10 w-10 object-contain" />
                                <p class="text-sm font-semibold text-slate-900">Brasil</p>
                            </div>

                            <div class="rounded-2xl bg-slate-50 px-4 py-2 text-center">
                                <p class="score-text text-lg font-bold text-slate-900">x</p>
                            </div>

                            <div class="text-center">
                                <img src="https://flagcdn.com/w320/es.png"
                                    class="mx-auto mb-2 h-10 w-10 object-contain" />
                                <p class="text-sm font-semibold text-slate-900">Espanha</p>
                            </div>
                        </div>

                        <div class="mt-5 grid grid-cols-[44px_1fr_44px] items-center gap-3">
                            <div></div>

                            <div class="grid grid-cols-[1fr_auto_1fr] items-center gap-2">
                                <div
                                    class="rounded-xl border border-slate-200 bg-white px-4 py-3 text-center text-sm font-bold">
                                    2
                                </div>

                                <span class="font-bold text-slate-400">x</span>

                                <div
                                    class="rounded-xl border border-slate-200 bg-white px-4 py-3 text-center text-sm font-bold">
                                    1
                                </div>
                            </div>

                            <div
                                class="flex h-11 w-11 items-center justify-center rounded-full bg-emerald-600 text-white shadow-sm">
                                <i class="pi pi-check text-sm" />
                            </div>
                        </div>
                    </article>

                    <article
                        class="landing-card relative z-10 mx-5 -mt-4 rounded-3xl bg-slate-900 p-4 text-white shadow-2xl">
                        <div class="flex items-center justify-between">
                            <div>
                                <p class="text-xs text-white/60">Ranking do grupo</p>
                                <h3 class="mt-1 font-black">Grupo dos amigos</h3>
                            </div>

                            <span class="rounded-full bg-emerald-500/20 px-3 py-1 text-xs font-bold text-emerald-300">
                                Ao vivo
                            </span>
                        </div>

                        <div class="mt-4 space-y-3">
                            <div v-for="item in rankingMock" :key="item.position" class="flex items-center gap-3">
                                <div class="flex h-9 w-9 items-center justify-center rounded-full font-bold"
                                    :class="item.class">
                                    {{ item.position }}
                                </div>

                                <div class="flex-1">
                                    <p class="text-sm font-semibold">{{ item.name }}</p>
                                    <p class="text-xs text-white/50">18 palpites</p>
                                </div>

                                <p class="score-text text-sm font-black"
                                    :class="{ 'score-pulse': item.position === 1 }">
                                    {{ item.points }} pts
                                </p>
                            </div>
                        </div>
                    </article>
                </div>
            </section>

            <section id="como-funciona" class="mx-auto max-w-[1400px] px-8 py-12">
                <div class="grid gap-4 md:grid-cols-3">
                    <div v-for="(card, index) in howItWorks" :key="card.title" v-motion
                        :initial="{ opacity: 0, y: 24, scale: 0.98 }" :visible-once="motionEnter(index * 120)"
                        class="landing-card rounded-3xl bg-white p-6 shadow-sm">
                        <i :class="[card.icon, 'text-2xl text-emerald-600']" />

                        <h3 class="mt-4 text-xl font-black">
                            {{ card.title }}
                        </h3>

                        <p class="mt-2 text-sm leading-6 text-slate-600">
                            {{ card.text }}
                        </p>
                    </div>
                </div>
            </section>

            <section class="mx-auto max-w-[1400px] px-8 py-12">
                <div v-motion :initial="{ opacity: 0, y: 24 }" :visible-once="motionEnter(0)"
                    class="rounded-[2rem] bg-slate-900 p-6 text-white shadow-xl md:p-8">
                    <div>
                        <p class="text-sm font-bold uppercase tracking-widest text-emerald-300">
                            Durante os jogos
                        </p>

                        <h2 class="mt-3 text-3xl font-black md:text-5xl">
                            Veja o ranking mudar em tempo real.
                        </h2>

                        <p class="mt-3 max-w-2xl text-sm leading-6 text-white/60 md:text-base">
                            Quando os resultados saem, seus palpites são calculados e a disputa fica ainda mais
                            emocionante.
                        </p>
                    </div>

                    <div class="mt-8 grid gap-4 md:grid-cols-3">
                        <div class="rounded-3xl bg-white/10 p-5">
                            <p class="text-xs font-bold uppercase text-emerald-300">
                                Jogo ao vivo
                            </p>

                            <div class="mt-4 flex items-center justify-between gap-3">
                                <span class="font-bold">Brasil</span>
                                <Transition name="score-pop" mode="out-in">
                                    <span :key="liveScore"
                                        class="score-text rounded-2xl bg-white/10 px-4 py-2 text-xl font-black">
                                        {{ liveScore }}
                                    </span>
                                </Transition>
                                <span class="font-bold">Espanha</span>
                            </div>

                            <div class="mt-4 h-1.5 overflow-hidden rounded-full bg-white/10">
                                <div class="live-progress h-full rounded-full bg-emerald-400"></div>
                            </div>

                            <p class="mt-2 text-xs text-white/50">
                                Atualizando automaticamente
                            </p>
                        </div>

                        <div class="rounded-3xl bg-white/10 p-5">
                            <p class="text-xs font-bold uppercase text-yellow-300">
                                Seu palpite
                            </p>

                            <p class="mt-4 text-4xl font-black score-text">
                                2 x 1
                            </p>

                            <p class="mt-2 text-sm text-white/60">
                                Você acertou o placar exato.
                            </p>
                        </div>

                        <div class="rounded-3xl bg-emerald-500 p-5">
                            <p class="text-xs font-bold uppercase text-white/70">
                                Pontuação
                            </p>

                            <p class="mt-4 text-5xl font-black score-pulse">
                                +5 pts
                            </p>

                            <p class="mt-2 text-sm text-white/80">
                                Subiu para o top 3 do grupo.
                            </p>
                        </div>
                    </div>
                </div>
            </section>

            <section v-motion :initial="{ opacity: 0, y: 28 }" :visible-once="motionEnter(100)"
                class="mx-auto max-w-[1400px] px-8 py-20 text-center">
                <h2 class="text-3xl font-black md:text-5xl">
                    Pronto para começar o bolão?
                </h2>

                <p class="mx-auto mt-4 max-w-2xl text-slate-600">
                    Crie sua conta, monte um grupo e convide seus amigos para disputar a Copa do Mundo 2026.
                </p>

                <RouterLink to="/register"
                    class="landing-button mt-8 inline-flex rounded-2xl bg-emerald-600 px-8 py-4 text-sm font-black text-white">
                    Criar minha conta
                </RouterLink>
            </section>
        </main>

        <footer class="border-t border-slate-200 bg-white/70 backdrop-blur-sm">
            <div
                class="mx-auto flex max-w-[1400px] flex-col items-center justify-between gap-5 px-8 py-8 text-center md:flex-row md:text-left">
                <div>
                    <p class="text-sm font-semibold text-slate-900">
                        Bolão Copa 2026
                    </p>

                    <p class="mt-1 max-w-xl text-sm leading-6 text-slate-500">
                        Projeto desenvolvido para estudos, prática de desenvolvimento
                        e diversão entre amigos durante a Copa do Mundo 2026.
                    </p>

                    <p class="mt-3 text-xs text-slate-400">
                        © 2026 Bolão Copa 2026
                    </p>
                </div>

                <div class="flex flex-col items-center gap-2 md:items-end">
                    <p class="text-sm text-slate-500">
                        Desenvolvido por
                        <span class="font-semibold text-slate-700">
                            Leonardo Müller
                        </span>
                    </p>

                    <a href="https://github.com/leonardo-gmuller/world-cup-2026" target="_blank"
                        class="landing-button inline-flex items-center gap-2 rounded-xl border border-slate-200 bg-white px-4 py-2 text-sm font-semibold text-slate-700 shadow-sm">
                        <i class="pi pi-github" />
                        GitHub
                    </a>

                    <p class="text-xs text-slate-400">
                        Vue.js • Go • PostgreSQL • AWS
                    </p>
                </div>
            </div>
        </footer>
    </section>
</template>

<script setup>
import { onMounted, onBeforeUnmount, ref, computed } from 'vue'
import logo from '@/assets/logo.png'

const motionEnter = (delay = 0) => ({
    opacity: 1,
    x: 0,
    y: 0,
    scale: 1,
    transition: {
        delay,
        duration: 700,
        ease: [0.16, 1, 0.3, 1],
    },
})

const rankingMock = [
    {
        position: 1,
        name: 'Leonardo',
        points: 42,
        class: 'bg-yellow-100 text-yellow-700',
    },
    {
        position: 2,
        name: 'Júlia',
        points: 38,
        class: 'bg-slate-100 text-slate-600',
    },
    {
        position: 3,
        name: 'Thunder',
        points: 35,
        class: 'bg-orange-100 text-orange-700',
    },
]

const howItWorks = [
    {
        icon: 'pi pi-users',
        title: 'Crie seu grupo',
        text: 'Monte um grupo e convide seus amigos para disputar a Copa juntos.',
    },
    {
        icon: 'pi pi-check-circle',
        title: 'Faça seus palpites',
        text: 'Escolha os placares antes das partidas começarem.',
    },
    {
        icon: 'pi pi-chart-bar',
        title: 'Acompanhe o ranking',
        text: 'Veja quem está liderando e dispute ponto a ponto.',
    },
]

onMounted(() => {
    const app = document.getElementById('app')

    app?.classList.add('app-scroll-free')
    app?.classList.add('hide-scrollbar')

    app?.classList.remove('app-scroll-locked')

    setInterval(() => {
        scoreIndex.value =
            (scoreIndex.value + 1) % scores.length
    }, 4000)
})

onBeforeUnmount(() => {
    const app = document.getElementById('app')

    app?.classList.remove('app-scroll-free')
    app?.classList.remove('hide-scrollbar')
    app?.classList.remove('app-scroll-locked')
})

function scrollToHowItWorks() {
    const element = document.getElementById('como-funciona')

    if (!element) return

    element.scrollIntoView({
        behavior: 'smooth',
        block: 'start',
    })
}

const scores = [
    '0 x 0',
    '1 x 0',
    '1 x 1',
    '2 x 1',
]

const scoreIndex = ref(0)

const liveScore = computed(() => scores[scoreIndex.value])
</script>

<style scoped>
.hero-background {
    background:
        radial-gradient(circle at top left, rgba(16, 185, 129, .10), transparent 28%),
        radial-gradient(circle at top right, rgba(253, 224, 71, .10), transparent 24%),
        #f8fafc;
    background-size: 120% 120%;
    animation: hero-gradient 12s ease infinite;
}

@keyframes hero-gradient {

    0%,
    100% {
        background-position: 0% 50%;
    }

    50% {
        background-position: 100% 50%;
    }
}

.landing-header {
    position: sticky;
    top: 0;
    z-index: 50;
    backdrop-filter: blur(14px);
    background: rgba(248, 250, 252, 0.75);
    border-bottom: 1px solid rgba(226, 232, 240, 0.7);
    box-shadow: 0 1px 0 rgba(15, 23, 42, .04);
}

.landing-card {
    transition:
        transform 180ms ease,
        box-shadow 180ms ease,
        border-color 180ms ease,
        background-color 180ms ease;
}

.landing-card:hover {
    transform: translateY(-4px);
    box-shadow: 0 20px 45px rgba(15, 23, 42, 0.12);
}

.landing-button {
    transition:
        transform 160ms ease,
        box-shadow 160ms ease,
        background-color 160ms ease;
}

.landing-button:hover {
    transform: translateY(-2px) scale(1.01);
    box-shadow: 0 16px 30px rgba(5, 150, 105, 0.25);
}

.landing-button:active {
    transform: scale(0.97);
}

.landing-mock {
    will-change: transform;
    animation: floating 6s ease-in-out infinite;
}

@keyframes floating {

    0%,
    100% {
        transform: translateY(0px) rotate(0deg);
    }

    25% {
        transform: translateY(-4px) rotate(-0.3deg);
    }

    50% {
        transform: translateY(-10px) rotate(0deg);
    }

    75% {
        transform: translateY(-6px) rotate(0.3deg);
    }
}

.score-pulse {
    animation: score-pulse 1.8s ease-in-out infinite;
}

@keyframes score-pulse {

    0%,
    100% {
        transform: scale(1);
    }

    50% {
        transform: scale(1.08);
    }
}

.live-progress {
    width: 65%;
    animation: live-progress 12s ease-in-out infinite;
}

@keyframes live-progress {
    0% {
        width: 18%;
    }

    20% {
        width: 32%;
    }

    40% {
        width: 48%;
    }

    60% {
        width: 66%;
    }

    80% {
        width: 82%;
    }

    100% {
        width: 94%;
    }
}
.score-pop-enter-active,
.score-pop-leave-active {
  transition: all 0.5s ease;
}

.score-pop-enter-from {
  opacity: 0;
  transform: scale(1.25);
}

.score-pop-leave-to {
  opacity: 0;
  transform: scale(0.9);
}
</style>