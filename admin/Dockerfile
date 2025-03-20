# 构建阶段
FROM node:23-alpine AS builder

WORKDIR /app

# 安装 pnpm
RUN npm install -g pnpm

# 复制包管理文件
COPY package.json pnpm-lock.yaml ./

# 配置 npm 镜像并安装依赖
RUN npm config set registry https://registry.npmmirror.com && \
    pnpm install --frozen-lockfile

# 复制所有源代码
COPY . .

# 构建应用
RUN pnpm build

EXPOSE 9002

# 启动 Nginx
CMD ["pnpm", "preview"]