package handler

import (
	"context"
	"monorepo/pkg/client"
	svc1 "monorepo/pkg/proto/svc1"
	"monorepo/pkg/proto/svc2"

	"github.com/monaco-io/logger"
)

type Svc1 struct{}

// CallSvc2 is a single request handler called via client.Call or the generated client code
func (e *Svc1) CallSvc2(ctx context.Context, req *svc1.CallSvc2Request, rsp *svc1.CallSvc2Response) error {
	logger.I("Received Svc1.CallSvc2 request", "req", req)
	in := svc2.Request{
		Name: req.Name,
	}
	resp, err := client.Svc2.Call(context.TODO(), &in, client.Svc2ClientOptions)
	// rsp.Msg = resp.Msg
	if err != nil {
		return err
	}
	rsp.Msg = resp.Msg
	return nil
}
