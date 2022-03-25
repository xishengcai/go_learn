FROM registry.cn-hangzhou.aliyuncs.com/launcher/alpine:latest

MAINTAINER xishengcai <cc710917049@163.com>

COPY ./bin/sockjs /usr/local/bin
COPY ./web /opt/web
COPY ./cert /opt/cert

EXPOSE 8080/tcp
EXPOSE 443/tcp

RUN chmod +x /usr/local/bin/sockjs

WORKDIR /opt

ENTRYPOINT ["sockjs"]
