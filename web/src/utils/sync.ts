import {ElLoading, ElNotification, NotificationParamsTyped} from 'element-plus'
import {debounce} from 'lodash'

export const notifyError = debounce((options: NotificationParamsTyped) => {
  ElNotification.error(options)
}, 500)

export const sync = async (fn: any, loading: boolean = false, options = {throwError: false}) => {

  const loadingInstance = loading ? ElLoading.service({
    fullscreen: true,
    background: 'rgba(0, 0, 0, 0)',
  }) : null

  try {
    return await fn()
  } catch (e: any) {
    console.info('sync error', e)
    if (options.throwError) {
      throw e
    }

    if (!e.ignore) {
      notifyError({
        title: '温馨提醒',
        message: e instanceof Error || 'message' in e ? e.message : JSON.stringify(e),
      })
    }
  } finally {
    if (loadingInstance) {
      loadingInstance.close()
    }
  }
}
