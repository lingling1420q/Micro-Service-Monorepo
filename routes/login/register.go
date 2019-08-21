package login

import (
	"encoding/json"
	db "gin-demo/db"

	"net/http"

	"github.com/gin-gonic/gin"
)

//AddUser add user to db
func AddUser(c *gin.Context) {
	name := c.PostForm("name")
	pwd := c.PostForm("pwd")

	gormConn := db.ConnGormMysql()
	// gormConn.CreateTable(&defs.TBL_USERS{})
	// 构建User
	user := db.TBL_USERS{Name: name, Pwd: pwd}
	// 设置表名为单数形式，如user表，不设置为users
	// 存入数据库
	gormConn.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"pwd":  pwd,
	})

}

//GetUserInfo  get user info by ID
func GetUserInfo(c *gin.Context) {
	userID := c.Param("id")
	gormConn := db.ConnGormMysql()
	user := db.TBL_USERS{}
	// select from user by ID
	gormConn.Find(&user, userID)
	userInfo, _ := json.Marshal(user)
	//返回数据
	c.JSON(http.StatusOK, gin.H{
		"userInfo": string(userInfo),
	})
}
