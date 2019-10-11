package print

import (
	db "gin-demo/db"
	logger "gin-demo/logger"
	"io/ioutil"
	http "net/http"

	gin "github.com/gin-gonic/gin"
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
	logger.Noticef("JsonForm: %v\r\n", c.Request.PostForm)

	data, _ := ioutil.ReadAll(c.Request.Body)
	logger.Noticef("BodyForm: %v\r\n", string(data))

	vl := db.TBL_VISIT_LOG{
		Host:  c.Request.Host,
		Query: "query",
		Body:  string(data),
	}
	db.InsertColume(&vl)

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
