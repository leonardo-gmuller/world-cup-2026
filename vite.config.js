import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import tailwindcss from '@tailwindcss/vite'
import { fileURLToPath, URL } from 'node:url'

export default defineConfig({
  plugins: [vue(), tailwindcss()],

  root: 'interface',

  server: {
    port: 3000,
  },

  resolve: {
    alias: {
      '@': fileURLToPath(
        new URL('./interface/src', import.meta.url)
      ),
    },
  },
})