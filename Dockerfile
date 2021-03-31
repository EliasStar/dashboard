FROM --platform=linux/arm/v6 debian:buster AS lib_builder

WORKDIR /lib

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y build-essential cmake git

RUN git clone https://github.com/jgarff/rpi_ws281x.git && \
    mkdir rpi_ws281x/build/ && \
    cd rpi_ws281x/build/ && \
    cmake -D BUILD_SHARED=OFF -D BUILD_TEST=OFF .. && \
    cmake --build . && \
    make install


FROM golang:1.16-buster

COPY --from=lib_builder /usr/local/lib/libws2811.a /usr/local/lib/
COPY --from=lib_builder /usr/local/include/ws2811 /usr/local/include/ws2811

ENV GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=1 CC=arm-linux-gnueabi-gcc

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y crossbuild-essential-armel

VOLUME [ "/go/src/app/" ]
WORKDIR /go/src/app/

CMD [ "/go/src/app/build.sh" ]