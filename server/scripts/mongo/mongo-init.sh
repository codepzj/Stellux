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
db.user.createIndex({"username": 1}, {"unique": true});

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
    "v2": "GET"
},{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "admin",
    "v1": "*",
    "v2": "POST"
},{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "admin",
    "v1": "*",
    "v2": "PATCH"
},{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "admin",
    "v1": "*",
    "v2": "PUT"
},{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "admin",
    "v1": "*",
    "v2": "DELETE"
}]);

// 普通用户权限
db.casbin_rule.insertMany([{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "user",
    "v1": "/user/list",
    "v2": "POST"
}]);

// 测试用户权限
db.casbin_rule.insertMany([{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "*",
    "v1": "/posts/list",
    "v2": "GET"
}, {
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "*",
    "v1": "/posts/*",
    "v2": "GET"
}, {
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "*",
    "v1": "/user/login",
    "v2": "POST"
}, {
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "*",
    "v1": "/images/*",
    "v2": "GET"
}, {
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "*",
    "v1": "/picture/list*",
    "v2": "GET"
}]);

// 为用户授权
db.casbin_rule.insertMany([{
    "_id": ObjectId(),
    "ptype": "g",
    "v0": AdminId,
    "v1": "admin"
}, {
    "_id": ObjectId(),
    "ptype": "g",
    "v0": UserId,
    "v1": "user"
}, {
    "_id": ObjectId(),
    "ptype": "g",
    "v0": TestId,
    "v1": "test"
}]);

// 初始化用户
db.user.insertMany([{
    "_id": AdminId,
    "username": "admin",
    "password": "123456",
    "role_id": 0,
    "created_at": new Date(),
    "updated_at": new Date()
}, {
    "_id": UserId,
    "username": "alice",
    "password": "123456",
    "role_id": 1,
    "created_at": new Date(),
    "updated_at": new Date()
}, {
    "_id": TestId,
    "username": "test",
    "password": "123456",
    "role_id": 2,
    "created_at": new Date(),
    "updated_at": new Date()
}]);

