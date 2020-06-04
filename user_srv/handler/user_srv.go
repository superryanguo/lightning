package handler

import (
	"context"
	"encoding/json"

	"github.com/astaxie/beego/orm"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/superryanguo/lightning/utils"

	"github.com/superryanguo/lightning/basic/cache"
	"github.com/superryanguo/lightning/models"
	user_srv "github.com/superryanguo/lightning/user_srv/proto/user_srv"
)

type User_srv struct{}

func (e *User_srv) PostLogin(ctx context.Context, req *user_srv.Request, rsp *user_srv.Response) error {
	log.Info("ServicePostLogin  /api/v1.0/sessions")

	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)
	//database query
	var user models.User
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	err := qs.Filter("email", req.Email).One(&user)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	//compare the password
	pwd_hash := utils.Sha256Encode(req.Password)
	if pwd_hash != user.Password_hash {
		rsp.Errno = utils.RECODE_PWDERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//bm, err := utils.GetRedisConnector()
	//if err != nil {
	//log.Debug("redis connection failure in postlogin", err)
	//rsp.Errno = utils.RECODE_DBERR
	//rsp.Errmsg = utils.RecodeText(rsp.Errno)
	//return nil
	//}

	//TODO: should put this part into the session mgr
	sessionId := utils.Sha256Encode(pwd_hash)
	rsp.SessionId = sessionId
	user.Password_hash = ""
	userInfo, _ := json.Marshal(user)
	//bm.Put(sessionId, userInfo, time.Second*600)

	//ca := redis.GetRedis()
	err = cache.SaveToCache(sessionId, userInfo)
	if err != nil {
		log.Debug("redis save sessionid failure in postlogin", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

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
