FROM debian AS lib_builder

WORKDIR /lib

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y build-essential cmake git

RUN git clone https://github.com/jgarff/rpi_ws281x.git && \
    mkdir rpi_ws281x/build && \
    cd  rpi_ws281x/build && \
    cmake -D BUILD_SHARED=OFF -D BUILD_TEST=OFF .. && \
    cmake --build . && \
    make install


FROM golang

COPY --from=lib_builder /usr/local/lib/libws2811.a /usr/local/lib/
COPY --from=lib_builder /usr/local/include/ws2811 /usr/local/include/ws2811

VOLUME [ "/app" ]

WORKDIR /app
ENTRYPOINT /app/build.sh