import api from '@/services/api'


export async function getGroupRanking(groupId, filters = {}) {
    const { data } = await api.get(`/ranking/group/${groupId}`, {
        params: {
            ...filters
        },
    })
    return data
}