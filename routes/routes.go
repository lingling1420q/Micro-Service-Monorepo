package routes

import (
	api1 "gin-demo/routes/api1"
	api2 "gin-demo/routes/api2"
	login "gin-demo/routes/login"
	mw "gin-demo/routes/middleware"
	index "gin-demo/routes/index"
	doc "gin-demo/routes/swagger"

	gin "github.com/gin-gonic/gin"
)

// InitRouter ...
func InitRouter() *gin.Engine {
	r := gin.Default()

	// middlewares
	r.Use(mw.ParseFormMiddleware())
	r.Use(mw.RequestIdMiddleware())
	r.Use(mw.RequestLoggerMiddleware())

	// swagger router
	doc.Route(r)

	// custom routers
	index.Route(r)
	login.Route(r)
	api1.Route(r)
	api2.Route(r)
	return r
}
