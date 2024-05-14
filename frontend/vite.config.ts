import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server : {
    port : +process.env.APP_PORT! || 8000
  },
  preview : {
    port  : +process.env.APP_PORT! || 8000
  }
})
