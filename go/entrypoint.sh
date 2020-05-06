#! /bin/bash
cat > bug.conf << EOF
[http]
# 是否使用了代码，为了获取ip，可能不起作用
httpproxy=true

[bug]
# 监听地址
listenaddr=${PORT:-:10001}
# 存放图片的目录
imgdir=/data/bugimg
# 图片显示的地址
showbaseurl=${IMG_SERVER_URI:-http://127.0.0.1:10001}/showimg
# 私钥
privatekeyfile=pri.key
# 盐值，建议修改，然后用curl http://ip:10001/admin/reset?password=123 来修改root密码
salt=hjkkaksjdhfryuooweqzmbvc
# token 过期时间
redisexpiration=120
# 共享文件夹根目录
sharedir=/share
# 排除记录这些ip日志
exclude=[]
# api文档
apihelp=false
# 项目名
apiname=itflow
# ssl, 使用ssl
ssl = false
certfile=
keyfile=

[redis]
redispwd =${REDIS_PWD:-__@picker-redis}
redishost=${REDIS_HOST:-172.17.0.1:6379}
redisdb=${REDIS_DB:-0}

[mysql]
mysqluser=${MYSQL_USER:-cander}
mysqlpwd=${MYSQL_PWD:-rF7oJBKopxiLJYyW2XJ&h&tyHX95D4}
mysqlhost=${MYSQL_HOST:-hyahm.com}
mysqlport=${MYSQL_PORT:-3306}
mysqldb=${MYSQL_DB:-project}
sqldriver=${DRIVER:-mysql}

[log]
# 日志目录
logpath=/data/log
# 日志大小备份一次， 0为不切割大小
logsize=0
# 每天备份一次 大小也存在的话，此项优先 ，false为不每天备份一次
logeveryday=false
EOF

exec "$@"
