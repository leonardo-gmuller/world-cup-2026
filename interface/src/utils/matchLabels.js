export function stageLabel(stage) {
  const labels = {
    group_stage: 'Fase de grupos',
    last_32: '16 avos',
    round_of_16: 'Oitavas',
    round_of_32: 'Fase eliminatória',
    last_16: 'Oitavas',
    quarter_final: 'Quartas',
    quarter_finals: 'Quartas',
    semi_final: 'Semifinal',
    semi_finals: 'Semifinal',
    third_place: '3º lugar',
    final: 'Final',
  }

  return labels[String(stage || '').toLowerCase()] || stage
}

export function statusLabel(status) {
  const labels = {
    scheduled: 'Agendado',
    timed: 'Agendado',
    live: 'Ao vivo',
    in_play: 'Ao vivo',
    paused: 'Intervalo',
    finished: 'Encerrado',
    postponed: 'Adiado',
    cancelled: 'Cancelado',
  }

  return labels[String(status || '').toLowerCase()] || status
}