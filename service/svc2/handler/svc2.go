package handler

import (
	"context"

	"github.com/monaco-io/logger"

	svc2 "monorepo/pkg/proto/svc2"
)

type Svc2 struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Svc2) Call(ctx context.Context, req *svc2.Request, rsp *svc2.Response) error {
	logger.I("Received Svc1.Call request", "req", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Svc2) Stream(ctx context.Context, req *svc2.StreamingRequest, stream svc2.Svc2_StreamStream) error {
	logger.I("Received Svc1.Stream request", "count", req)
	for i := 0; i < int(req.Count); i++ {
		logger.I("Responding", "i", i)
		if err := stream.Send(&svc2.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Svc2) PingPong(ctx context.Context, stream svc2.Svc2_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		logger.I("Got ping", "req.Stroke", req.Stroke)
		if err := stream.Send(&svc2.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
