#!/bin/bash

mongosh -u "$MONGO_INITDB_ROOT_USERNAME" -p "$MONGO_INITDB_ROOT_PASSWORD" --authenticationDatabase admin <<EOF

// 创建目标数据库用户
db = db.getSiblingDB('$MONGO_INITDB_DATABASE');
db.createUser({
    user: '$MONGO_USERNAME',
    pwd: '$MONGO_PASSWORD',
    roles: [{ role: 'readWrite', db: '$MONGO_INITDB_DATABASE' }]
});

db.auth('$MONGO_USERNAME', '$MONGO_PASSWORD');

// 用户表为用户名创建唯一升序索引
db.user.createIndex({"username":1},{"unique":true});
// 文章表创建文本索引
db.posts.createIndex({
    "title": "text",
    "content": "text",
    "description": "text"
});

let AdminId = ObjectId();
let UserId = ObjectId();
let TestId = ObjectId();

// 管理员所有权限
db.casbin_rule.insertMany([{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "admin",
    "v1": "*",
    "v2": "GET",
},{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "admin",
    "v1": "*",
    "v2": "POST",
},{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "admin",
    "v1": "*",
    "v2": "PUT",
},{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "admin",
    "v1": "*",
    "v2": "DELETE",
}]);


// 普通用户部分权限
db.casbin_rule.insertMany([{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "user",
    "v1": "/user/list",
    "v2": "POST",
}]);

// 所有人均可访问（白名单）
db.casbin_rule.insertMany([{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "*",
    "v1": "/posts/list",
    "v2": "GET",
},{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "*",
    "v1": "/posts/*",
    "v2": "GET",
},{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "*",
    "v1": "/user/login",
    "v2": "POST",
},{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "*",
    "v1": "/images/*",
    "v2": "GET",
},{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "*",
    "v1": "/picture/list*",
    "v2": "GET",
}]);

// 为用户授权
db.casbin_rule.insertMany([{
    "_id": ObjectId(),
    "ptype": "g",
    "v0": AdminId,
    "v1": "admin",
},{
    "_id": ObjectId(),
    "ptype": "g",
    "v0": UserId,
    "v1": "user",
},{
    "_id": ObjectId(),
    "ptype": "g",
    "v0": TestId,
    "v1": "test",
}]);

// 初始化用户
db.user.insertMany([{
    "_id": AdminId,
    "username": "admin",
    "password": "123456",
    "role_id": 0,
    "created_at": new Date(),
    "updated_at": new Date()
},{
    "_id": UserId,
    "username": "alice",
    "password": "123456",
    "role_id": 1,
    "created_at": new Date(),
    "updated_at": new Date()
},{
    "_id": TestId,
    "username": "test",
    "password": "123456",
    "role_id": 2,
    "created_at": new Date(),
    "updated_at": new Date()
}])

// 插入文章
db.getCollection("posts").insert( {
    _id: ObjectId("67c453eda04b00c407b431fd"),
    created_at: new Date(),
    updated_at: new Date(),
    title: "stellux博客",
    content: "当你看见这段文字的时候，说明你已经成功了",
    author: "codepzj",
    description: "mangguo",
    category: "技术",
    tags: [
        "golang"
    ],
    cover: "https://avatars.githubusercontent.com/u/183695872?v=4",
    is_top: true
});

EOF
