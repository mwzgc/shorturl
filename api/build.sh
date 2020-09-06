#!/bin/bash

GOOS=linux go build shorturl.go
docker build -t shorturl:latest .
rm shorturl
docker image prune -f
