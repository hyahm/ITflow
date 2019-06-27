#! /bin/bash
cd  ${HOMEDIR:-/home/vue}
if ["$PUBKEY" = ""]; then
	PUBKEY=`cat pub.key`
fi	
cat > src/config/config.js << EOF
const baseUrl = '${BASE_URL:-http://127.0.0.1:10001}'

const g = {
// cookie 过期时间，单位分，与后端保持一致
  expirament: ${EXPIRETIME:-120},
  downloadUrl: baseUrl + '/share/down', // 下载用到的地址
  username: 'admin@qq.com',
  password: 'admin',
  uploadUrl: baseUrl + '/uploadimg',
  headImgUrl: baseUrl + '/upload/headimg',
  shareUpload: baseUrl + '/share/upload', // 共享文件夹上传接口
  apiUrl: baseUrl,
  pubkey: \`${PUBKEY}\`
}

export default g
EOF
exec "$@"
