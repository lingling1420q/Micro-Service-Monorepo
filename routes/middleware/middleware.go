package middleware

import (
	"encoding/json"
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
		// url parameters
		query := c.Request.URL.Query()
		logger.Noticef("Query: %v\r\n", query)

		// application/x-www-form-urlencoded or application/x-www-form-urlencoded
		formJSON := c.Request.PostForm
		logger.Noticef("JsonForm: %v\r\n", formJSON)

		// application/json
		jsonRaw, _ := c.GetRawData()
		var mapResult map[string]interface{}
		json.Unmarshal(jsonRaw, &mapResult)
		logger.Noticef("JsonRaw: %v\r\n", mapResult)

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
