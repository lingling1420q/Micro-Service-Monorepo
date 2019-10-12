package index

import gin "github.com/gin-gonic/gin"

func Route(r *gin.Engine) {
	r.Any("/", Index)
}
