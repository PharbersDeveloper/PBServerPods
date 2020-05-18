# builder 源镜像
FROM alpine:latest as builder

# 安装系统级依赖
RUN echo http://mirrors.aliyun.com/alpine/edge/main > /etc/apk/repositories \
	&& echo http://mirrors.aliyun.com/alpine/edge/community >> /etc/apk/repositories \
	&& apk update \
	&& apk add --no-cache bash nodejs npm git yarn \
	&& rm -rf /var/cache/apk/*

WORKDIR /app

# 难顶的网络
#RUN git clone -b Alex https://github.com/PharbersDeveloper/PBServerPods.git
COPY . .

WORKDIR /app/PBServerPods/server

RUN yarn && npm run build

# prod 源镜像
FROM alpine:latest

# 作者
MAINTAINER Pharbers "pqian@pharbers.com"

LABEL server.version=0.0.2

# 安装 主要 依赖
RUN echo http://mirrors.aliyun.com/alpine/edge/main > /etc/apk/repositories \
	&& echo http://mirrors.aliyun.com/alpine/edge/community >> /etc/apk/repositories \
	&& apk update \
	&& apk add --no-cache tzdata bash nodejs npm \
	&& ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
	&& echo "Asia/Shanghai" > /etc/timezone \
	&& apk del tzdata \
	&& rm -rf /var/cache/apk/*


WORKDIR /app

# 环境变量
ENV PHPRODSHOME /app/PBServerPods
ENV PHSERVER ${PHPRODSHOME}/server

# 提取执行文件
COPY --from=builder /app/PBServerPods/server ${PHSERVER}
COPY --from=builder /app/PBServerPods/conf ${PHPRODSHOME}/conf
COPY --from=builder /app/PBServerPods/log4js.json ${PHPRODSHOME}/log4js.json

WORKDIR ${PHSERVER}/log
WORKDIR ${PHSERVER}

EXPOSE 8080

ENTRYPOINT ["npm", "run", "start"]