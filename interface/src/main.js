import { createApp } from 'vue'
import { createPinia } from 'pinia'

import PrimeVue from 'primevue/config'
import Aura from '@primevue/themes/aura'

import ToastService from 'primevue/toastservice'
import { MotionPlugin } from '@vueuse/motion'

import 'primeicons/primeicons.css'
import './assets/main.css'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.use(PrimeVue, {
    theme: {
        preset: Aura,
        options: {
            darkModeSelector: false,
        },
    },
})

app.use(ToastService)
app.use(MotionPlugin)

app.mount('#app')

// Remove pinch zoom on mobile devices
document.addEventListener(
    'gesturestart',
    function (e) {
        e.preventDefault()
    },
    { passive: false }
)

document.addEventListener(
    'gesturechange',
    function (e) {
        e.preventDefault()
    },
    { passive: false }
)

document.addEventListener(
    'gestureend',
    function (e) {
        e.preventDefault()
    },
    { passive: false }
)

let lastTouchEnd = 0

document.addEventListener(
    'touchend',
    function (event) {
        const now = Date.now()

        if (now - lastTouchEnd <= 300) {
            event.preventDefault()
        }

        lastTouchEnd = now
    },
    false
)