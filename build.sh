#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o vm main.go
go build -o vm_mac main.go