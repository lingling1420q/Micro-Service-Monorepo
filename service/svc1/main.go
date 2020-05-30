package main

import (
	"monorepo/pkg/config"
	svc1 "monorepo/pkg/proto/svc1"
	"monorepo/service/svc1/handler"

	"github.com/micro/go-micro/v2"
	"github.com/monaco-io/logger"
)

func main() {
	service := micro.NewService(
		micro.Name(config.Svc1.MicroName),
		micro.Version(config.Svc1.Version),
	)
	service.Init()

	if err := svc1.RegisterSvc1Handler(service.Server(), new(handler.Svc1)); err != nil {
		logger.F("service register failed.", "err", err)
	}

	if err := service.Run(); err != nil {
		logger.F("service run failed.", "err", err)
	}
}
