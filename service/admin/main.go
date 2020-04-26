package main

import (
	"log"
	config "monorepo/config"
	logger "monorepo/pkg/logger"
	routes "monorepo/service/admin/routes"

	"github.com/gin-gonic/gin"
	web "github.com/micro/go-micro/web"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	service := web.NewService(
		web.Name("go.micro.api.gin"),
		web.Version("0.0.1"),
	)
	service.Init()
	serverCfg := config.Config.Server

	if !serverCfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	logger.Noticef("Gin is Running in Debug mode: %v", serverCfg.Debug)
	service.Handle("/", routes.InitRouter())

	if err := service.Run(); err != nil {
		log.Fatalf("Start Failed: %v", err)
	}
}
