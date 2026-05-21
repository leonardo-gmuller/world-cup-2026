import { defineStore } from 'pinia'
import { createGroup, joinGroup, listMyGroups } from '@/services/groupService'

export const useGroupStore = defineStore('group', {
    state: () => ({
        groups: [],
        loading: false,
    }),

    actions: {
        async fetchGroups() {
            this.loading = true
            try {
                this.groups = await listMyGroups()
            }   finally {
                this.loading = false
            }
        },

        async create(payload) {
            await createGroup(payload)
            await this.fetchGroups()
            toast.success('Grupo criado com sucesso')
        },

        async join(inviteCode) {
            await joinGroup({ invite_code: inviteCode })
            await this.fetchGroups()
            toast.success('Entrou no grupo com sucesso')
        },
    },
})