package db

import (
	_ "github.com/go-sql-driver/mysql" /* mysql driver init */
	"github.com/jinzhu/gorm"
	"monaco/config"
	"monaco/logger"
)

var (
	sqlClient *gorm.DB
	dbCfg     = config.Config.DB
	err       error
	mysqlCfg  = "Username:Password@tcp(Host:Port)/DbName?charset=utf8mb4&parseTime=True&loc=Local"
)

func init() {
	InitCfg(&mysqlCfg, dbCfg.Mysql)
	logger.Noticef("Mysql config: %s", mysqlCfg)
}

// DB gormMysql 数据库链接
func DB() *gorm.DB {
	if sqlClient == nil || sqlClient.DB().Ping() != nil {
		sqlClient, err = gorm.Open("mysql", mysqlCfg)
		sqlClient.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
			&TBL_USERS{},
			&TBL_VISIT_LOG{},
		)
		sqlClient.LogMode(dbCfg.Debug)
		if err != nil {
			logger.Error("defs.ConnDBErr", err.Error())
		}
	}
	return sqlClient
}
