# 第一阶段：构建阶段
FROM golang:1.21 as builder

# 设置环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn,direct \
    GOOS=linux \
    GOARCH=amd64

# 创建工作目录
RUN mkdir -p /app
WORKDIR /app

# 定义构建参数
ARG SERVICE

# 复制所有文件
COPY . .

# 安装依赖并编译
RUN go mod tidy
RUN mkdir -p output || exit 1
RUN cd "cmd/${SERVICE}" && go build -o "../../output/${SERVICE}"

# 第二阶段：运行阶段
FROM alpine:latest


# 设置工作目录
WORKDIR /app

# 复制依赖文件和可执行文件
COPY --from=builder /app/config /app/config
COPY --from=builder /app/output /app/output



