import api from '@/services/api'

export async function listMyGroups() {
  const { data } = await api.get(`/groups`)
  return data
}

export async function createGroup(payload) {
  const { data } = await api.post('/groups', payload)
  return data
}

export async function joinGroup(payload) {
  const { data } = await api.post('/groups/join', payload)
  return data
}