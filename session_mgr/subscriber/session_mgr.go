package subscriber

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	session_mgr "github.com/superryanguo/lightning/session_mgr/proto/session_mgr"
)

type Session_mgr struct{}

func (e *Session_mgr) Handle(ctx context.Context, msg *session_mgr.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *session_mgr.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
