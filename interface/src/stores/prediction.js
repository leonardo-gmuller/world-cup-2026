import { defineStore } from 'pinia'

import {
    listPredictionsByGroup,
    savePrediction,
} from '@/services/predictionService'

export const usePredictionStore = defineStore('prediction', {
    state: () => ({
        predictions: [],
        myPredictions: [],
        loading: false,
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
    },
})