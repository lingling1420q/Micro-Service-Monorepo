
GOPATH:=$(shell go env GOPATH)
MODIFY=Mgithub.com/micro/go-micro/api/proto/api.proto=github.com/micro/go-micro/v2/api/proto

.PHONY: proto
proto:
    
	protoc --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/svc2/svc2.proto
    

.PHONY: build
build: proto

	go build -o svc2-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t svc2-service:latest
