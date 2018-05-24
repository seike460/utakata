#!/bin/sh
echo "build-start" $(date +"%Y/%m/%d %H:%M:%S")
cd src/utakata && GOOS=linux GOARCH=amd64 go build -o ../../handlers/utakata.handler
echo "build-end  " $(date +"%Y/%m/%d %H:%M:%S")
