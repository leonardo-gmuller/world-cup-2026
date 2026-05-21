import { defineStore } from 'pinia'
import { mockMatches } from '@/mocks/matches'

export const useMatchStore = defineStore('match', {
  state: () => ({
    matches: [],
    loading: false,
  }),

  actions: {
    async fetchMatches() {
      this.loading = true

      await new Promise((resolve) => setTimeout(resolve, 300))

      this.matches = mockMatches
      this.loading = false
    },

    async importAll() {
      await this.fetchMatches()
    },
  },
})