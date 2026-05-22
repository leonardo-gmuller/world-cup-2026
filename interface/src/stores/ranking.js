import { defineStore } from 'pinia'
import { getGroupRanking } from '@/services/rankingService'

export const useRankingStore = defineStore('ranking', {
  state: () => ({
    ranking: [],
    loading: false,
  }),

  actions: {
    async fetchByGroup(groupId) {
      this.loading = true

      try {
        this.ranking = await getGroupRanking(groupId)
      } finally {
        this.loading = false
      }
    },
  },
})