package main

import (
	"gin-demo/config"
	"gin-demo/logger"
	"gin-demo/routes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
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
