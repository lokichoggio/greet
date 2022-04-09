#!/bin/bash

# infra
REGISTRY_URL="harbor.infra.ww5sawfyut0k.bitsvc.io/zhangqi"

#echo "build"
#
#GOOS=linux GOARCH=amd64 go build -o greet

echo "delete old image"

docker rm -f `docker ps -a | grep greet | awk '{print $1}'`

docker rmi greet

echo ">>> build image"

docker build . -t greet -f ./Dockerfile

echo ">>> push image"

docker tag greet $REGISTRY_URL/greet

docker push $REGISTRY_URL/greet

echo ">>> delete local image"

docker rmi greet

docker rmi $REGISTRY_URL/greet
