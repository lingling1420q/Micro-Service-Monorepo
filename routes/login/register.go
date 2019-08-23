package login

import (
	db "gin-demo/db"
	"gin-demo/logger"

	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @api {get} /user/:id Request User information
 * @apiName GetUser
 * @apiGroup User
 *
 * @apiParam {Number} id Users unique ID.
 *
 * @apiSuccess {String} firstname Firstname of the User.
 * @apiSuccess {String} lastname  Lastname of the User.
 */
//AddUser add user to db
func AddUser(c *gin.Context) {
	name := c.PostForm("name")
	pwd := c.PostForm("pwd")

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

/**
 * @api {get} /user2/:id Request User information
 * @apiName GetUser1
 * @apiGroup User1
 *
 * @apiParam {Number} id Users unique ID.
 *
 * @apiSuccess {String} firstname Firstname of the User.
 * @apiSuccess {String} lastname  Lastname of the User.
 */
//GetUserInfo  get user info by name
func GetUserInfo(c *gin.Context) {
	gormConn := db.ConnGormMysql()

	user := db.TBL_USERS{Name: c.Param("name")}

	gormConn.Select("name, pwd").Where(&user).Find(&user)
	logger.Debug(user)

	//返回数据
	c.JSON(http.StatusOK, user)
}
