# Stellux 知识库

Stellux - 于留白中沉淀，于极简中思索

后端仓库: https://github.com/StelluxWiki/Stellux-Server

前端仓库: https://github.com/StelluxWiki/Stellux-Web

后台仓库: https://github.com/StelluxWiki/Stellux-Admin

项目预览：[https://gowiki.site](https://gowiki.site/)

项目文档：[https://gowiki.site/doc/stellux/683c0b8e96f61777ec3e29ad](https://gowiki.site/doc/stellux/683c0b8e96f61777ec3e29ad)

## 项目部署文档

* 后端部署：[后端部署指南](https://gowiki.site/doc/stellux/683c0b8e96f61777ec3e29ad/683c0c9796f61777ec3e29b0)
* 后台部署：[后台部署指南](https://gowiki.site/doc/stellux/683c0b8e96f61777ec3e29ad/683c134096f61777ec3e29b1)
* 前端部署：[前端部署指南](https://gowiki.site/doc/stellux/683c0b8e96f61777ec3e29ad/683c135c96f61777ec3e29b2)

## 项目简介

**Stellux** 是一款极简风格的博客兼知识库系统，采用纯黑白色调设计，专注写作与阅读体验。它摒弃传统博客的友链、评论、说说等繁杂功能，融合了类似 VitePress 的文档功能，适合对写作环境有极致追求的用户。

## 技术栈

* 前端：React、Next.js、TailwindCSS、TypeScript、Zustand、React Hooks、HeroUI、ShadcnUI、Vercel、Docker
* 后端：Gin、MongoDB、Go-Mongox、Viper、Casbin、Slog、Wire、Go-Template
* 控制台：Vue、TailwindCSS、Ant Design Vue、Pinia

## 项目初衷

市面上的博客和 CMS 系统种类繁多，但没有一款让我真正满意。界面或花哨或臃肿，写作体验不够纯粹。我喜欢 Typora 那种极简且沉浸的写作风格，因此想打造一个在线版 Typora —— 知识点笔记同步呈现于网页，阅读简洁纯粹，同时具备动态博客应有的灵活性和可定制性。

虽然 VitePress、Hexo 等静态博客系统优秀且主题丰富，但自定义能力有限；Typecho、Halo 等动态博客插件丰富，但臃肿，不适合专注写作；思源笔记、语雀、飞书功能强大，但学习成本高且重量级。Stellux 保留静态博客写作舒适度，同时兼具适度动态功能和现代特性，使写作更简单高效。

## 项目演示

### 前端界面预览

首页
![首页](https://image.codepzj.cn/image/20250601151953328.png)

博客详情
![博客详情](https://image.codepzj.cn/image/20250601152027843.png)

文档详情
![文档详情](https://image.codepzj.cn/image/20250601152034301.png)

### 控制台界面预览

首页
![首页](https://image.codepzj.cn/image/20250601152110138.png)

用户管理
![用户管理](https://image.codepzj.cn/image/20250601152145354.png)

文档管理
![文档功能](https://image.codepzj.cn/image/20250601152224586.png)

图片墙
![图片墙](https://image.codepzj.cn/image/20250601152311570.png)

设置页面
![设置页面](https://image.codepzj.cn/image/20250601152356534.png)

## 项目运行指南

### 前端和控制台

环境需求：Node.js 18 及以上

```bash
pnpm i
pnpm run dev
```

默认后台账号密码：

* 管理员：admin / 123456
* 普通用户：alice / 123456
* 游客：test / 123456

在控制端根目录创建 `.env.production`，内容：

```env
VITE_API_URL=http://localhost:9002/api
```

在前端根目录创建 `.env.production`，内容：

```env
NEXT_PUBLIC_PROJECT_API=http://localhost:9001
NEXT_PUBLIC_SITE_URL=http://localhost:9003
```

### 后端

环境需求：Go 1.22+ 和 MongoDB 8+

```bash
go mod tidy
```

配置文件 `stellux.development.yaml` 示例：

```yaml
mongodb:
  HOST: "localhost"
  PORT: 27017
  MONGO_INITDB_ROOT_USERNAME: "admin"
  MONGO_INITDB_ROOT_PASSWORD: "123456"
  MONGO_INITDB_DATABASE: "stellux"
  MONGO_USERNAME: "stellux"
  MONGO_PASSWORD: "123456"

server:
  PORT: 9001
  JWT_SECRET: "stellux"
```

### 使用 Makefile 启动

```makefile
start:
	docker compose -f docker-compose.development.yaml up -d

rebuild:
	docker compose -f docker-compose.development.yaml down
	rm -rf data/mongo
	mkdir -p data/mongo
	docker compose -f docker-compose.development.yaml up -d

stop:
	docker compose -f docker-compose.development.yaml down
```

在 WSL 环境中执行 `make` 启动开发环境的 Docker 容器，然后在根目录执行：

```bash
go run main.go
```

成功启动截图：

![运行成功](https://image.codepzj.cn/image/20250601153403462.png)

## 项目结构

### 控制端 (admin)

```
admin
├── public                  # 静态资源，如图片、字体等
├── src                     # 源代码目录
│   ├── assets              # 图片、样式等资源
│   ├── components          # 通用 UI 组件
│   ├── views               # 页面文件夹
│   ├── router              # 路由配置
│   ├── stores              # 状态管理
│   └── utils               # 工具函数
├── index.html              # 入口 HTML
├── package.json            # 依赖配置
└── vite.config.ts          # Vite 配置
```

### 后端 (server)

```
server
├── internal
│   ├── document
│   │   ├── internal
│   │   │   ├── domain        # 领域模型
│   │   │   ├── repository    # 数据访问层
│   │   │   ├── service       # 业务逻辑层
│   │   │   └── web           # API 相关
│   │   ├── module.go         # 模块定义
│   │   └── wire.go           # 依赖注入配置
│   ├── file                  # 文件模块
│   ├── label                 # 标签模块
│   ├── post                  # 文章模块
│   ├── setting               # 配置模块
│   ├── user                  # 用户模块
│   ├── ioc
│   │   ├── gin.go            # Gin 框架初始化
│   │   ├── middleware.go     # 中间件
│   │   └── mongodb.go        # MongoDB 连接配置
│   └── pkg
│       ├── apiwrap           # API 包装工具
│       ├── middleware        # 中间件工具包
│       └── utils             # 通用工具
├── main.go                   # 程序入口
├── go.mod                    # Go 模块
└── go.sum                    # 依赖校验
```

### 前端 (web)

```
web
├── api                      # 前端接口请求封装
├── app                      # 页面及布局
│   ├── (blog)               # 博客相关
│   ├── (document)           # 文档相关
│   ├── error.tsx            # 错误页
│   ├── layout.tsx           # 全局布局
│   ├── loading.tsx          # 加载动画
│   ├── not-found.tsx        # 404 页面
│   ├── providers.tsx        # 上下文及状态提供者
│   ├── robots.ts            # robots.txt 配置
│   └── sitemap.ts           # 网站地图生成
├── assets                   # 静态资源
│   ├── png
│   └── svg
├── components               # 通用组件
│   ├── Card
│   ├── Empty
│   ├── Footer
│   ├── Loading
│   ├── Md                   # Markdown 相关组件
│   ├── SearchModal
│   ├── Sidebar
│   ├── SvgIcon
│   ├── ThemeSwitcher
│   ├── Toc
│   ├── Tool
│   └── ui                   # 基础 UI 组件库
├── public                   # 公共资源目录
├── styles                   # 样式文件
│   ├── mdx.css
│   └── tailwind.css
├── tailwind.config.js       # TailwindCSS 配置
├── tsconfig.json            # TypeScript 配置
├── next.config.js           # Next.js 配置
└── package.json             # 依赖配置

```

## 未来计划

- 增加更丰富的写作辅助功能（如全文检索、标签联动）
- 优化移动端体验
- 增强权限体系及团队协作功能

感谢使用 Stellux，欢迎贡献代码和建议，一起打造极致写作体验！  
项目地址：[https://github.com/StelluxWiki/Stellux](https://github.com/StelluxWiki/Stellux)  
开发者：浩瀚星河  
