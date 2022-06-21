import {request} from '@/utils'
import {AllowRegister, DatabaseQuery, LoginUser, PageQuery, Pagination, User} from './types'

export * from './types'

export type UserRequest = Pick<User, 'username' | 'password'>

// 登录
export const loginApi = (data: UserRequest) => request<LoginUser>('/login', data)
export const registerApi = (data: UserRequest) => request<LoginUser>('/register', data)
export const allowRegisterApi = () => request<AllowRegister>('/allow-register', {}, {method: 'GET'})

// 用户相关
type UserResponse = Pick<User, 'username'>

export type UserListRequest = Partial<User> & Partial<PageQuery> & { query?: any }
export const userCreateApi = (data: UserRequest) => request<UserResponse>('/user/create', data)
export const userUpdateApi = (data: Partial<User>) => request<UserResponse>('/user/update', data)
export const userListApi = (data: UserListRequest) => request<Pagination<User>>('/user/list', data)

// 查询
export const databaseQueryApi = (data: DatabaseQuery) => request<Pagination<any>>('/database/query', data)
export const databaseCollectionsApi = () => request<{ items: string[] }>('/database/collections')
