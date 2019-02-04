FROM golang:1.11.1-alpine3.8 as builder

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
    go build -i -o /build/app . && \
    \
    cp .env /build/ && \
    \
    apk del .build-dependencies

ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /build/wait.sh
RUN chmod +x /build/wait.sh

CMD /build/wait.sh --host=${DATABASE_HOST} --port=${DATABASE_PORT} --timeout=60 -- /build/app