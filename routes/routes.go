package routes

import (
	agolet "monaco/routes/agolet"
	api1 "monaco/routes/api1"
	api2 "monaco/routes/api2"
	index "monaco/routes/index"
	login "monaco/routes/login"
	mw "monaco/routes/middleware"
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
	agolet.Route(r)
	return r
}
