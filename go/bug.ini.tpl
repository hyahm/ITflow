# 是否使用了代码，为了获取ip，可能不起作用
httpproxy = true
# 监听地址
listenaddr = :10001
# 存放图片的目录
imgdir = /data/bugimg/
# 图片显示的地址(用接口的地址)
showbaseurl = http://127.0.0.1:10001/showimg/
# 盐值，建议修改，然后用curl http://127.0.0.1:10001/admin/reset?password=123 来修改admin密码
salt = hjkkaksjdhfryuooweqzmbvc

# 默认管理员id
adminid = 1
# 默认头像地址
defaulthead = 
# 跨域字段
cross = *
# token 过期时间
expiration = 120m

# ssl, 使用ssl
[ssl]
on = false
cert = 
key = 

[log]
# 日志目录, 不设置就控制台输出
path = {{ .LOG_PATH }}
# 日志大小备份一次， 0为不切割大小
size = 0
# 每天备份一次 大小也存在的话，此项优先 ，false为不每天备份一次
everyday = false

[mysql]
user = {{ .MYSQL_USER }}
pwd = {{ .MYSQL_PASSWORD }}
host = {{ .MYSQL_HOST }}
port = {{ .MYSQL_PORT }}
db = {{ .MYSQL_DB }}
