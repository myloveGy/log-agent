import {LocalStore} from './storage'
import {LoginUser} from '@/api'

export const userStore = new LocalStore<LoginUser>('user')

