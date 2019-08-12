package api1

import (
	db "gin-demo/db"
	"log"
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

// Index ...
func Index(c *gin.Context) {
	conn := db.ConnMysql()
	log.Println(conn)
	c.String(http.StatusOK, "hello, index, api1")
}

//Helloworld ...
func Helloworld(c *gin.Context) {
	log.Print()
	c.String(http.StatusOK, "hello, world, api1")
}
