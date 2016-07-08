#!/bin/bash

docker run -d --net=host -e KUBERNETES_URL=http://localhost:8080 -e NGINX_CONFIG=/etc/nginx/nginx.conf   hub.qingyuanos.com/admin/ingress-controller:nginx
