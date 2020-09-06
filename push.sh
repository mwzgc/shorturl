#!/bin/bash

docker login --username=mawenzhong

./build.sh

docker tag shorturl mawenzhong/go_shorturl:latest
docker push mawenzhong/go_shorturl:latest

docker tag transform mawenzhong/go_transform:latest
docker push mawenzhong/go_transform:latest

docker image prune -f