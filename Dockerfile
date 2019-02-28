FROM golang:1.11.1-alpine3.8

RUN \
    apk update && \
    apk upgrade && \
    \
    apk add --no-cache --virtual .build-dependencies \
        libc-dev \
        gcc \
        git \
        dep

ENV APP_REPOSITORY 'github.com/umirode/go-rest'

WORKDIR $GOPATH/src/${APP_REPOSITORY}

COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only

COPY . ./
RUN \
    GOOS=linux \
    go build -i -o /build/app . && \
    go build -i -o /build/cmd Cli/main.go && \
    \
    cp .env database.yaml /build/ && \
    cp -R ignore /build/ && \
    \
    apk del .build-dependencies

WORKDIR /build

CMD /build/app