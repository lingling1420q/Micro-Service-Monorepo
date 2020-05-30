package monorepo

//go:generate protoc --go_out=. --micro_out=. pkg/proto/svc1/svc1.proto
//go:generate protoc --go_out=. --micro_out=. pkg/proto/svc2/svc2.proto
