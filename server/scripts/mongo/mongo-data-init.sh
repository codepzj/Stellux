mongosh -- "$MONGO_INITDB_DATABASE" <<EOF
db = db.getSiblingDB('$MONGO_INITDB_DATABASE')
db.auth('$MONGO_USERNAME', '$MONGO_PASSWORD');


// 用户表为用户名创建唯一升序索引
db.user.createIndex({"username":1},{"unique":true});

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
    created_at: ISODate("2025-03-02T13:44:05.737Z"),
    updated_at: ISODate("2025-03-02T13:44:05.737Z"),
    title: "mongodb学习指南",
    content: "",
    author: "codepzj",
    description: "mangguo",
    category: "技术",
    tags: [
        "golang"
    ],
    cover: "345",
    is_top: true
} );
db.getCollection("posts").update( { _id: ObjectId("67c460a5b89a82055f8d3f7e") }, { $set: { content: "# mongodb学习指南\r\n\r\n这是一个非关系型数据库\r\n\r\n然后学习过程可以使用vscode的mongo插件以及本地部署docker容器\r\n\r\n### 使用数据库\r\n\r\n```sh\r\nuse(\"admin\") # 使用admin数据库\r\nuse(\"test\") # 使用test数据库\r\n```\r\n\r\n### 权限认证\r\n\r\n```sh\r\ndb.auth(\"admin\",\"123456\")\r\n```\r\n\r\n### 获取ObjectId\r\n\r\n```sh\r\nlet AdminUuid = ObjectId();\r\nprint(AdminUuid.toString()) # 转16进制字符串\r\n```\r\n\r\n### 插入一条记录\r\n\r\n```sh\r\ndb.user.insertOne({\r\n \"_id\": AdminUuid,\r\n \"username\": \"admin\",\r\n \"password\": \"123456\",\r\n \"role_id\": 1,\r\n \"created_at\": new Date(),\r\n \"updated_at\": new Date()\r\n})\r\n```\r\n\r\n### 插入多条记录\r\n\r\n```sh\r\ndb.user.insertMany([{\r\n \"_id\": AdminId,\r\n \"username\": \"admin\",\r\n \"password\": \"123456\",\r\n \"role_id\": 0,\r\n \"created_at\": new Date(),\r\n \"updated_at\": new Date()\r\n},{\r\n \"_id\": TestId,\r\n \"username\": \"test\",\r\n \"password\": \"123456\",\r\n \"role_id\": 1,\r\n \"created_at\": new Date(),\r\n \"updated_at\": new Date()\r\n}])\r\n```\r\n\r\n### 创建索引\r\n\r\n```sh\r\ndb.user.createIndex({\"username\":1},{\"unique\":true}); # 创建升序索引和唯一索引\r\ndb.user.createIndex({\"username\":1},{\"unique\":true}); # 创建降序索引和唯一索引\r\n```\r\n\r\n" } } )

EOF
