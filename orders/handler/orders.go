package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	orders "github.com/superryanguo/lightning/orders/proto/orders"
)

type Orders struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Orders) Call(ctx context.Context, req *orders.Request, rsp *orders.Response) error {
	log.Info("Received Orders.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Orders) Stream(ctx context.Context, req *orders.StreamingRequest, stream orders.Orders_StreamStream) error {
	log.Infof("Received Orders.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&orders.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Orders) PingPong(ctx context.Context, stream orders.Orders_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&orders.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
