package client

import (
	"monorepo/pkg/config"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
)

func NewRegistry() registry.Registry {
	if config.EnableKubernetes != "YES" {
		return etcdv3.NewRegistry(
			func(op *registry.Options) {
				op.Addrs = config.Etcd.Addrs
			},
		)
	}
	return kubernetes.NewRegistry()
}

func NewMicroClient() client.Client {
	return micro.NewService(
		micro.Registry(NewRegistry()),
	).Client()
}

func NewOriginalClient() client.Client {
	return micro.NewService().Client()
}
