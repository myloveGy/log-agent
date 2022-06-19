import {request} from '@/utils'
import {DatabaseQuery, PageQuery, Pagination, User} from './types'

export * from './types'

export type UserRequest = Pick<User, 'username' | 'password'>

export interface LoginUser extends User {
  token: string
}

// 登录
export const userLoginApi = (data: UserRequest) => request<LoginUser>('/login', data)

// 用户相关
type UserResponse = Pick<User, 'username'>

export type UserListRequest = Partial<User> & Partial<PageQuery> & { query?: any }
export const userCreateApi = (data: UserRequest) => request<UserResponse>('/user/create', data)
export const userUpdateApi = (data: Partial<User>) => request<UserResponse>('/user/update', data)
export const userListApi = (data: UserListRequest) => request<Pagination<User>>('/user/list', data)

// 查询
export const databaseQueryApi = (data: DatabaseQuery) => request<Pagination<any>>('/database/query', data)
