import { defineConfig } from 'vite'
import unocss from 'unocss/vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import components from 'unplugin-vue-components/vite'
import { NaiveUiResolver } from 'unplugin-vue-components/resolvers'
import icons from 'unplugin-icons/vite'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    unocss(),
    icons({
      compiler: 'vue3',
    }),
    AutoImport({
      imports: ['vue'],
      dts: true,
    }),
    components({
      resolvers: [NaiveUiResolver()],
    }),
  ],
})
