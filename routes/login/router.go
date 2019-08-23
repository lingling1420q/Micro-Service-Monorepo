package login

import (
	"net/http"

	gin "github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	router := r.Group("/login")
	ctrl := &controller{}
	{
		router.GET("/wx", ctrl.ctrlWx)
		router.GET("/qywx", ctrl.ctrlQywx)
		router.GET("/qywxtp", ctrl.ctrlQywxThirdParty)
		router.POST("/register", AddUser)
		router.GET("/get/:name", GetUserInfo)
	}
}

type controller struct{}

func (*controller) ctrlWx(c *gin.Context) {
	queryMap := c.Request.URL.Query()
	c.String(http.StatusOK, Wx(queryMap))
}

//Qywx user login
func (*controller) ctrlQywx(c *gin.Context) {
	queryMap := c.Request.URL.Query()
	resp := Qywx(queryMap)
	c.String(http.StatusOK, resp)
}

func (*controller) ctrlQywxThirdParty(c *gin.Context) {
	queryMap := c.Request.URL.Query()
	resq := QywxThirdParty(queryMap)
	c.String(http.StatusOK, resq)
}
