package main

import (
	"log"
	config "micro-service-monorepo/admin/server/config"
	logger "micro-service-monorepo/admin/server/logger"
	routes "micro-service-monorepo/admin/server/routes"

	"github.com/gin-gonic/gin"
	web "github.com/micro/go-micro/web"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

// @title Web Server Framework API
// @version 1.0.1

// @contact.name API Support
// @contact.url https://github.com/luxuze/Web-Server-Framework-Golang
// @contact.email luxuze@agora.io

// @license.name DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
// @license.url https://github.com/luxuze/Web-Server-Framework-Golang/blob/master/LICENSE

// @host localhost:4096
// @BasePath /
// func main() {
// 	serverCfg := config.Config.Server

// 	if !serverCfg.Debug {
// 		gin.SetMode(gin.ReleaseMode)
// 	}
// 	logger.Noticef("Gin is Running in Debug mode: %v", serverCfg.Debug)

// 	s := &http.Server{
// 		Addr:           serverCfg.Port,
// 		Handler:        routes.InitRouter(),
// 		ReadTimeout:    10 * time.Second,
// 		WriteTimeout:   10 * time.Second,
// 		MaxHeaderBytes: 1 << 20,
// 	}
// 	logger.Noticef("Server Listing %s", serverCfg.Port)
// 	if err := s.ListenAndServe(); err != nil {
// 		logger.Errorf("Failed to Start Server: %s", err)
// 	}
// }

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
