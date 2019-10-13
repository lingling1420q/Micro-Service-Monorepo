package swagger

import (
	_ "monaco/docs" /* doc files */

	gin "github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func Route(r *gin.Engine) {
	r.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/devdoc", func(c *gin.Context) {
		c.Redirect(301, "/doc/index.html")
	})
}
