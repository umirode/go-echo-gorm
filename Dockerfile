FROM golang:1.11.1-alpine3.8 as builder

ENV GOLANG_APP_REPOSITORY='github.com/umirode/go-rest'

WORKDIR $GOPATH/src/$GOLANG_APP_REPOSITORY

COPY Gopkg.toml Gopkg.lock ./
RUN \
    apk update && \
    apk upgrade && \
    \
    apk add --no-cache \
        bash && \
    \
    apk add --no-cache --virtual .build-dependencies \
        libc-dev \
        gcc \
        git \
        dep && \
    \
    dep ensure --vendor-only

COPY . ./
RUN \
    GOOS=linux \
    go build -o /build/app . && \
    \
    cp .env /build/ && \
    cp wait-for-it.sh /build/ && \
    \
    apk del .build-dependencies

CMD /build/wait-for-it.sh --host=${DATABASE_HOST} --port=${DATABASE_PORT} --timeout=60 -- /build/app