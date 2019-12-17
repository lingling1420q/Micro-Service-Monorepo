# docker build -t 17199911145/monaco .
FROM golang:latest

WORKDIR /server
COPY . /server

ENV GOPROXY="https://goproxy.io"
ENV APP_ENV="PRODUCTION"

RUN go mod download
RUN go build -x -o /server/build/server

EXPOSE 9099

ENTRYPOINT ["/server/build/server"]
