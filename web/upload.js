let OSS = require('ali-oss') // 阿里云oss模块
let fs = require('fs') // 文件模块
let path = require('path')
const client = require('ftp')
const dotenv = require('dotenv')
dotenv.config({path: path.resolve('.env.local')})

const remotePath = process.argv[2]
if (process.argv.length < 3 || remotePath == null) {
  throw new Error('缺少目录参数！')
}

// 本地目录
const localPath = path.resolve('dist')
if (!fs.existsSync(localPath)) {
  throw new Error('本地目录' + localPath + '不存在！')
}

// 上传列表
const fileDic = []
// 阿里云OSS配置
const ossClient = new OSS({
  region: process.env.OSS_REGION,
  accessKeyId: process.env.OSS_ACCESS_ID,
  endpoint: process.env.OSS_ENDPOINT,
  accessKeySecret: process.env.OSS_ACCESS_SECRET,
  bucket: process.env.OSS_BUCKET,
})

console.log('---------上传OSS---------')
console.log('【Step1】 分析目录')
const readDir = (filePath) => {
  filePath = path.resolve(filePath)
  // 遍历文件目录
  let pa = fs.readdirSync(filePath)
  pa.forEach(function (filename, index) {
    let file = path.join(filePath, filename)
    let info = fs.statSync(file)
    // 目录
    if (info.isDirectory()) {
      readDir(file)
    }
    // 文件
    else {
      // 添加到上传列表
      let localDir = path.join(filePath, filename)
      let remoteDir = path.join(remotePath, localDir.replace(localPath, '')).replace(/[\\]/g, '/')
      fileDic[localDir] = remoteDir
      console.log(localDir, ' => ', remoteDir)
    }
  })
}

readDir(localPath)

console.log('【Step2】 上传文件')
const upload = async (dir) => {
  for (let localDir in dir) {
    await ossClient.put(dir[localDir], localDir)
    console.log('upload from \'' + localDir + '\' to \'' + dir[localDir] + '\'')
  }
}

upload(fileDic).then(() => console.log('【Step3】 完成')).catch(e => {
  throw new Error(e)
})

// 第一步：连接
const targetOptions = {
  host: process.env.FTP_HOST,
  port: process.env.FTP_PORT,
  user: process.env.FTP_USER,
  password: process.env.FTP_PASSWORD,
}

const ftpClient = new client()
ftpClient.connect(targetOptions)
ftpClient.on('ready', async () => {
  ftpClient.put(localPath + '/index.html', 'log-agent/index.html', () => console.info('上传文件: index.html'))
  ftpClient.end()
})
