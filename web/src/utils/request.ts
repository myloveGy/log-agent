import {userStore} from './user'

const API_URL = import.meta.env.VITE_API_URL

export const request = async <T, >(url: string, data: any = {}, options: RequestInit = {}): Promise<T> => {
  const {headers, ...otherOptions} = options
  const init: RequestInit = {
    method: 'POST',
    body: data,
    headers: {
      'Authorization': 'Bearer ' + userStore.fetch('token'),
      'Content-Type': 'application/json',
      ...headers,
    },
    ...otherOptions,
  }

  if (init.method === 'GET') {
    init.body = null
  } else if (init.body && typeof init.body === 'object') {
    init.body = JSON.stringify(init.body)
  }

  const result = await fetch(`${API_URL}${url}`, init)
  const response = await result.json()
  if (result.status == 200) {
    return response
  }

  if (response.code === 'Unauthenticated') {
    userStore.flush()
  }

  throw response
}
