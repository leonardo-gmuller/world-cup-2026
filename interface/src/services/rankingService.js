import api from '@/services/api'


export async function getGroupRanking(groupId) {
  const { data } = await api.get(`/ranking/group/${groupId}`)
  return data
}