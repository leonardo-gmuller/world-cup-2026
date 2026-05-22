import { defineStore } from 'pinia'
import { listMatches, getMatchById } from '@/services/matchService'

export const useMatchStore = defineStore('match', {
    state: () => ({
        matches: [],
        loading: false,
        loadingMatchIds: {},
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

        async fetchMatch(id) {
            this.loadingMatchIds[id] = true

            try {
                const match = await getMatchById(id)

                const index = this.matches.findIndex(m => m.id === match.id)

                if (index !== -1) {
                    this.matches[index] = match
                } else {
                    this.matches.push(match)
                }
            } finally {
                delete this.loadingMatchIds[id]
            }
        }
    },
})