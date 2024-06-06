#!/bin/sh
#使用docker分别构建微服务镜像
# 定义服务列表
services="user video interaction social api"
# 循环编译每个服务
for service in $services
do
  docker build --build-arg SERVICE="${service}" -t "${service}"_image .
done

# 打印完成消息
echo "All Docker images have been built successfully."