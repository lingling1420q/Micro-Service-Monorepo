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

/**
 * @api {get} /test/:id Test api doc
 * @apiName GetUser
 * @apiGroup HAHA
 *
 * @apiParam {Number} id Users unique ID.
 *
 * @apiSuccess {String} firstname Firstname of the User.
 * @apiSuccess {String} lastname  Lastname of the User.
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "firstname": "John",
 *       "lastname": "Doe"
 *     }
 *
 * @apiError UserNotFound The id of the User was not found.
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 404 Not Found
 *     {
 *       "error": "UserNotFound"
 *     }
 */
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
