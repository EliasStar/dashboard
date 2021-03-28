FROM --platform=linux/arm/v6 debian:buster AS lib_builder

WORKDIR /lib

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y build-essential cmake git

RUN git clone https://github.com/jgarff/rpi_ws281x.git && \
    cd rpi_ws281x && \
    sed -i "s/ARCHIVE DESTINATION \${DEST_LIB_DIR}/ARCHIVE DESTINATION \${DEST_LIB_DIR}\n    LIBRARY DESTINATION \${DEST_LIB_DIR}/" CMakeLists.txt && \
    mkdir build/ && \
    cd  build/ && \
    cmake -D BUILD_SHARED=ON -D BUILD_TEST=OFF .. && \
    cmake --build . && \
    make install


FROM golang:1.16-buster

COPY --from=lib_builder /usr/local/lib/libws2811.so /usr/local/lib/
COPY --from=lib_builder /usr/local/include/ws2811 /usr/local/include/ws2811

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y build-essential

CMD [ "bash" ]

# VOLUME [ "/go/src/app/" ]
# WORKDIR /go/src/app/

# CMD [ "/go/src/app/build.sh" ]