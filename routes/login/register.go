package login

import (
	db "gin-demo/db"
	"gin-demo/logger"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Add user godoc
// @Summary user
// @Description add user
// @Tags user
// @Accept  json
// @Produce  json
// @Param  user  body db.TBL_USERS true "add user"
// @Success 200 {object} db.TBL_USERS
// @Failure 400 {object} defs.HTTPError
// @Router /login/register [post]
//AddUser add user to db
func AddUser(c *gin.Context) {
	name := c.PostForm("name")
	pwd := c.PostForm("pwd")
	logger.Debug(name)
	gormConn := db.ConnGormMysql()
	// 构建User
	user := db.TBL_USERS{Name: name, Pwd: pwd}
	// 存入数据库
	gormConn.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"pwd":  pwd,
	})

}

//GetUserInfo  get user info by name
func GetUserInfo(c *gin.Context) {
	gormConn := db.ConnGormMysql()

	user := db.TBL_USERS{Name: c.Param("name")}

	gormConn.Select("name, pwd").Where(&user).Find(&user)
	logger.Debug(user)

	//返回数据
	c.JSON(http.StatusOK, user)
}
