package main

import (
	"monorepo/service/svc2/handler"
	"monorepo/service/svc2/subscriber"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	"monorepo/pkg/config"
	svc2 "monorepo/pkg/proto/svc2"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name(config.Svc2.MicroName),
		micro.Version(config.Svc2.Version),
	)

	// Initialise service
	service.Init()

	// Register Handler
	svc2.RegisterSvc2Handler(service.Server(), new(handler.Svc2))

	// Register Struct as Subscriber
	micro.RegisterSubscriber(config.Svc2.MicroName, service.Server(), new(subscriber.Svc2))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
