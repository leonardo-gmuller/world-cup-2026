import { defineStore } from 'pinia'

import {
    listPredictionsByGroup,
    savePrediction,
    getPredictionReminders,
} from '@/services/predictionService'

export const usePredictionStore = defineStore('prediction', {
    state: () => ({
        predictions: [],
        myPredictions: [],
        reminders: [],
        loading: false,
        remindersLoading: false,
    }),

    actions: {
        async fetchByGroup(groupId) {
            this.loading = true

            try {
                this.predictions = await listPredictionsByGroup(groupId)
            } finally {
                this.loading = false
            }
        },

        async fetchMyPredictions(groupId) {
            this.loading = true

            try {
                this.myPredictions = await listPredictionsByGroup(groupId)
            } finally {
                this.loading = false
            }
        },

        async save(payload) {
            this.loading = true

            try {
                await savePrediction(payload)
            } finally {
                this.loading = false
            }
        },

        async getReminders() {
            this.remindersLoading = true

            try {
                this.reminders = await getPredictionReminders()
            } finally {
                this.remindersLoading = false
            }
        },
    },
})