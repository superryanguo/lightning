package subscriber

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	orders "github.com/superryanguo/lightning/orders/proto/orders"
)

type Orders struct{}

func (e *Orders) Handle(ctx context.Context, msg *orders.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *orders.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
