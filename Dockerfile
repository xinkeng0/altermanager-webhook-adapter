# 打包依赖阶段使用golang作为基础镜像
FROM golang:1.19.2 as build

# 启用go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

# 指定OS等，并go build
RUN GOOS=linux GOARCH=amd64 go build .

# 运行阶段指定scratch作为基础镜像
FROM alpine

WORKDIR /app

# 将上一个阶段的文件复制进来
COPY --from=build /app/main /

# 指定运行时环境变量
ENV GIN_MODE=release \
    PORT=8080

EXPOSE 8080

ENTRYPOINT ["./main"]