// 插入文章
db.posts.insertMany([{
    _id: ObjectId("67c453eda04b00c407b431fd"),
    created_at: new Date(),
    updated_at: new Date(),
    title: "stellux博客",
    content: "## 主要内容\r\n> #### Markdown*是什么*？\r\n> #### *谁*创造了它？\r\n> #### *为什么*要使用它？\r\n> #### *怎么*使用？\r\n> #### *谁*在用？\r\n> #### 尝试一下\r\n\r\n## 正文\r\n### 1. Markdown*是什么*？\r\n**Markdown**是一种轻量级**标记语言**，它以纯文本形式(*易读、易写、易更改*)编写文档，并最终以HTML格式发布。    \r\n**Markdown**也可以理解为将以MARKDOWN语法编写的语言转换成HTML内容的工具。    \r\n\r\n### 2. *谁*创造了它？\r\n它由[**Aaron Swartz**](http://www.aaronsw.com/)和**John Gruber**共同设计，**Aaron Swartz**就是那位于去年（*2013年1月11日*）自杀,有着**开挂**一般人生经历的程序员。维基百科对他的[介绍](http://zh.wikipedia.org/wiki/%E4%BA%9A%E4%BC%A6%C2%B7%E6%96%AF%E6%B2%83%E8%8C%A8)是：**软件工程师、作家、政治组织者、互联网活动家、维基百科人**。    \r\n\r\n他有着足以让你跪拜的人生经历：    \r\n+ **14岁**参与RSS 1.0规格标准的制订。     \r\n+ **2004**年入读**斯坦福**，之后退学。   \r\n+ **2005**年创建[Infogami](http://infogami.org/)，之后与[Reddit](http://www.reddit.com/)合并成为其合伙人。   \r\n+ **2010**年创立求进会（Demand Progress），积极参与禁止网络盗版法案（SOPA）活动，最终该提案被撤回。   \r\n+ **2011**年7月19日，因被控从MIT和JSTOR下载480万篇学术论文并以免费形式上传于网络被捕。     \r\n+ **2013**年1月自杀身亡。    \r\n\r\n![Aaron Swartz](https://github.com/younghz/Markdown/raw/master/resource/Aaron_Swartz.jpg)\r\n\r\n天才都有早逝的归途。\r\n\r\n### 3. *为什么*要使用它？\r\n+ 它是易读（看起来舒服）、易写（语法简单）、易更改**纯文本**。处处体现着**极简主义**的影子。\r\n+ 兼容HTML，可以转换为HTML格式发布。\r\n+ 跨平台使用。\r\n+ 越来越多的网站支持Markdown。\r\n+ 更方便清晰地组织你的电子邮件。（Markdown-here, Airmail）\r\n+ 摆脱Word（我不是认真的）。\r\n\r\n### 4. *怎么*使用？\r\n如果不算**扩展**，Markdown的语法绝对**简单**到让你爱不释手。\r\n\r\nMarkdown语法主要分为如下几大部分：\r\n**标题**，**段落**，**区块引用**，**代码区块**，**强调**，**列表**，**分割线**，**链接**，**图片**，**反斜杠 `\\`**，**符号'`'**。\r\n\r\n#### 4.1 标题\r\n两种形式：  \r\n1）使用`=`和`-`标记一级和二级标题。\r\n> 一级标题   \r\n> `=========`   \r\n> 二级标题    \r\n> `---------`\r\n\r\n效果：\r\n> 一级标题   \r\n> =========   \r\n> 二级标题\r\n> ---------  \r\n\r\n2）使用`#`，可表示1-6级标题。\r\n> \\# 一级标题   \r\n> \\## 二级标题   \r\n> \\### 三级标题   \r\n> \\#### 四级标题   \r\n> \\##### 五级标题   \r\n> \\###### 六级标题    \r\n\r\n效果：\r\n> # 一级标题   \r\n> ## 二级标题   \r\n> ### 三级标题   \r\n> #### 四级标题   \r\n> ##### 五级标题   \r\n> ###### 六级标题\r\n\r\n#### 4.2 段落\r\n段落的前后要有空行，所谓的空行是指没有文字内容。若想在段内强制换行的方式是使用**两个以上**空格加上回车（引用中换行省略回车）。\r\n\r\n#### 4.3 区块引用\r\n在段落的每行或者只在第一行使用符号`>`,还可使用多个嵌套引用，如：\r\n> \\> 区块引用  \r\n> \\>> 嵌套引用  \r\n\r\n效果：\r\n> 区块引用  \r\n>> 嵌套引用\r\n\r\n#### 4.4 代码区块\r\n代码区块的建立是在每行加上4个空格或者一个制表符（如同写代码一样）。如    \r\n普通段落：\r\n\r\nvoid main()    \r\n{    \r\n    printf(\"Hello, Markdown.\");    \r\n}    \r\n\r\n代码区块：\r\n\r\n    void main()\r\n    {\r\n        printf(\"Hello, Markdown.\");\r\n    }\r\n\r\n**注意**:需要和普通段落之间存在空行。\r\n\r\n#### 4.5 强调\r\n在强调内容两侧分别加上`*`或者`_`，如：\r\n> \\*斜体\\*，\\_斜体\\_    \r\n> \\*\\*粗体\\*\\*，\\_\\_粗体\\_\\_\r\n\r\n效果：\r\n> *斜体*，_斜体_    \r\n> **粗体**，__粗体__\r\n\r\n#### 4.6 列表\r\n使用`·`、`+`、或`-`标记无序列表，如：\r\n> \\-（+\\*） 第一项\r\n> \\-（+\\*） 第二项\r\n> \\- （+\\*）第三项\r\n\r\n**注意**：标记后面最少有一个_空格_或_制表符_。若不在引用区块中，必须和前方段落之间存在空行。\r\n\r\n效果：\r\n> + 第一项\r\n> + 第二项\r\n> + 第三项\r\n\r\n有序列表的标记方式是将上述的符号换成数字,并辅以`.`，如：\r\n> 1 . 第一项   \r\n> 2 . 第二项    \r\n> 3 . 第三项    \r\n\r\n效果：\r\n> 1. 第一项\r\n> 2. 第二项\r\n> 3. 第三项\r\n\r\n#### 4.7 分割线\r\n分割线最常使用就是三个或以上`*`，还可以使用`-`和`_`。\r\n\r\n#### 4.8 链接\r\n链接可以由两种形式生成：**行内式**和**参考式**。    \r\n**行内式**：\r\n> \\[younghz的Markdown库\\]\\(https:://github.com/younghz/Markdown \"Markdown\"\\)。\r\n\r\n效果：\r\n> [younghz的Markdown库](https:://github.com/younghz/Markdown \"Markdown\")。\r\n\r\n**参考式**：\r\n> \\[younghz的Markdown库1\\]\\[1\\]    \r\n> \\[younghz的Markdown库2\\]\\[2\\]    \r\n> \\[1\\]:https:://github.com/younghz/Markdown \"Markdown\"    \r\n> \\[2\\]:https:://github.com/younghz/Markdown \"Markdown\"    \r\n\r\n效果：\r\n> [younghz的Markdown库1][1]    \r\n> [younghz的Markdown库2][2]\r\n\r\n[1]: https:://github.com/younghz/Markdown \"Markdown\"\r\n[2]: https:://github.com/younghz/Markdown \"Markdown\"\r\n\r\n**注意**：上述的`[1]:https:://github.com/younghz/Markdown \"Markdown\"`不出现在区块中。\r\n\r\n#### 4.9 图片\r\n添加图片的形式和链接相似，只需在链接的基础上前方加一个`！`。\r\n#### 4.10 反斜杠`\\`\r\n相当于**反转义**作用。使符号成为普通符号。\r\n#### 4.11 符号'`'\r\n起到标记作用。如：\r\n>\\`ctrl+a\\`\r\n\r\n效果：\r\n>`ctrl+a`    \r\n\r\n#### 5. *谁*在用？\r\nMarkdown的使用者：\r\n+ GitHub\r\n+ 简书\r\n+ Stack Overflow\r\n+ Apollo\r\n+ Moodle\r\n+ Reddit\r\n+ 等等\r\n\r\n#### 6. 尝试一下\r\n+ **Chrome**下的插件诸如`stackedit`与`markdown-here`等非常方便，也不用担心平台受限。\r\n+ **在线**的dillinger.io评价也不错   \r\n+ **Windowns**下的MarkdownPad也用过，不过免费版的体验不是很好。    \r\n+ **Mac**下的Mou是国人贡献的，口碑很好。\r\n+ **Linux**下的ReText不错。    \r\n\r\n**当然，最终境界永远都是笔下是语法，心中格式化 :)。**\r\n\r\n****\r\n**注意**：不同的Markdown解释器或工具对相应语法（扩展语法）的解释效果不尽相同，具体可参见工具的使用说明。\r\n虽然有人想出面搞一个所谓的标准化的Markdown，[没想到还惹怒了健在的创始人John Gruber]\r\n(http://blog.codinghorror.com/standard-markdown-is-now-common-markdown/ )。\r\n****\r\n以上基本是所有traditonal markdown的语法。\r\n\r\n### 其它：\r\n列表的使用(非traditonal markdown)：\r\n\r\n用`|`表示表格纵向边界，表头和表内容用`-`隔开，并可用`:`进行对齐设置，两边都有`:`则表示居中，若不加`:`则默认左对齐。\r\n\r\n|代码库                              |链接                                |\r\n|:------------------------------------:|------------------------------------|\r\n|MarkDown                              |[https://github.com/younghz/Markdown](https://github.com/younghz/Markdown \"Markdown\")|\r\n|MarkDownCopy                              |[https://github.com/younghz/Markdown](https://github.com/younghz/Markdown \"Markdown\")|\r\n\r\n\r\n关于其它扩展语法可参见具体工具的使用说明。",
    author: "codepzj",
    description: "stellux博客",
    category: "技术",
    tags: [
        "golang"
    ],
    cover: "https://avatars.githubusercontent.com/u/183695872?v=4",
    is_top: true,
    is_publish: true
}]);

EOF