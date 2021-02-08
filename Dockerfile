FROM golang:alpine AS builder

# 为镜像设置环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录: /build
WORKDIR /build

# 创建一个app-runner用户，-D表示无密码
# RUN adduser -u 10001 -D app-runner

# 复制项目中的go.mod和go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将代码编译成二进制可执行文件app
RUN go build -o app .

#创建一个小镜象
FROM scratch

#这里可以复制项目里的静态文件和配置文件
# COPY ./static /static
COPY ./conf /conf

# 从builder镜像中把/build/app拷贝到当前目录
COPY --from=builder /build/app /

# 声明服务端口
EXPOSE 8090

#USER app-runner
#RUN chmod -R 755 /app

# 启动容器时运行命令
ENTRYPOINT [ "/app", "conf/config.yaml"]