import {onMounted, ref} from 'vue'
import {sync, userStore} from '@/utils'
import {useRouter} from 'vue-router'
import {allowRegisterApi, LoginUser} from '@/api'

export const useLogin = () => {
  const router = useRouter()
  const allowRegister = ref<string>('N')
  const action = ref<string>('login')

  const handler = async (data: LoginUser) => {
    userStore.save(data)
    return router.push('/database')
  }

  onMounted(() => {
    if (userStore.fetch('token')) {
      router.push('/database')
    }

    sync(async () => {
      const response = await allowRegisterApi()
      allowRegister.value = response.allow
    })
  })

  return {
    allow: allowRegister,
    handler,
    action,
  }
}


