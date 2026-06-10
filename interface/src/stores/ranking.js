import { defineStore } from 'pinia'
import { getGroupRanking } from '@/services/rankingService'

export const useRankingStore = defineStore('ranking', {
    state: () => ({
        ranking: [],
        loading: false,
    }),

    actions: {
        async fetchByGroup(groupId, filters = {}) {
            this.loading = true

            try {
                this.ranking = await getGroupRanking(groupId, filters)
            } finally {
                this.loading = false
            }
        },
    },
})