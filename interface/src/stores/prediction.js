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
                const savedPrediction = await savePrediction(payload)
                const index = this.predictions.findIndex((item) => {
                    return item.match_id === savedPrediction.match_id
                })

                if (index >= 0) {
                    this.predictions[index] = savedPrediction
                } else {
                    this.predictions.push(savedPrediction)
                }

                return savedPrediction
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