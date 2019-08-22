package main

import (
	config "gin-demo/config"
	logger "gin-demo/logger"
	routes "gin-demo/routes"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	/* loading toml configs */
	cfg := config.Config()

	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	logger.Noticef("gin is debug mode ? %v", cfg.Debug)

	s := &http.Server{
		Addr:           cfg.Server.Port,
		Handler:        routes.InitRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logger.Noticef("Server Listing %s", cfg.Server.Port)
	s.ListenAndServe()
}
