package routes

import (
	p1 "demo/routes/api"

	gin "github.com/gin-gonic/gin"
)

// InitRouter ...
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	apip1 := r.Group("/api/p1")
	{
		apip1.GET("/", p1.Index)
		apip1.GET("/hello", p1.Helloworld)
	}
	return r
}
