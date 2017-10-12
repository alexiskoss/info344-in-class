#!/usr/bin/env bash
set -e 
echo "building linux executable"
GOOS=linux go build
docker build -t alexiskoss/zipsvr .
docker push alexiskoss/zipsvr
go clean