package db

import (
	config "monorepo/config"
	"testing"
)

func TestInitCfg(t *testing.T) {
	dbCfg := config.Config.DB
	mysqlCfg := "Username:Password@tcp(Host:Port)/DbName?charset=utf8"
	InitCfg(&mysqlCfg, dbCfg.Mysql)
	t.Log(mysqlCfg)
}
