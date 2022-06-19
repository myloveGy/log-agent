import {request} from '../utils/request'

// 登录
export const loginApi = (data: any) => request<any>('/login', data)

// 用户相关
export const userUpdateApi = (data: any) => request<any>('/user/update', data)

// 查询
export const databaseQueryApi = (data: any) => request<any>('/database/query', data)
