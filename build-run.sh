#!/bin/bash

export GOPROXY=https://goproxy.io

go build ./...

# 这里要根据操作系统做判断 ...
./app
