package subscriber

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	user_srv "github.com/superryanguo/lightning/user_srv/proto/user_srv"
)

type User_srv struct{}

func (e *User_srv) Handle(ctx context.Context, msg *user_srv.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *user_srv.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
