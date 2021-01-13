package subscriber

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	auth "github.com/superryanguo/lightning/auth/proto/auth"
)

type Auth struct{}

func (e *Auth) Handle(ctx context.Context, msg *auth.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *auth.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
