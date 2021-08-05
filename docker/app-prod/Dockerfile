FROM golang:1.16.6-alpine3.14 as builder

# 增加允许访问内网镜像地址
# RUN GOPRIVATE=bitbucket.chartertechs.info
# 修改镜像加速和添加压缩工具
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache upx ca-certificates tzdata git openssh

# 指定服务名称
ARG SERVICE
# 指定casbin 模型配置文件名称
ARG CASBIN_MODEL_CONF_PATH=casbin_model.conf
# 指定内网镜像用户名
# ARG BITBUCKETUSERNAME
# 指定内网镜像密码
# ARG BITBUCKETPASSWORD
RUN mkdir /internal
COPY ./internal /internal
# COPY ./db/migrations/admin /internal/admin/migrations
# # 拷贝内网镜像地址访问所需秘钥
# COPY ./id_rsa  /root/.ssh/id_rsa
# # 修改下载内网镜像用户名和密码映射
# RUN git config --global url."https://${BITBUCKETUSERNAME}:${BITBUCKETPASSWORD}@bitbucket.chartertechs.info/".insteadOf "https://bitbucket.chartertechs.info/"
# 指定工作目录
WORKDIR /internal/$SERVICE
# 构建APP二进制执行文件
RUN go build -o /app .

# 构建轻量级应用,保留二进制执行文件,所需数据库迁移文件,其他文件
FROM alpine as runner
WORKDIR /app

RUN apk add curl

# 配置时区和ssl
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# 拷贝二进制文件
COPY --from=builder /app .
# 拷贝casbin模型配置文件
COPY --from=builder /internal/$SERVICE/$CASBIN_MODEL_CONF_PATH .
# # 拷贝migrations文件夹
# COPY --from=builder /internal/admin/migrations ./migrations
# # 拷贝admin权限文件
# COPY --from=builder /internal/admin/output.admin_casbin.json .

CMD ["./app"]