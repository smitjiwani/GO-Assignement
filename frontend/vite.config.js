import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  server: {
    host: true, // Add this to allow network access
    port: 5173,
    proxy: {
      '/api': {
        target: 'http://server:8080' || 'http://localhost:8080',
        secure: true,
        changeOrigin: true,
      },
    },
    watch: {
      usePolling: true,
    },
    hmr: {
      host: 'localhost',
    },
  },
})
