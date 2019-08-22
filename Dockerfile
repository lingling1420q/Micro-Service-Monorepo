# TODO: complete docker file
# docker build -t gin-demo .
FROM golang:latest

WORKDIR /server
COPY . /server
RUN  go build -x -o /server/build/server

EXPOSE 8080

ENTRYPOINT ["/server/build/server"]
