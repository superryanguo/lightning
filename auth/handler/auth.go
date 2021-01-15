package handler

import (
	"context"
	"strconv"

	log "github.com/micro/go-micro/v2/logger"

	"github.com/superryanguo/lightning/auth/jwtoken"
	auth "github.com/superryanguo/lightning/auth/proto/auth"
)

type Auth struct{}

func Init() {
}

func (e *Auth) MakeAccessToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Debug("[MakeAccessToken] receive the create req")

	token, err := jwtoken.MakeAccessToken(&jwtoken.Subject{
		ID:   strconv.FormatInt(req.UserId, 10),
		Name: req.UserName,
	})
	if err != nil {
		rsp.Error = &auth.Error{
			Detail: err.Error(),
		}

		log.Errorf("[MakeAccessToken] fail err：%s", err)
		return err
	}

	rsp.Token = token
	return nil
}
func (e *Auth) AuthAccessToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Debug("[AuthAccessToken] receive the auth req")

	return nil
}
func (s *Auth) DelUserAccessToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Debug("[DelUserAccessToken]...")
	err := jwtoken.DelUserAccessToken(req.Token)
	if err != nil {
		rsp.Error = &auth.Error{
			Detail: err.Error(),
		}

		log.Errorf("[DelUserAccessToken] fail err：%s", err)
		return err
	}
	return nil
}

func (s *Auth) GetCachedAccessToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Debugf("[GetCachedAccessToken] token:%d", req.UserId)
	token, err := jwtoken.GetCachedAccessToken(&jwtoken.Subject{
		ID: strconv.FormatInt(req.UserId, 10),
	})
	if err != nil {
		rsp.Error = &auth.Error{
			Detail: err.Error(),
		}

		log.Errorf("[GetCachedAccessToken] fail err：%s", err)
		return err
	}

	rsp.Token = token
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Auth) Stream(ctx context.Context, req *auth.StreamingRequest, stream auth.Auth_StreamStream) error {
	log.Infof("Received Auth.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&auth.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Auth) PingPong(ctx context.Context, stream auth.Auth_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&auth.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
