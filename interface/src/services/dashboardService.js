import api from '@/services/api'

export async function getHomeDashboard() {
  const { data } = await api.get('/dashboard')
  return data
}