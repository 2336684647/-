# === 第一阶段：编译构建 ===
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY main.go .
# 【新增这一行】自动初始化 Go 模块配置文件
RUN go mod init app 
RUN CGO_ENABLED=0 GOOS=linux go build -o app-service .

# === 第二阶段：运行环境 ===
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app-service .

EXPOSE 8080
CMD ["./app-service"]