### 简介
  一个开源的bug管理系统，IT人员开发全过程
  bug管理功能基本能满足90%以上的需求， 其他需求可以自己开发，也可以提交issues    
   


### 功能
- [x] 增加bug，改变bug状态，转交bug 
- [x] 部门管理
- [x] 显示bug列表,搜索、分页
- [x] 用户创建及其操作  
- [x] 上传个人头像  
- [x] 增加邮件通知功能  
- [x] 可以修改邮箱
- [x] 增加admin用户的信息重置接口  
   admin用户有且只有一个，注册admin账户建议直接操作数据库，然后修改密码即可
   如果忘记admin的密码，可以执行下面命令重置密码，如下所示，只能在go服务器那台机器上执行
```
   curl http://127.0.0.1:10001/admin/reset?password=123
```
- [x] 增加修改邮箱，昵称，姓名页面 
- [x] bug可以指定多人，自己的bug才可以转交，删除bug内部转交功能，增加缓存,增加查看所有bug的权限  
- [x] 增加用户禁用功能，当此用户存在bug时，无法被删除  
- [x] 禁用用户，此用户的所有发布的bug都将移动至垃圾箱，垃圾箱里面的bug只有管理员才能查看，启用用户会将此用户的bug改为非垃圾箱  
- [x] 增加操作日志，只有管理员才能查看   
- [x] 状态实时保存 
- [x] 项目管理增加用户权限
- [x] 为了更好的使用，增加消息提示
- [x] 管理层可以对下级表用户表管理

### 展示页面： 
   展示页面会更新为最新可使用的代码  
   [ITflow](http://bug.hyahm.com "ITflow")  
   
### 部署  
需要 mysql >= 5.7   node 和go 最近即可，  然后还需要nginx 代理前端代码
```
git clone https://github.com/hyahm/ITflow.git
cd ITflow
```
###### 后端(安装最新版的go >= 1.16.0， 并将其目录下的bin目录添加进环境变量, 保证有go命令),  有安装好mysql数据库   
```shell
cd go
```

> 设置代理
```
export GOPROXY=https://goproxy.cn   // 国内的机器需要执行代理， 国外的机器不需要
```

> 修改配置文件   

```
go run cmd/makeconfig/config.go   # 自动生成默认配置文件到本目录   bug.ini
showbaseurl = http://127.0.0.1:10001/showimg/    #  127.0.0.1 换成外网的IP地址
salt = hjkkakoweqzmbvc   # 修改salt值后   服务启动后用 curl http://127.0.0.1:10001/admin/reset?password=123 修改admin 密码
cross=*   # 设置跨域的域名   eg:  http://127.0.0.1

# mysql 配置必须正确， 会自动创建表
[mysql]
user = root
pwd = "123456"
host = 127.0.0.1
port = 3306
db = itflow

```  
> 启动后端服务  

```
go build main.go
./main
```
###### 前端(最新版node, 保证有npm命令)
```
cd vue

优先使用cnpm 安装 
npm install  --registry=https://registry.npm.taobao.org  cnpm -g
cnpm install
或 npm 安装
npm install  --registry=https://registry.npm.taobao.org  # 安装依赖

或 yarm 安装
npm install  --registry=https://registry.npm.taobao.org  yarm -g
yarm install
```
> 修改配置文件  .env.production
```
VUE_APP_BASE_API = 'http://120.26.164.125:10001'  # 改为后端服务器外网地址
VUE_APP_USERNAME = ''  # 设置为空
VUE_APP_PASSWORD = ''   # 设置为空
```
> 打包
```
npm run build:prod

```
>  使用nginx 部署    

```
server {
        listen 80;
        server_name 127.0.0.1;
        root <ITflow_dir>/vue/dist;
        index index.html;


        location / {
                try_files $uri $uri/ @router;
                index index.html;
        }


        location @router {
                rewrite ^.*$ /index.html last;
        }
}
```

然后通过  http://<server_name>:port  访问

### 项目优势   
 部署简单, 使用简单, 高定制, 永久开源  可以自己二次开发   

  
### QQ群  
    928790087

