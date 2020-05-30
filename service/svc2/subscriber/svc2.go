package subscriber

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	svc2 "monorepo/pkg/proto/svc2"
)

type Svc2 struct{}

func (e *Svc2) Handle(ctx context.Context, msg *svc2.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *svc2.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
