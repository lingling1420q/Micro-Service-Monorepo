package client

import (
	"monorepo/pkg/config"
	"monorepo/pkg/proto/svc2"
	"time"

	"github.com/micro/go-micro/v2/client"
)

var Svc2 = svc2.NewSvc2Service(
	config.Svc2.MicroName,
	NewOriginalClient())

var Svc2ClientOptions = func(o *client.CallOptions) {
	o.Retries = 2
	o.RequestTimeout = time.Second * 5
	o.DialTimeout = time.Second * 5
	o.StreamTimeout = time.Second * 5
}
