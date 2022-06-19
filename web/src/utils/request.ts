const API_URL = import.meta.env.VITE_API_URL

export const request = async <T, >(url: string, data: any = {}, options: RequestInit = {}) => {
  const {headers, ...otherOptions} = options
  const init: RequestInit = {
    method: 'POST',
    body: data,
    headers: {
      'Authorization': 'VERY-JWT ',
      'Content-Type': 'application/json',
      ...headers,
    },
    ...otherOptions,
  }

  if (init.body instanceof FormData) {
    // @ts-ignore
    delete init?.headers['Content-Type']
  } else if (init.body && typeof init.body === 'object') {
    init.body = JSON.stringify(init.body)
  }

  const result = await fetch(`${API_URL}${url}`, init)
  const response = await result.json()
  if (result.status < 200 || result.status > 300) {
    // notification.error({
    //   message: `请求错误 ${result.status}: ${result.url}`,
    // })

    throw response
  }

  return response

  // if (response.code === 40003) {
  //   redirectInfo.save({redirect: window.location.href}, 3600)
  //   userStore.flush()
  //   window.location.href = '/'
  //   throw response
  // }
  //
  // if (response.code === 10000) {
  //   return response.data
  // }

  throw response
}
