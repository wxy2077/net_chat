FROM golang:1.22.2 AS builder

WORKDIR /app

COPY . .

RUN go mod download

# 编译应用程序
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ichat-gin-app main.go

# 使用一个更小的镜像来运行应用程序
FROM alpine:latest

# 设置时区
RUN apk --no-cache add tzdata
ENV TZ=Asia/Shanghai

# 安装证书和其他必要的依赖项
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件和配置文件
COPY --from=builder /app/ichat-gin-app .
COPY --from=builder /app/config ./config
COPY --from=builder /app/runtime ./runtime

# 暴露端口
EXPOSE 8081

# 运行应用程序
CMD ["./ichat-gin-app"]
