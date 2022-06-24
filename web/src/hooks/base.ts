import {sync} from '@/utils'
import {ref} from 'vue'

export const useSync = <T>(fn: (data?: T) => any, defaultLoading = true) => {
  const loading = ref<boolean>(defaultLoading)
  const fetch = (data?: T) => sync(async () => {
    loading.value = true
    await fn(data)
  }).finally(() => loading.value = false)
  return {loading, fetch}
}
