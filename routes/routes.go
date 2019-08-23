package routes

import (
	api1 "gin-demo/routes/api1"
	api2 "gin-demo/routes/api2"
	login "gin-demo/routes/login"

	_ "gin-demo/docs"

	gin "github.com/gin-gonic/gin"
	// ginSwagger "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/gin-swagger/swaggerFiles"
)
import ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
import swaggerFiles "github.com/swaggo/files"     // swagger embed files
// InitRouter ...
func InitRouter() *gin.Engine {
	r := gin.Default()

	login.Route(r)
	api1.Route(r)
	api2.Route(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
