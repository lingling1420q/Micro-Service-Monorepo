package db

import (
	"database/sql"
	config "gin-demo/config"
	logger "gin-demo/logger"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var (
	connection *sql.DB
	err        error
)

func initConn() {
	logger.Info("database connection initialization ...")
	dbCfg := config.Config().DB
	path := strings.Join([]string{dbCfg.Username, ":", dbCfg.Password, "@tcp(", dbCfg.Host, ":", dbCfg.Port, ")/", dbCfg.DbName, "?charset=utf8"}, "")
	connection, err = sql.Open("mysql", path)
	if err != nil {
		logger.Error(err.Error())
		panic(err.Error())
	}
}

/*Conn  get db connection*/
func Conn() *sql.DB {
	if connection == nil || connection.Ping() != nil {
		initConn()
	}
	return connection
}
