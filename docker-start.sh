#!/bin/bash
set -e
GOOS=linux go build chrisgr.go
set +e
docker rm chrisgr
set -e
docker build -t chrisgr .
docker run -t -p 8080:8080 --name chrisgr -it chrisgr
