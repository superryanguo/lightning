package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	user_srv "github.com/superryanguo/lightning/user_srv/proto/user_srv"
)

type User_srv struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *User_srv) Call(ctx context.Context, req *user_srv.Request, rsp *user_srv.Response) error {
	log.Info("Received User_srv.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *User_srv) Stream(ctx context.Context, req *user_srv.StreamingRequest, stream user_srv.UserSrv_StreamStream) error {
	log.Infof("Received User_srv.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&user_srv.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *User_srv) PingPong(ctx context.Context, stream user_srv.UserSrv_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&user_srv.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
