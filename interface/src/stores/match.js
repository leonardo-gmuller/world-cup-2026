import { defineStore } from 'pinia'
import { importMatches, listMatches } from '@/services/matchService'

export const useMatchStore = defineStore('match', {
  state: () => ({
    matches: [],
    loading: false,
  }),

  actions: {
    async fetchMatches() {
      this.loading = true
      try {
        this.matches = await listMatches()
      } finally {
        this.loading = false
      }
    },

    async importAll() {
      this.loading = true
      try {
        await importMatches()
        await this.fetchMatches()
      } finally {
        this.loading = false
      }
    },
  },
})