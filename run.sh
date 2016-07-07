#!/bin/bash

export KUBERNETES_URL='http://localhost:8080'
export HAPROXY_CONFIG='/etc/haproxy/haproxy.cfg'
./ingress-controller
