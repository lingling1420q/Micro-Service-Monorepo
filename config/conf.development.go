package config

import "os"

func initDevelopmentConfig() config {

	return config{
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
			Debug: true,
		},
		Server: server{
			Debug:   true,
			Port:    ":4096",
			LogPath: "./tmp/logs/gin.log",
		},
		TmpFolder: os.TempDir(),
	}
}
