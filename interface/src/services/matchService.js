import api from '@/services/api'

export async function listMatches() {
  const { data } = await api.get('/matches')
  return data
}

export async function listMatchesByStage(stage) {
  const { data } = await api.get(`/matches/stage/${stage}`)
  return data
}

export async function getMatchById(id) {
  const { data } = await api.get(`/matches/${id}`)
  return data
}