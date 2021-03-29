#!/bin/sh

rm -rf build/
mkdir build/

cp /usr/local/lib/libws2811.so build/



mkdir build/monolith/

go build -o=build/monolith/ ./screen/ ./ledstrip/



mkdir build/shared/

go install -buildmode=shared -linkshared ./common/pins/ ./common/utils/
go build -linkshared -o=build/shared/ ./screen/ #./ledstrip/

cp /usr/local/go/pkg/linux_arm_dynlink/*.so build/shared