package user

import (
	"encoding/json"
	db "gin-demo/db"
	"gin-demo/defs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetUserInfo  get user info by ID
func GetUserInfo(c *gin.Context) {
	userID := c.Param("id")
	/* connection := db.ConnMysql()
	stmtOut, err := connection.Prepare("SELECT name,pwd FROM user WHERE id = ?")
	if err != nil {
		log.Printf("%s", err)
		return
	}

	var name string
	var pwd string
	err = stmtOut.QueryRow(userID).Scan(&name, &pwd)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	defer stmtOut.Close() */
	gormConn := db.ConnGormMysql()
	gormConn.SingularTable(true)
	user := defs.User{}
	// select from user by ID
	gormConn.Find(&user, userID)
	userInfo, _ := json.Marshal(user)
	//返回数据
	c.JSON(http.StatusOK, gin.H{
		"userInfo": string(userInfo),
	})
}
