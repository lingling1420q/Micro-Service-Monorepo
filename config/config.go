package config

import (
	"log"
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
)

var (
	cfg  *TomlConfig
	once sync.Once
)

/*TomlConfig ... */
type TomlConfig struct {
	Debug   bool
	DB      database `toml:"database"`
	Server  server
	Clients clients
}

type database struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
}

type server struct {
	Port string
}

type clients struct {
	Data  [][]interface{}
	Hosts []string
}

/*Config ... */
func Config() *TomlConfig {
	once.Do(func() {
		filePath, err := filepath.Abs("./config/config.toml")
		if err != nil {
			panic(err)
		}
		log.Printf("parse toml file singleton. filePath: %s\n", filePath)
		if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
			panic(err)
		}
	})
	return cfg
}
