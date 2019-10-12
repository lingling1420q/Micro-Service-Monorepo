package config

var (
	/*Config : Global Configuration*/
	Config config
)

func init() {
	Config = config{
		DB: database{
			Mysql: mysql{
				Host:     "www.luxuze.club",
				Port:     "3306",
				Username: "root",
				Password: "pl,okm123",
				DbName:   "gin",
			},
			Redis: redis{
				Host:     "47.105.181.69",
				Port:     "6379",
				Password: "123456",
				DbName:   0,
			},
		},
		Server: server{
			Debug:   true,
			Port:    ":4096",
			LogPath: "./tmp/logs/gin.log",
		},
		Wechat: wechat{
			WechatURL: "https://api.weixin.qq.com/sns/jscode2session",
			Appid:     "wx4bd3f0044643ec2c",
			Secret:    "ae33804c97317d0cfc66239027d106ca",
			GrantType: "authorization_code",
		},
		QyWechat: qyWechat{
			WechatURL:   "https://qyapi.weixin.qq.com/cgi-bin/miniprogram/jscode2session",
			GetTokenURL: "https://qyapi.weixin.qq.com/cgi-bin/gettoken",
			Corpid:      "ww17f8d10783494584",
			Secret:      "i5t-rh8bXeNCgihcYPrG9ZPpWkivzPJ69sv570osk6I",
			GrantType:   "authorization_code",
			ThirdParty: thirdParty{
				SuiteTokenURL: "https://qyapi.weixin.qq.com/cgi-bin/service/get_suite_token",
				SuiteID:       "ww831f04b6b486f6d9",
				SuitSecret:    "JhY2CM6MQw216n7CzVJ68WfrqAS1RWLjDgRxtPyAznU",
				TpURL:         "https://qyapi.weixin.qq.com/cgi-bin/service/miniprogram/jscode2session",
			},
		},
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
