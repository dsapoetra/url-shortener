import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';


export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      // The backend URL is http://localhost:8080/api/urls
      '/api/urls': 'http://localhost:8080', // Replace with your backend URL
    },
  },
});
