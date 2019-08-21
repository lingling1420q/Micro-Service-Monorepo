package routes

import (
	api1 "gin-demo/routes/api1"
	api2 "gin-demo/routes/api2"
	login "gin-demo/routes/login"
	user "gin-demo/routes/user"

	gin "github.com/gin-gonic/gin"
)

// InitRouter ...
func InitRouter() *gin.Engine {
	r := gin.Default()

	login.Route(r)
	api1.Route(r)
	api2.Route(r)
	user.Route(r)
	return r
}
