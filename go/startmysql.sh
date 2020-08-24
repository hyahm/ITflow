#! /bin/bash

# docker 安装mysql
# linux 建议用host网络
docker run --restart=always --name mysql -v /data/mysql:/var/lib/mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=bug -d mysql:5.7.30