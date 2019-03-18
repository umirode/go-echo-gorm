FROM golang:1.12.1-alpine3.9

WORKDIR /root/app

COPY . ./

RUN \
    apk update && \
    apk upgrade && \
    \
    apk add --no-cache --virtual .build-dependencies \
        libc-dev \
        gcc \
        git && \
    \
    GOOS=linux \
    go build -i -o /build/app . && \
    go build -i -o /build/cmd Cli/main.go && \
    \
    cp .env database.yaml /build/ && \
    mkdir /build/ignore && \
    \
    apk del .build-dependencies

WORKDIR /build

CMD /build/app