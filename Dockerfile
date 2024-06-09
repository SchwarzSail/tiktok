# 使用 skywalking-go 基础镜像
FROM apache/skywalking-go:0.4.0-go1.22
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
RUN cd ./cmd/${SERVICE} && go build -toolexec="skywalking-go-agent" -o ./${SERVICE}

