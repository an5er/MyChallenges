#!/bin/bash

service docker start;
service nginx start;
sleep 8;
docker build -t sctf/filenode1 -f /app/node/Dockerfile .;
docker run -d -p 8080:8080 sctf/filenode1;
node /app/server/server.js;