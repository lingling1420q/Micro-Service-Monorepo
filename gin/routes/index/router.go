package index

import gin "github.com/gin-gonic/gin"

func Route(r *gin.RouterGroup) {
	r.Any("/", Index)
}
