#!/bin/sh
path=$(cd "$(dirname "$0")"; pwd)

# Mac
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/modoc-mac -v main.go
chmod +x bin/modoc-mac

# linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/modoc-linux -v main.go
chmod +x bin/modoc-linux

# Win
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/modoc-win.exe -v main.go
chmod +x bin/modoc-win.exe



