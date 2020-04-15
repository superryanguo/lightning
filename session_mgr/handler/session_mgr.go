package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	session_mgr "github.com/superryanguo/lightning/session_mgr/proto/session_mgr"
)

type Session_mgr struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Session_mgr) Call(ctx context.Context, req *session_mgr.Request, rsp *session_mgr.Response) error {
	log.Log("Received Session_mgr.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Session_mgr) Stream(ctx context.Context, req *session_mgr.StreamingRequest, stream session_mgr.Session_mgr_StreamStream) error {
	log.Logf("Received Session_mgr.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&session_mgr.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Session_mgr) PingPong(ctx context.Context, stream session_mgr.Session_mgr_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&session_mgr.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
