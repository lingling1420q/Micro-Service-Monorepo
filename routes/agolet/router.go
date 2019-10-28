package agolet

import gin "github.com/gin-gonic/gin"

func Route(r *gin.Engine) {
	router := r.Group("/agolet")
	{
		router.GET("/", Index)
		router.GET("/hw", Helloworld)
	}
}
