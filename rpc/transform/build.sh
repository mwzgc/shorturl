#!/bin/bash

GOOS=linux go build transform.go
docker build -t transform:latest .
rm transform
docker image prune -f
