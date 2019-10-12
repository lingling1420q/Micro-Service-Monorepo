package swagger

import (
	_ "monaco/docs" /* doc files */

	gin "github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func Route(r *gin.Engine) {
	//TODO: redirect to index.html
	r.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
