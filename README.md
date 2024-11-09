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
   [ITflow](http://itflow.hyahm.com "ITflow")  



### 获取源码

```
git clone https://github.com/hyahm/ITflow.git
```



### 依赖服务

- mysql >= 5.7
- nginx 或 apache 或 caddy(推荐)


### 本地测试

- 需要 mysql >= 5.7   node 和go 最近即可

- 后端启动（注意路径）要在go路径下

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

  > 安装依赖（注意路径，在ITflow/vue下）

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

> 后端接口也可以通过caddy转为https, 参考如下

```
itflow.api.hyahm.com:443 {
    reverse_proxy 127.0.0.1:10001
}
```



### 线上部署


- 后端

  >  直接在ITflow/go下执行， 会生成 可执行文件

  ```
  go build main.go
  ```

  可以直接通过进程管理工具或systemd或者docker启动

  

- 前端

  > 配置文件需要修改的地方 .env.production(需要自己新建一个文件)

  ```
    # 改为后端服务器外网地址
  VUE_APP_BASE_API = 'http://127.0.0.1:10001'
   # 设置为空
  VUE_APP_USERNAME = '' 
     # 设置为空
  VUE_APP_PASSWORD = ''
  ```

  > 打包成静态文件

  ```
  npm run build
  ```

  > 会在当前目录下生成dist文件夹， 这个文件夹就是前端的根目录文件

  > caddy 的配置文件参考

  ```
  itflow.hyahm.com:443 {
      root * <ITflow_dir>/vue/dist
      try_files {path} {path}/ /index.html
      file_server
  }
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




### docker 部署

我也很想用docker compose部署， 但是前端打包前配置文件需要修改，不知道怎么弄

### 项目优势   

 部署简单, 使用简单, 高定制, 永久开源  可以自己二次开发   


### QQ群  
    928790087

