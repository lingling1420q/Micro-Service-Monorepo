package config

import "log"

var (
	// Config : Global Configuration
	Config config

	// switch env here
	isProduction bool = false
)

func init() {
	initConfiguration()
}

func initConfiguration() {
	if isProduction {
		Config = initProductionConfig()
		log.Println("Initialization Production Configuration ...")
	} else {
		Config = initDevelopmentConf()
		log.Println("Initialization Development Configuration ...")
	}
}

type config struct {
	Debug    bool
	DB       database
	Server   server
	Wechat   wechat
	QyWechat qyWechat
}

type database struct {
	Mysql mysql
	Redis redis
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

type wechat struct {
	WechatURL string
	Appid     string
	Secret    string
	GrantType string
}

type thirdParty struct {
	SuiteTokenURL string
	SuiteID       string
	SuitSecret    string
	TpURL         string
}

type qyWechat struct {
	WechatURL   string
	GetTokenURL string
	Corpid      string
	Secret      string
	GrantType   string
	ThirdParty  thirdParty
}
