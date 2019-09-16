package print

import (
	http "net/http"
	logger "gin-demo/logger"
	gin "github.com/gin-gonic/gin"
)

func IndexGet(c *gin.Context) {
	query := c.Request.URL.Query()
	logger.Info(query)
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func IndexPost(c *gin.Context) {
	query := c.Request.URL.Query()
	form := c.PostFormMap("")
	logger.Info(query)
	logger.Info(form)
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
