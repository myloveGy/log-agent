import moment from 'moment'

export const saveJson = (filename: string, data: any[]) => {
  const blob = new Blob([JSON.stringify(data)], {type: 'text/json'})
  const link = document.createElement('a')
  link.download = moment().format('YYYY-MM-DD HH:mm:ss') + filename + '.json'
  link.href = window.URL.createObjectURL(blob)
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
}
