#源镜像
FROM golang:1.12.4-alpine

#作者
MAINTAINER Pharbers "pqian@pharbers.com"

RUN apk add --no-cache git gcc musl-dev mercurial bash gcc g++ make pkgconfig openssl-dev

# 设置工程配置文件的环境变量
#ENV PKG_CONFIG_PATH /usr/lib/pkgconfig
#ENV DOWNLOAD /go/files/
ENV UCB_HOME $GOPATH/src/github.com/PharbersDeveloper/SandBoxServiceDeploy/deploy-config
#ENV BM_KAFKA_CONF_HOME $GOPATH/src/github.com/PharbersDeveloper/SandBoxServiceDeploy/deploy-config/resource/kafkaconfig.json
#ENV BM_XMPP_CONF_HOME $GOPATH/src/github.com/PharbersDeveloper/SandBoxServiceDeploy/deploy-config/resource/xmppconfig.json
ENV GO111MODULE on

ENV LOGGER_USER "Alex"
ENV LOGGER_DEBUG "false"
ENV LOG_PATH $GOPATH/logs

#LABEL
LABEL SandBoxPods.version="0.0.1" maintainer="Alex"


# 下载kafka
#RUN git clone https://github.com/edenhill/librdkafka.git $GOPATH/librdkafka

#WORKDIR $GOPATH/librdkafka
#RUN ./configure --prefix /usr  && \
#make && \
#make install

# 下载依赖
RUN git clone https://github.com/PharbersDeveloper/SandBoxServiceDeploy.git  $GOPATH/src/github.com/PharbersDeveloper/SandBoxServiceDeploy && \
    git clone https://github.com/PharbersDeveloper/SandBoxPods.git $GOPATH/src/github.com/PharbersDeveloper/SandBoxPods

# 构建可执行文件
RUN cd $GOPATH/src/github.com/PharbersDeveloper/SandBoxPods && \
    go build && go install

# 暴露端口
EXPOSE 36415

# 设置工作目录
WORKDIR $GOPATH/bin

ENTRYPOINT ["SandBox"]