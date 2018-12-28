#!/usr/bin/env bash
echo "编译中：..."
CGO_ENABLED=0 go build -a -o gin-blog
echo "编译完成：$PWD/gin-blog"