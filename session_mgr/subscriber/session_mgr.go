package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	session_mgr "github.com/superryanguo/lightning/session_mgr/proto/session_mgr"
)

type Session_mgr struct{}

func (e *Session_mgr) Handle(ctx context.Context, msg *session_mgr.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *session_mgr.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}