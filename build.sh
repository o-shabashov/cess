#!/usr/bin/env bash
env GOOS=windows GOARCH=amd64 go build -o bin/cess.exe
env GOOS=linux GOARCH=amd64 go build -o bin/cess.sh
env GOOS=darwin GOARCH=amd64 go build -o bin/cess.mac