# docker build -t 17199911145/monaco .
FROM golang:latest

WORKDIR /app
COPY . .

ENV GOPROXY="https://goproxy.io"
ENV APP_ENV="PRODUCTION"

RUN go mod download
RUN go build -x -o admin_build

EXPOSE 9099

ENTRYPOINT ["/app/admin_build"]
