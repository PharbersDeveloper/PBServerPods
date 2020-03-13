FROM alpine:latest

RUN echo http://mirrors.aliyun.com/alpine/edge/main > /etc/apk/repositories \
	&& echo http://mirrors.aliyun.com/alpine/edge/community >> /etc/apk/repositories \
	&& apk update \
	&& apk add --no-cache tzdata git bash nodejs npm yarn \
	&& rm -rf /var/cache/apk/* \
	&& cp -r -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime


LABEL pods-server.version=0.0.1

WORKDIR /app/PBServerPods
COPY ./PBServerPods .
# WORKDIR /app
# RUN git clone https://github.com/PharbersDeveloper/PBServerPods.git

# 环境变量
ENV PHPRODSHOME /app/PBServerPods
ENV PHSERVER ${PHPRODSHOME}/server

# 编译Node项目
WORKDIR ${PHSERVER}
RUN yarn

EXPOSE 8080

ENTRYPOINT ["npm", "run", "start"]