#!/bin/bash

export KUBERNETES_URL='http://localhost:8080'
#export HAPROXY_CONFIG='/etc/haproxy/haproxy.cfg'
export NGINX_CONFIG='/etc/nginx/nginx.conf'
./ingress-controller --lb-provider=nginx
