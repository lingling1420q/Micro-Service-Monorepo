package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTags(c *gin.Context) {
}

func AddTag(c *gin.Context) {
}

func EditTag(c *gin.Context) {
}

func DeleteTag(c *gin.Context) {
}

func Helloworld(c *gin.Context) {
	c.String(http.StatusOK, "hello, world")
}
