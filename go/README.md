### bug.ini配置文件说明
```ini

暂时只有下面的环境变量定义
MYSQL_USER
MYSQL_PASSWORD
MYSQL_HOST
MYSQL_PORT
MYSQL_DB
ADMINID
SHOW_URL
ACAO

- 安装之前请确保安装了go>=1.16并配置了环境变量
- 确保安装了mysql >= 5.7

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
# 共享文件夹根目录(已经废弃)
sharedir = /share/
# 默认管理员id
adminid = 1
# 默认头像地址
defaulthead = 
# token 过期时间
expiration = 120m
debug=true

# ssl, 使用ssl
[ssl]
on = false
cert = 
key = 

[log]
# 日志目录, 不设置就控制台输出
path = log/bug.log
# 日志大小备份一次， 0为不切割大小
size = 0
# 每天备份一次 大小也存在的话，此项优先 ，false为不每天备份一次
everyday = false

[mysql]
user = test
pwd = "123456"
host = 192.168.50.250
port = 3306
db = itflow

[scs]
domain=http://127.0.0.1:11111
path = D:\myproject\ITflow\go

```
