FROM ubuntu:trusty
MAINTAINER  Jayson  <yjge@qingyuanos.com>

RUN echo "deb http://mirrors.aliyun.com/ubuntu precise main universe" > /etc/apt/sources.list
RUN apt-get update
RUN apt-get install -y nginx

COPY ingress-controller  /
COPY ./provider/nginx/nginx_template.cfg  /etc/nginx/
CMD ["/ingress-controller", "--lb-provider=nginx"]
