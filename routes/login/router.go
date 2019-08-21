package login

import gin "github.com/gin-gonic/gin"

func Route(r *gin.Engine) {
	router := r.Group("/login")
	{
		router.GET("/wx", Wx)
		router.GET("/qywx", Qywx)
		router.GET("/qywxtp", QywxThirdParty)
		router.POST("/register", AddUser)
		router.GET("/get/:name", GetUserInfo)
	}
}
