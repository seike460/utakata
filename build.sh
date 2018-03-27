#!/bin/sh
echo "build-start" $(date +"%Y/%m/%d %H:%M:%S")
cd src/getItem && GOOS=linux GOARCH=amd64 go build -o ../../handlers/getItem.handler
cd ../../
cd src/setItem && GOOS=linux GOARCH=amd64 go build -o ../../handlers/setItem.handler
echo "build-end  " $(date +"%Y/%m/%d %H:%M:%S")
