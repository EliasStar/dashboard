#!/bin/sh

GOOS=linux
GOARCH=arm
GOARM=6
CGO_ENABLED=1
# CC=arm-linux-gnueabi-gcc

go build -buildmode=shared -o ./build/ ./pins/ ./utils/
go build -linkshared -o ./build/ ./screen/ ./ledstrip/