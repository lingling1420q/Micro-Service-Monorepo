package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/web"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

// func main() {
// 	service := web.NewService(
// 		web.Name("go.micro.api.gin"),
// 		web.Version("0.0.1"),
// 	)
// 	service.Init()
// 	serverCfg := config.Config.Server

// 	if !serverCfg.Debug {
// 		gin.SetMode(gin.ReleaseMode)
// 	}
// 	logger.Noticef("Gin is Running in Debug mode: %v", serverCfg.Debug)
// 	service.Handle("/", routes.InitRouter())

// 	if err := service.Run(); err != nil {
// 		log.Fatalf("Start Failed: %v", err)
// 	}
// }

func main() {
	// Create service
	service := web.NewService(
		web.Name("go.micro.api.gin"),
	)

	service.Init()

	// setup gin Server Client
	cl = hello.NewSayService("go.micro.srv.gin", client.DefaultClient)

	// Create RESTful handler (using Gin)
	router := gin.Default()
	router.GET("/gin", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "gin"}) })
	router.GET("/gin/:name", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "gin2"}) })

	// Register Handler
	service.Handle("/", router)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
