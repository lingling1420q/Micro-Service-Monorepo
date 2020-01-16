package swagger

import (
	_ "micro-service-monorepo/gin/docs" /* doc files */
	"net/http"

	gin "github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func Route(r *gin.RouterGroup) {
	r.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/devdoc", func(c *gin.Context) {
		c.Redirect(http.StatusSeeOther, "/gin/doc/index.html")
	})
}
