package db

import (
	"database/sql"
	"fmt"
	config "gin-demo/config"
	"gin-demo/defs"
	logger "gin-demo/logger"
	"strings"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	gormMysqlClient *gorm.DB
	sqlClient       *sql.DB
	redisClient     *redis.Client
	dbCfg           = config.Config().DB
	err             error
	mysqlPath       = strings.Join([]string{
		dbCfg.Mysql.Username, ":",
		dbCfg.Mysql.Password, "@tcp(",
		dbCfg.Mysql.Host, ":",
		dbCfg.Mysql.Port, ")/",
		dbCfg.Mysql.DbName, "?charset=utf8"}, "")
)

func initGormMysql() {

	gormMysqlClient, _ = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/?charset=utf8")
	defer gormMysqlClient.Close()
	//db.DB().SetMaxOpenConns(1)
	gormMysqlClient.Exec("use gin")
	if err != nil {
		logger.Error(defs.ConnDBErr, err.Error())
	}
}

// ConnGormMysql gormMysql
func ConnGormMysql() *gorm.DB {
	if gormMysqlClient == nil {
		initGormMysql()
	}
	return gormMysqlClient
}

func initConnMysql() {
	logger.Info("mysql database connection initialization ...")

	sqlClient, err = sql.Open("mysql", mysqlPath)
	if err != nil {
		logger.Error(defs.ConnDBErr, err.Error())
		panic(err.Error())
	}
}

// ConnMysql  connect to mysql
func ConnMysql() *sql.DB {
	if sqlClient == nil || sqlClient.Ping() != nil {
		initConnMysql()
	}
	return sqlClient
}

func initRedis() {
	logger.Info("redis database connection initialization ...")
	client := redis.NewClient(&redis.Options{
		Addr:     dbCfg.Redis.Host + ":" + dbCfg.Redis.Port,
		Password: dbCfg.Redis.Password,
		DB:       dbCfg.Redis.DbName,
	})
	fmt.Println(client.Ping().Result())
}

// ConnRedis connect to redis
func ConnRedis() *redis.Client {
	if redisClient == nil {
		initRedis()
	}
	return redisClient
}
