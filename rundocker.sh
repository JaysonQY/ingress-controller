#!/bin/bash

docker run -d --net=host -e KUBERNETES_URL=http://localhost:8080 -e HAPROXY_CONFIG=/etc/haproxy/haproxy.cfg  hub.qingyuanos.com/admin/ingress-controller:haproxy
