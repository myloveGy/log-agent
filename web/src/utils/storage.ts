export class LocalStore<T> {
  name: string

  constructor(name: string) {
    this.name = name
  }

  flush(key: any = null) {
    const returnValue = key === null ? null : this.fetch(key)
    localStorage.removeItem(this.name)
    return returnValue
  }

  fetch(key: string = '', defaultValue: any = null): T | any {
    const data = localStorage.getItem(this.name)
    if (!data) {
      return defaultValue
    }

    if (data[0] !== '{') {
      this.flush()
    }

    const values = JSON.parse(data)

    // 设置了有效时间的数据
    if (values.hasOwnProperty('value') && values.hasOwnProperty('expire_time')) {
      // 有超时时间
      if (values.expire_time !== 0 && (new Date()).getTime() > values.expire_time) {
        this.flush()
        return defaultValue
      }

      return key ? (values.value?.[key] || defaultValue) : values.value
    }

    return key ? (values?.[key] || defaultValue) : values
  }

  /**
   * 保存信息
   * @param data 数据信息
   * @param expire_time
   */
  save(data: any, expire_time: number = 0) {
    let expireTime = expire_time

    // 有效时间
    if (expireTime > 0) {
      expireTime = (new Date()).getTime() + expire_time * 1000
    }

    localStorage.setItem(this.name, JSON.stringify({
      value: {
        ...this.fetch(),
        ...data,
      },
      expire_time: expireTime,
    }))
  }
}
