import api from '@/services/api'

export async function savePrediction(payload) {
  const { data } = await api.post('/predictions', payload)
  return data
}

export async function listPredictionsByGroup(groupId) {
  const { data } = await api.get(`/predictions/group/${groupId}`)
  return data
}

export async function getPredictionReminders() {
    const { data } = await api.get('/predictions/reminders')

    return data
}