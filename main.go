package main

import (
	"log"
	config "monaco/config"
	logger "monaco/logger"
	routes "monaco/routes"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
