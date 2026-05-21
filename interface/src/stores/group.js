import { defineStore } from 'pinia'
import { mockGroups } from '@/mocks/groups'

export const useGroupStore = defineStore('group', {
  state: () => ({
    groups: [],
    loading: false,
  }),

  actions: {
    async fetchGroups() {
      this.loading = true

      await new Promise((resolve) => setTimeout(resolve, 300))

      this.groups = mockGroups
      this.loading = false
    },

    async create(payload) {
      this.groups.unshift({
        id: Date.now(),
        internal_id: Date.now(),
        invite_code: Math.random().toString(36).slice(2, 8).toUpperCase(),
        is_active: true,
        ...payload,
      })
    },

    async join(inviteCode) {
      this.groups.unshift({
        id: Date.now(),
        internal_id: Date.now(),
        name: `Grupo ${inviteCode}`,
        description: 'Grupo adicionado por convite',
        invite_code: inviteCode,
        is_active: true,
      })
    },
  },
})