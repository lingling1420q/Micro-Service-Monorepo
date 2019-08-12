package user

import (
	db "gin-demo/db"
	"gin-demo/defs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AddUser add user to db
func AddUser(c *gin.Context) {
	name := c.PostForm("name")
	pwd := c.PostForm("pwd")
	//connection := db.ConnMysql()
	//插入数据库
	// stmtIns, err := connection.Prepare("INSERT INTO user (name, pwd) VALUES (?, ?)")
	// if err != nil {
	// 	return
	// }
	// _, err = stmtIns.Exec(name, pwd)
	// if err != nil {
	// 	return
	// }
	//返回数据
	// c.JSON(http.StatusOK, gin.H{
	// 	"name": name,
	// 	"pwd":  pwd,
	// })
	//defer stmtIns.Close()

	gormConn := db.ConnGormMysql()
	// 构建User
	user := defs.User{Name: name, Pwd: pwd}
	// 设置表名为单数形式，如user表，不设置为users
	gormConn.SingularTable(true)
	// 存入数据库
	gormConn.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"pwd":  pwd,
	})

}
