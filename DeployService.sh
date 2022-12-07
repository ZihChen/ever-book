#!/bin/bash

WORK_PATH=$(dirname $(readlink -f $0))
# 專案名稱(取當前資料夾路徑最後一個資料夾名稱)
PROJECT_NAME=${WORK_PATH##*/}
# 當前環境
ENV="local"

# 第一個參數為 LineBot Channel Secret
if [ -z "$1" ]
then
  echo "CHANNEL_SECRET is required arguments"
  exit
fi

# 第二個參數為 LineBot Channel Token
if [ -z "$2" ]
then
  echo "CHANNEL_TOKEN is required arguments"
  exit
fi

# 創建 docker network
docker network ls | grep "ever-book-service" >/dev/null 2>&1
    if  [ $? -ne 0 ]; then
        docker network create ever-book-service
    fi

# 存入.env
echo "ENV=$ENV">.env
echo "PROJECT_NAME=$PROJECT_NAME">>.env
echo "CHANNEL_SECRET=$1">>.env
echo "CHANNEL_TOKEN=$2">>.env

# 啟動容器服務
docker-compose up -d