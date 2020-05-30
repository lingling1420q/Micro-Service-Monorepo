.PHONY:proto
proto:
	@go generate

.PHONY:micro.web
micro.web:
	@micro web

.PHONY:svc1
svc1:
	@go run service/svc1/main.go

.PHONY:svc2
svc2:
	@go run service/svc2/main.go
