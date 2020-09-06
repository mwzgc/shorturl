#!/bin/bash

cd api
./build.sh
cd ..

cd rpc/transform
./build.sh
cd ../..
