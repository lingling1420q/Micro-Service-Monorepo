package api1

import (
	"log"
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

// Index ...// @Summary Add new user to the database
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Router /user [post]
func Index(c *gin.Context) {
	// conn := db.ConnMysql()
	// log.Println(conn)
	// c.String(http.StatusOK, "hello, index, api1")
	c.HTML(http.StatusOK, "doc/index.html", nil)
}

//Helloworld ...
func Helloworld(c *gin.Context) {
	log.Print()
	c.String(http.StatusOK, "hello, world, api1")
}
