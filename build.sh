#!/bin/sh
#使用docker分别构建微服务镜像
# 定义服务列表
services="user video interaction social api"
# 循环编译每个服务
for service in $services
do
  make ${service}
done
