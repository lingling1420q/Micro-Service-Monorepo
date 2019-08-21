package user

import gin "github.com/gin-gonic/gin"

func Route(r *gin.Engine) {

	router := r.Group("/user")
	{
		router.POST("/register", AddUser)
		router.GET("/get/:id", GetUserInfo)
	}
}
