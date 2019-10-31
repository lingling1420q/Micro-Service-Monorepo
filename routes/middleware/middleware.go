package middleware

import (
	"fmt"
	"monaco/db"
	"monaco/logger"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func ParseFormMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.ParseMultipartForm(1024)
		c.Next()
	}
}

func RequestIdMiddleware() gin.HandlerFunc {
	// Do some initialization logic here
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Request-Id", uuid.NewV4().String())
		c.Next()
	}
}

func RequestLoggerMiddleware() gin.HandlerFunc {
	// Do some initialization logic here

	return func(c *gin.Context) {
		query := c.Request.URL.Query()
		logger.Noticef("Query: %v\r\n", query)

		formJSON := c.Request.PostForm
		logger.Noticef("JsonForm: %v\r\n", formJSON)

		jsonRaw, _ := c.GetRawData()
		logger.Noticef("JsonData: %v\r\n", string(jsonRaw))

		vl := db.TBL_VISIT_LOG{
			Host:     c.Request.Host,
			Method:   c.Request.Method,
			URL:      fmt.Sprintf("%v", c.Request.URL),
			Header:   fmt.Sprintf("%v", c.Request.Header),
			Query:    fmt.Sprintf("%v", query),
			JSONForm: fmt.Sprintf("%v", formJSON),
			JSONRaw:  fmt.Sprintf("%v", string(jsonRaw)),
		}
		db.InsertColume(&vl)

		c.Next()
	}
}
