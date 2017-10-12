#!/usr/bin/env bash
set -e 
echo "building linux executable"
GOOS=linux go build
docker build -t alexiskoss/testserver .
docker push alexiskoss/testserver
go clean