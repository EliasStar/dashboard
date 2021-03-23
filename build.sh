#!/bin/sh

GOOS=linux
GOARCH=arm
GOARM=6

go build -o ./build/screen ./screen
go build -o ./build/ledstrip ./ledstrip