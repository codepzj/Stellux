# 拉取golang:1.24-alpine3.21轻量镜像
FROM golang:1.24-alpine3.21 AS builder

# 配置模块代理
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

## 在docker的根目录下创建相应的使用目录
RUN mkdir -p /app

## 设置工作目录
WORKDIR /app

## 将后端的gin代码复制到docker容器中
COPY . .

RUN go mod download
RUN go mod tidy

## go构建可执行文件
RUN go build main.go

RUN chmod +x main

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app .

EXPOSE 9001

CMD ["./main"]