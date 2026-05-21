<template>
    <section class="app-page">
        <div>
            <h1 class="app-title">Regras do Bolão</h1>
            <p class="app-subtitle">
                Entenda como funciona a pontuação e os palpites.
            </p>
        </div>

        <div v-motion :initial="{ opacity: 0, y: 20 }" :enter="{ opacity: 1, y: 0 }" class="app-card p-5">
            <div class="flex items-center gap-3">
                <div class="rounded-2xl bg-emerald-100 p-3">
                    <i class="pi pi-clock text-emerald-700" />
                </div>

                <div>
                    <h2 class="font-bold text-slate-900">
                        Prazo para palpites
                    </h2>

                    <p class="text-sm text-slate-500">
                        Os palpites podem ser enviados até 5 minutos antes do início da partida.
                    </p>
                </div>
            </div>
        </div>

        <div v-motion :initial="{ opacity: 0, y: 20 }" :enter="{ opacity: 1, y: 0, transition: { delay: 80 } }"
            class="app-card p-5">
            <div class="flex items-center gap-3">
                <div class="rounded-2xl bg-yellow-100 p-3">
                    <i class="pi pi-star-fill text-yellow-600" />
                </div>

                <div>
                    <h2 class="font-bold text-slate-900">
                        Pontuação
                    </h2>

                    <ul class="mt-2 space-y-2 text-sm text-slate-600">
                        <li>• Acertar vencedor: <strong>3 pontos</strong></li>
                        <li>• Acertar empate: <strong>3 pontos</strong></li>
                        <li>• Acertar ao menos um placar de um time: <strong>1 ponto</strong></li>
                        <li>• Acertar placar exato: <strong>5 pontos</strong></li>
                    </ul>
                </div>
            </div>
        </div>

        <div v-motion :initial="{ opacity: 0, y: 20 }" :enter="{ opacity: 1, y: 0, transition: { delay: 160 } }"
            class="app-card p-5">
            <div class="flex items-center gap-3">
                <div class="rounded-2xl bg-emerald-100 p-3">
                    <i class="pi pi-trophy text-emerald-700" />
                </div>

                <div>
                    <h2 class="font-bold text-slate-900">
                        Peso das fases
                    </h2>

                    <ul class="mt-2 space-y-2 text-sm text-slate-600">
                        <li>• Fase de grupos: <strong>x1</strong></li>
                        <li>• Fase eliminatória: <strong>x1,5</strong></li>
                        <li>• Oitavas: <strong>x2</strong></li>
                        <li>• Quartas: <strong>x3</strong></li>
                        <li>• Semifinal: <strong>x4</strong></li>
                        <li>• Disputa de 3º lugar: <strong>x4</strong></li>
                        <li>• Final: <strong>x5</strong></li>
                    </ul>
                </div>
            </div>
        </div>

        <div v-motion :initial="{ opacity: 0, y: 20 }" :enter="{ opacity: 1, y: 0, transition: { delay: 240 } }"
            class="app-card p-5">
            <div class="flex items-center gap-3">
                <div class="rounded-2xl bg-slate-100 p-3">
                    <i class="pi pi-chart-bar text-slate-700" />
                </div>

                <div>
                    <h2 class="font-bold text-slate-900">
                        Ranking
                    </h2>

                    <p class="mt-2 text-sm text-slate-600">
                        O ranking é atualizado automaticamente após o encerramento dos jogos.
                    </p>
                </div>
            </div>
        </div>
    </section>
</template>