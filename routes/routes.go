package routes

import (
	api1 "monaco/routes/api1"
	index "monaco/routes/index"
	mw "monaco/routes/middleware"
	doc "monaco/routes/swagger"

	gin "github.com/gin-gonic/gin"
)

// InitRouter ...
func InitRouter() *gin.Engine {
	r := gin.Default()

	// CORS
	r.Use(mw.CORS())

	// middlewares
	r.Use(mw.ParseFormMiddleware())
	r.Use(mw.RequestIdMiddleware())
	r.Use(mw.RequestLoggerMiddleware())

	// swagger router
	doc.Route(r)

	// custom routers
	index.Route(r)
	api1.Route(r)

	return r
}
