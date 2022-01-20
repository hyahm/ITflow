### 简介
  一个开源的bug管理系统，IT人员开发全过程
  bug管理功能基本能满足90%以上的需求， 其他需求可以自己开发，也可以提交issues    



### 升级文档

为了完善开发流程， 增加tag, 同时支持后端 docker 镜像

版本更新日志 [版本更新日志](UPDATE.md)



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



### 获取源码

```
git clone https://github.com/hyahm/ITflow.git
```



### 依赖服务

- mysql >= 5.7
- nginx 或 apache(生产环境)



### 本地测试

- 需要 mysql >= 5.7   node 和go 最近即可

- 后端启动

  ```
  cd ITflow/go
  ```

  > 设置代理(国外不需要)

  ```
  export GOPROXY=https://goproxy.cn   // 国内的机器需要执行代理， 国外的机器不需要
  ```

  > 生成默认配置文件模板   

  ```
  go run .\main.exe -c bug.ini     
  ```

  >  bug.ini 修改mysql配置文件

  ```
  # mysql 配置必须正确， 会自动创建表
  [mysql]
  user = root
  pwd = "123456"
  host = 127.0.0.1
  port = 3306
  db = itflow
  
  ```

  > 启动

  ```
  go run main.go
  ```

- 前端启动

  > 安装依赖

  ```
  cd vue
  
  优先使用cnpm 安装 https://github.com/hyahm/ITflow/blob/master/README.md
  npm install  --registry=https://registry.npm.taobao.org  cnpm -g
  cnpm install
  或 npm 安装
  npm install  --registry=https://registry.npm.taobao.org  # 安装依赖
  
  或 yarm 安装
  npm install  --registry=https://registry.npm.taobao.org  yarn -g
  yarn install
  ```

  > 启动

  ```
  npm run dev
  ```

  

### 打包运行（生产环境）

- 需要 mysql >= 5.7   node 和go 最新即可

- nginx 或 apache

- 后端打包成二进制。 打包后的二进制文件只能是打包相同系统的能运行， 如果想要打包成不同系统需要设置GOOS为 linux | windows | darwin

  > 配置文件需要修改的地方

  ```
  #  127.0.0.1 换成外网的IP地址
  showbaseurl = http://127.0.0.1:10001/showimg/    
   # 修改salt值后   服务启动后用 curl http://127.0.0.1:10001/admin/reset?password=123 修改admin 密码
  salt = hjkkakoweqzmbvc  
   eg:  http://127.0.0.1
  cross=*   # 设置跨域的域名  
  
  # mysql 配置必须正确， 会自动创建表
  [mysql]
  user = root
  pwd = "123456"
  host = 127.0.0.1
  port = 3306
  db = itflow
  ```

  > 后端编译成2进制文件

  ```
  go run main.go
  ```

  >  运行

  ```
  ./main
  ```



- 前端

  > 配置文件需要修改的地方

  ```
    # 改为后端服务器外网地址
  VUE_APP_BASE_API = 'http://120.26.164.125:10001'
   # 设置为空
  VUE_APP_USERNAME = '' 
     # 设置为空
  VUE_APP_PASSWORD = ''
  ```

  > 打包成静态文件

  ```
  npm run build
  ```

  > nginx 配置文件参考

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

  

### 使用[scsctl](https://github.com/hyahm/scs 'scs')通过docker打包(生产环境)， 推荐使用

- 确保已经安装docker并启动了docker, 安装了scs

> 后端go打包

1. 修改配置文件

   ```
   # linux | windows | darwin, 设置打包后对应的系统
   GOOS: linux
   # 部分配置文件
   LOG_PATH: log/itflow.log
   MYSQL_USER: root
   MYSQL_PASSWORD: 123456
   MYSQL_HOST: 127.0.0.1
   MYSQL_PORT: 3306
   MYSQL_DB: test
   ```

   

2. 打包并将服务托管给scs管理

   ```
   scsctl install -f go/sc_build.yaml
   # 执行完成后等待生成 二进制文件
   ```

   



> vue 打包



1. 修改配置文件  sc_build.yaml

   ```
   API_DOMAIN: http://127.0.0.1:10001  # 主要是修改为后端的地址
   ```

   

2. 打包

   ```
   cd ITflow
   scsctl install -f vue/sc_build.yaml
   # 执行完成后等待生成 dist 打包的文件
   ```

3.  部署参考上面nginx配置



### 项目优势   
 部署简单, 使用简单, 高定制, 永久开源  可以自己二次开发   


### QQ群  
    928790087

