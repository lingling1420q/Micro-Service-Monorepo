package db

import (
	"database/sql"
	"fmt"
	config "gin-demo/config"
	"gin-demo/defs"
	logger "gin-demo/logger"
	"regexp"

	"github.com/go-redis/redis"

	_ "github.com/go-sql-driver/mysql" /* mysql driver init */
	"github.com/jinzhu/gorm"
	"reflect"
)

/* TODO: split dbs */
var (
	gormMysqlClient *gorm.DB
	sqlClient       *sql.DB
	redisClient     *redis.Client
	dbCfg           = config.Config().DB
	err             error
	mysqlCfg        = "Username:Password@tcp(Host:Port)/DbName?charset=utf8"
)

func init() {
	initCfg(dbCfg.Mysql)

}

func initGormMysql() {
	gormMysqlClient, _ = gorm.Open("mysql", mysqlCfg)
	defer gormMysqlClient.Close()
	//db.DB().SetMaxOpenConns(1)
	// gormMysqlClient.Exec("use gin")
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

	sqlClient, err = sql.Open("mysql", mysqlCfg)
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

func initCfg(cfgStruct interface{}) {
	t, v := reflect.TypeOf(cfgStruct), reflect.ValueOf(cfgStruct)
	for k := 0; k < t.NumField(); k++ {
		r, _ := regexp.Compile(t.Field(k).Name)
		mysqlCfg = r.ReplaceAllString(mysqlCfg, fmt.Sprintf(
			"%v", v.Field(k).Interface()))
	}
	logger.Debug(mysqlCfg)
}
