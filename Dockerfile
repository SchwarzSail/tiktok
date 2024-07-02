# 第一阶段：构建应用程序
FROM apache/skywalking-go:0.4.0-go1.22 as builder

# 定义构建参数
ARG SERVICE

# 设置环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn,direct \
    GOOS=linux \
    GOARCH=amd64 \
    SW_AGENT_REPORTER_GRPC_BACKEND_SERVICE=skywalking-oap.skywalking:11800 \
    SW_AGENT_NAME="${SERVICE}"

# 创建工作目录
RUN mkdir -p /app
WORKDIR /app

# 复制所有文件到工作目录
COPY . .

# 下载依赖
RUN go mod tidy

# 注入 SkyWalking agent
RUN skywalking-go-agent -inject ./

# 编译应用程序
RUN cd ./cmd/${SERVICE} && go build -toolexec="skywalking-go-agent" -o /app/${SERVICE}

# 第二阶段：创建最终运行环境
FROM alpine:3.17

# 安装必要的依赖
RUN apk --no-cache add ca-certificates

# 定义构建参数
ARG SERVICE

# 设置环境变量
ENV SW_AGENT_REPORTER_GRPC_BACKEND_SERVICE=skywalking-oap.skywalking:11800 \
    SW_AGENT_NAME="${SERVICE}"

# 创建工作目录
WORKDIR /app

# 从构建阶段复制应用程序二进制文件
COPY --from=builder /app/cmd/${SERVICE}/config /app/config
COPY --from=builder /app/${SERVICE} /app/${SERVICE}

# 设置 SkyWalking agent 的运行时参数
ENV SW_AGENT_COLLECTOR_BACKEND_SERVICES=skywalking-oap.skywalking:11800
ENV SW_AGENT_NAME=${SERVICE}

