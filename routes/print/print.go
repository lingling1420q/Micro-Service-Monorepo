package print

import (
	http "net/http"
	logger "gin-demo/logger"
	gin "github.com/gin-gonic/gin"
	"io/ioutil"
)

func IndexGet(c *gin.Context) {
	query := c.Request.URL.Query()
	logger.Notice(query)
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func IndexPost(c *gin.Context) {
	query := c.Request.URL.Query()
	logger.Noticef("Query: %v\r\n", query)
	c.Request.ParseForm()
	logger.Noticef("JsonForm: %v\r\n", c.Request.PostForm )
	data, _ := ioutil.ReadAll(c.Request.Body)
	logger.Noticef("BodyForm: %v\r\n", string(data))
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
