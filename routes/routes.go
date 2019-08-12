package routes

import (
	api1 "gin-demo/routes/api1"
	api2 "gin-demo/routes/api2"
	login "gin-demo/routes/login"

	gin "github.com/gin-gonic/gin"
)

// InitRouter ...
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	loginRouter := r.Group("/login")
	{
		loginRouter.GET("/wx", login.Wx)
		loginRouter.GET("/qywx", login.Qywx)
		loginRouter.GET("/qywxtp", login.QywxThirdParty)
	}

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
	// a3 := r.Group("/user")
	// {
	// 	a3.POST("/register", user.AddUser)
	// 	a3.GET("/get/:id", user.GetUserInfo)
	// }

	return r
}
