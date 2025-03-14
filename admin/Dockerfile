FROM node:23-alpine3.21 AS builder

# 复制package.json和package-lock.json
COPY package*.json .

# 配置 npm 镜像源
RUN npm config set registry https://registry.npmmirror.com

RUN npm install -g pnpm

RUN pnpm install

# 排除node_modules
COPY . .

RUN pnpm build

EXPOSE 4173

# 运行项目
CMD ["npm", "run", "preview", "--", "--host", "0.0.0.0"]