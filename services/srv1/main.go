package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/monaco-io/logger"
)

func main() {
	service := micro.NewService(
		micro.Registry(svc.NewRegistry()),
		micro.Name(config.Srv1.MicroName),
		micro.Version(config.config.Srv1.Version),
		micro.Address(config.config.Srv1.Address),
	)
	service.Init()

	if err := Srv1.RegisterSrv1Handler(service.Server(), new(handler.Srv1)); err != nil {
		logger.F("service register failed.", "err", err)
	}

	if err := service.Run(); err != nil {
		logger.F("service run failed.", "err", err)
	}
}
