package handler

import (
	"context"
	"encoding/json"

	"github.com/superryanguo/lightning/models"

	"github.com/garyburd/redigo/redis"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/superryanguo/lightning/utils"

	session_mgr "github.com/superryanguo/lightning/session_mgr/proto/session_mgr"
)

type Session_mgr struct{}

func (e *Session_mgr) GetSession(ctx context.Context, req *session_mgr.Request, rsp *session_mgr.Response) error {
	log.Info("获取Session url：api/v1.0/session")

	rsp.Errno = utils.RECODE_OK

	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	bm, err := utils.GetRedisConnector()
	if err != nil {
		log.Info("获取缓存连接失败")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	//从缓存中拿到用户信息
	log.Info(req.SessionId)
	userInfo_redis := bm.Get(req.SessionId)
	userInfo_string, _ := redis.String(userInfo_redis, nil)
	log.Info(userInfo_string)
	userInfo := []byte(userInfo_string)
	user := models.User{}
	err = json.Unmarshal(userInfo, &user)
	if err != nil {
		log.Info("Json解析异常")
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	rsp.Data = user.Name

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
