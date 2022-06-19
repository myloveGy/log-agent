import {ConfigEnv, loadEnv, UserConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import {resolve} from 'path'

export default ({mode}: ConfigEnv): UserConfig => {
  const root = process.cwd()
  const env = loadEnv(mode, root)
  const {VITE_PUBLIC_URL} = env
  return {
    base: VITE_PUBLIC_URL || '',
    plugins: [vue()],
    resolve: {
      alias: {
        '@': resolve(__dirname, 'src'),
      },
      extensions: ['.js', '.json', '.ts'],
    },
    server: {
      proxy: {
        '/api': {
          target: 'http://127.0.0.1:8080',
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api/, ''),
        },
      },
    },
  }
}
