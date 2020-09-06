#!/bin/bash
./build.sh
docker-compose up --scale transform=3
