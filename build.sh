#!/bin/sh

rm -rf build/
mkdir build/

# Desktop
mkdir build/windows/

export GOOS=windows
export GOARCH=amd64

cd DashConnect/
go build -o=../build/windows/ .
cd ..

mkdir build/linux/

export GOOS=linux
export GOARCH=amd64

cd DashConnect/
go build -o=../build/linux/ .
cd ..

# Dashboard
mkdir build/dashboard/

export GOOS=linux
export GOARCH=arm
export GOARM=6
export CGO_ENABLED=1
export CC=arm-linux-gnueabi-gcc

cd ScrnBtn/
go build -o=../build/dashboard/ .
cd ..

cd LedstripCtrl/
go build -o=../build/dashboard/ .
cd ..

cd DashD/
go build -o=../build/dashboard/ .
cd ..

unset GOOS
unset GOARCH
unset GOARM
unset CGO_ENABLED
unset CC