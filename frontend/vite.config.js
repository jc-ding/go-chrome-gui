import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  build: {
    assetsDir: "static",
    /* rollupOptions: {
      output: {
        chunkFileNames:'app.js',
        entryFileNames:"root.js",//入口
        assetFileNames:"app.[ext]"//其他文件
      }
    } */
  }
})
