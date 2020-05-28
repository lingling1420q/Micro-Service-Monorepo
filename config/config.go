package config

import (
	"os"

	"github.com/monaco-io/logger"
)

// init config
var (
	Production       = os.Getenv("PRODUCTION")
	EnableKubernetes = os.Getenv("ENABLE_KUBERNETES")
)

var (
	mock interface{}
)

// service config
var (
	Etcd etcd
	Srv1 srv
)

func init() {
	logger.I("System environment", "Production", Production, "EnableKubernetes", EnableKubernetes)
	switch Production {
	case "YES":
		production()
	default:
		staging()
		switch mock.(type) {
		case func():
			mock.(func())()
		}
	}
}

type etcd struct {
	Addrs []string
}

type srv struct {
	MicroName string
	Version   string
	Address   string
	Debug     bool
	DbURI     string
	DbDebug   bool
}
