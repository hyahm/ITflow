// const baseUrl = 'http://127.0.0.1:10001'
const baseUrl = 'https://api.hyahm.coms:10001'

const g = {
// cookie 过期时间，单位分，与后端保持一致
  expirament: 120,
  downloadUrl: baseUrl + '/share/down', // 下载用到的地址
  username: 'admin@qq.com',
  password: 'admin',
  uploadUrl: baseUrl + '/uploadimg',
  headImgUrl: baseUrl + '/upload/headimg',
  shareUpload: baseUrl + '/share/upload', // 共享文件夹上传接口
  apiUrl: baseUrl,
  pubkey: `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCorvPHEzrJUXiqd+qZwEFC7E9d
VIyrXq4Cchwc3r301KWQuIJ4l/m4lMFRXIh18yxItRMKrDlp1pNCzPKy6LBTuOwu
CDwodJSM3UyAm3ezi9vzKK4ci7LOm3Gv7uUGCzW7IEyi3OpN+QTalLlaXHO0w2e6
aE4A1HF1gL/y8BXF/wIDAQAB
-----END PUBLIC KEY-----`
}

export default g
