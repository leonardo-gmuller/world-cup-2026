import { defineStore } from 'pinia'
import { getHomeDashboard } from '@/services/dashboardService'

export const useDashboardStore = defineStore('dashboard', {
  state: () => ({
    home: null,
    loading: false,
  }),

  getters: {
    nextMatch: (state) => state.home?.next_match || null,

    groupsCount: (state) => state.home?.groups_count || 0,

    predictionsCount: (state) => state.home?.predictions_count || 0,

    matchesCount: (state) => state.home?.matches_count || 0,
  },

  actions: {
    async fetchHome() {
      this.loading = true

      try {
        this.home = await getHomeDashboard()
      } finally {
        this.loading = false
      }
    },
  },
})