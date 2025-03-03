#!/bin/bash
set -u

# 打印所有环境变量
env

# 打印 MONGO 相关变量
echo "MONGO_INITDB_DATABASE: ${MONGO_INITDB_DATABASE:-未定义}"
echo "MONGO_USERNAME: ${MONGO_USERNAME:-未定义}"
echo "MONGO_PASSWORD: ${MONGO_PASSWORD:-未定义}"

which mongosh

mongosh <<EOF


EOF
