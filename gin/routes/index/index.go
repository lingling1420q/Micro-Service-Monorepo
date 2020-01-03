package index

import (
	"log"
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	log.Println(c)
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
