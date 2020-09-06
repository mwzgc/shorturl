#!/bin/bash

docker-compose pull
docker-compose up -d --scale transform=3
