#!/usr/bin/env bash


CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -o ./build/win32/vm.exe main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -o ./build/win64/vm.exe main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./build/mac/vm main.go
