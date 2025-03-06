#!/bin/bash

# 默认用户名和密码
username=${1:-admin}
password=${2:-123456}

# 连接 MongoDB，等待服务可用
echo "等待 MongoDB 服务启动..."
until mongosh --quiet --eval "db.runCommand({ ping: 1 })" &>/dev/null; do
  echo "MongoDB 未就绪，5 秒后重试..."
  sleep 5
done

echo "MongoDB 已启动，开始检查副本集状态..."

# 获取副本集状态
status=$(mongosh "mongodb://$username:$password@localhost:27017/admin" --quiet --eval "var s=rs.status().ok; if (s == 1) print(1);")

# 如果 status 为空，则认为未初始化
if [ -z "$status" ]; then
  echo "副本集未初始化，正在初始化..."
  
  # 初始化副本集配置
  mongosh "mongodb://$username:$password@localhost:27017/admin" --eval "
    print('初始化副本集配置');
    var config = {
      _id: 'rs0',
      members: [
        { _id: 0, host: 'localhost:27017' }
      ]
    };
    print('初始化副本集');
    rs.initiate(config);
    print('副本集初始化完成');
  "

  # 等待副本集进入 PRIMARY 状态
  echo "等待副本集进入 PRIMARY 状态..."
  until mongosh "mongodb://$username:$password@localhost:27017/admin" --quiet --eval "rs.status().myState" | grep -q "1"; do
    echo "副本集未就绪，5 秒后重试..."
    sleep 5
  done

  echo "副本集初始化完成，并已进入 PRIMARY 状态！"
else
  echo "副本集已初始化，跳过初始化步骤。"
fi
