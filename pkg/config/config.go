package config

import (
	"monorepo/pkg/global"
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
	Svc1 service
	Svc2 service
)

func init() {
	logger.I("System environment", "Production", Production, "EnableKubernetes", EnableKubernetes)
	switch Production {
	case global.YES:
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

type service struct {
	MicroName string
	Version   string
	Address   string
	Debug     bool
	DbURI     string
	DbDebug   bool
}
