#! /bin/bash
GOOS=linux GOARCH=amd64 go build -o bin/casino ./cmd/main.go
GOOS=darwin GOARCH=amd64 go build -o bin/xxxx-mac ./cmd/main.go
GOOS=darwin GOARCH=arm64 go build -o bin/xxxx-mac-m1 ./cmd/main.go


#复制文件到测试服务器
scp bin/casino debug@47.104.227.116:~