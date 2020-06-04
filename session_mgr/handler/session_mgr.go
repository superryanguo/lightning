package handler

import (
	"context"
	"encoding/json"

	log "github.com/micro/go-micro/v2/logger"
	"github.com/superryanguo/lightning/basic/cache"
	"github.com/superryanguo/lightning/models"
	"github.com/superryanguo/lightning/utils"

	session_mgr "github.com/superryanguo/lightning/session_mgr/proto/session_mgr"
)

type Session_mgr struct{}

func Init() {
	//TODO: do something here?
}

func (e *Session_mgr) GetSession(ctx context.Context, req *session_mgr.Request, rsp *session_mgr.Response) error {
	log.Info("GetSession url：api/v1.0/session")

	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	log.Info(req.SessionId)
	userInfo, err := cache.GetFromCache(req.SessionId)
	if err != nil {
		log.Info("No session data in cache")
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	user := models.User{}
	err = json.Unmarshal([]byte(userInfo), &user)
	if err != nil {
		log.Info("Data unmarshal json error")
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	rsp.Data = user.Name

	return nil
}
func (e *Session_mgr) SaveSession(ctx context.Context, ses *session_mgr.Session, rsp *session_mgr.Response) error {
	log.Info("SaveSession url：api/v1.0/session")
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Session_mgr) Stream(ctx context.Context, req *session_mgr.StreamingRequest, stream session_mgr.SessionMgr_StreamStream) error {
	log.Infof("Received Session_mgr.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&session_mgr.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Session_mgr) PingPong(ctx context.Context, stream session_mgr.SessionMgr_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&session_mgr.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
