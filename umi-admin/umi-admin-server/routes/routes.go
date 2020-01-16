package routes

import (
	gin "github.com/gin-gonic/gin"
	api1 "micro-service-monorepo/gin/routes/api1"
	index "micro-service-monorepo/gin/routes/index"
	mw "micro-service-monorepo/gin/routes/middleware"
	doc "micro-service-monorepo/gin/routes/swagger"
)

// InitRouter ...
func InitRouter() *gin.Engine {
	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// CORS
	r.Use(mw.CORS())

	// middlewares
	r.Use(mw.RequestIdMiddleware())
	r.Use(mw.RequestLoggerMiddleware())
	root := r.Group("/gin")
	{
		// swagger router
		doc.Route(root)

		// custom routers
		index.Route(root)
		api1.Route(root)
	}

	return r
}
