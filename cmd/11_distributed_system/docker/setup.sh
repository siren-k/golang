#!/usr/bin/env bash

pushd .
env GOOS=linux go build -ldflags "-X main.version=1.0 -X main.builddate=$(date +%s)"
popd
docker build -t example .
docker run -d --name example -p 8000:8000 example