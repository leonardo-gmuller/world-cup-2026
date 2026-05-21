import api from '@/services/api'

export async function listMyGroups(userId) {
  const { data } = await api.get(`/groups/user/${userId}`)
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