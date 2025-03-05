#!/bin/bash

# 默认用户名和密码
username=${1:-admin}
password=${2:-123456}

# 检查副本集状态
status=$(mongosh "mongodb://$username:$password@localhost:27017/admin" --eval "rs.status().ok" --quiet)

# 如果 rs.status().ok 返回 1，说明副本集已初始化
if [ "$status" -eq 1 ]; then
  echo "副本集已初始化，跳过初始化步骤。"
else
  echo "副本集未初始化，正在初始化..."

  # 初始化副本集配置和状态检查
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

    print('检查副本集状态');
    printjson(rs.status());
  "
fi
