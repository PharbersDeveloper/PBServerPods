
# builder 源镜像
FROM golang:1.12.4-alpine as builder

# 作者
MAINTAINER Pharbers "pqian@pharbers.com"

RUN echo http://mirrors.aliyun.com/alpine/edge/main > /etc/apk/repositories \
	&& echo http://mirrors.aliyun.com/alpine/edge/community >> /etc/apk/repositories \
	&& apk update \
	&& apk add --no-cache bash tzdata bash git gcc g++ openssl-dev librdkafka-dev pkgconf \
	&& rm -rf /var/cache/apk/* \
	&& cp -r -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# 环境变量
ENV GOPROXY https://goproxy.io
ENV GO111MODULE on

# WORKDIR /go/src/github.com/PharbersDeveloper/PBServerPods
WORKDIR /go/src/github.com/PharbersDeveloper/
RUN git clone -b Alex https://github.com/PharbersDeveloper/PBServerPods.git

ENV PHPRODSHOME /go/src/github.com/PharbersDeveloper/PBServerPods
ENV PHSANDBOX ${PHPRODSHOME}/modules/PhSandBox
ENV EMAIL_TEMPLATE=${PHPRODSHOME}/conf/email-template.html
ENV EMAILADDRESS=${PHPRODSHOME}/conf/emails.json

# COPY ./PBServerPods ${PHPRODSHOME}


# Kafka Config
ENV KAFKA_HOME_CONFIG=${PHPRODSHOME}/conf/kafka_config.json
ENV BP_KAFKA_SCHEMA_REGISTRY_URL http://pharbers.com:8081


WORKDIR ${PHSANDBOX}
RUN go build

# pord 镜像
FROM alpine:latest

RUN echo http://mirrors.aliyun.com/alpine/edge/main > /etc/apk/repositories \
	&& echo http://mirrors.aliyun.com/alpine/edge/community >> /etc/apk/repositories \
	&& apk update \
	&& apk add --no-cache bash tzdata gcc g++ openssl-dev librdkafka-dev pkgconf \
	&& rm -rf /var/cache/apk/* \
	&& cp -r -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime


WORKDIR /app/PBServerPods

ENV PHPRODSHOME /app/PBServerPods
ENV EMAIL_TEMPLATE=${PHPRODSHOME}/conf/email-template.html
ENV EMAILADDRESS=${PHPRODSHOME}/conf/emails.json

ENV PROJECT_NAME PhSandBox
ENV BP_LOG_TIME_FORMAT "2006-01-02 15:04:05"
ENV BP_LOG_OUTPUT /app/PBServerPods/log/PhSandBox.log
ENV BP_LOG_LEVEL info

# Kafka Config
ENV KAFKA_HOME_CONFIG=${PHPRODSHOME}/conf/kafka_config.json
ENV BP_KAFKA_SCHEMA_REGISTRY_URL http://pharbers.com:8081

WORKDIR ${PHPRODSHOME}/log

WORKDIR ${PHPRODSHOME}/conf

# 提取资源文件
# COPY --from=0 /go/src/github.com/PharbersDeveloper/PBServerPods/conf/email-template.html ./
# COPY --from=0 /go/src/github.com/PharbersDeveloper/PBServerPods/conf/emails.json ./
# COPY --from=0 /go/src/github.com/PharbersDeveloper/PBServerPods/conf/kafka_config.json ./

WORKDIR /go/bin

# COPY --from=0 /go/src/github.com/PharbersDeveloper/PBServerPods/conf/snakeoil-ca-1.crt ./
# COPY --from=0 /go/src/github.com/PharbersDeveloper/PBServerPods/conf/kafkacat-ca1-signed.pem ./
# COPY --from=0 /go/src/github.com/PharbersDeveloper/PBServerPods/conf/kafkacat.client.key ./

# 提取执行文件
COPY --from=0 /go/src/github.com/PharbersDeveloper/PBServerPods/modules/PhSandBox/PhSandBox ./

EXPOSE 30001

ENTRYPOINT ["./PhSandBox"]