package config

import (
	"log"
	"os"
)

var (
	// Config : Global Configuration
	Config config

	// switch env here
	appENV string = os.Getenv("APP_ENV")
)

func init() {
	initConfiguration()
}

func initConfiguration() {
	if appENV == "PRODUCTION" {
		Config = initProductionConfig()
		log.Println("Initialization Production Configuration ...")
	} else if appENV == "STAGING" {
		Config = initStagingConfig()
		log.Println("Initialization Staging Configuration ...")
	} else {
		Config = initDevelopmentConfig()
		log.Println("Initialization Development Configuration ...")
	}
}

type config struct {
	Debug     bool
	DB        database
	Server    server
	TmpFolder string
}

type database struct {
	Mysql mysql
	Redis redis
	Debug bool
}

type mysql struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
}

type redis struct {
	Host     string
	Port     string
	Password string
	DbName   int
}

type server struct {
	Debug   bool
	Port    string
	LogPath string
}
