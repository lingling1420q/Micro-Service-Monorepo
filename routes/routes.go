package routes

import (
	api1 "monaco/routes/api1"
	api2 "monaco/routes/api2"
	login "monaco/routes/login"
	mw "monaco/routes/middleware"
	index "monaco/routes/index"
	doc "monaco/routes/swagger"

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
