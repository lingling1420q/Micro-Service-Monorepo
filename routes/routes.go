package routes

import (
	api1 "gin-demo/routes/api1"
	api2 "gin-demo/routes/api2"

	gin "github.com/gin-gonic/gin"
)

// InitRouter ...
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	a1 := r.Group("/api1")
	{
		a1.GET("/", api1.Index)
		a1.GET("/h1", api1.Helloworld)
	}
	a2 := r.Group("/api2")
	{
		a2.GET("/", api2.Index)
		a2.GET("/h2", api2.Helloworld)
	}
	return r
}
