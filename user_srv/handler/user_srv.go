package handler

import (
	"context"
	"encoding/json"
	"image/color"
	"math/rand"
	"path"
	"strconv"
	"time"

	"github.com/afocus/captcha"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/superryanguo/lightning/basic/cache"
	"github.com/superryanguo/lightning/models"
	sm "github.com/superryanguo/lightning/session_mgr/proto/session_mgr"
	user_srv "github.com/superryanguo/lightning/user_srv/proto/user_srv"
	"github.com/superryanguo/lightning/utils"
)

type User_srv struct{}

var (
	smClient sm.SessionMgrService
)

func Init() {
	smClient = sm.NewSessionMgrService("micro.super.lightning.service.session_mgr", client.DefaultClient)
}

func (e *User_srv) PutUserInfo(ctx context.Context, req *user_srv.PutRequest, rsp *user_srv.PutResponse) error {
	log.Info("PutUserInfo->  url：api/v1.0/user/name")

	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	userInfo_redis, err := cache.GetFromCache(req.SessionId)
	if err != nil {
		log.Debug("PutUserInfo->cache problem or empty data:", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	userOld := models.User{}
	json.Unmarshal(userInfo_redis, &userOld)
	user := models.User{Uid: userOld.Uid, Name: req.Username}

	db := models.GetGorm()
	err = db.Debug().Model(&user).Update("name", req.Username).Error
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	userOld.Name = req.Username
	userInfo, _ := json.Marshal(userOld)
	err = cache.SaveToCache(req.SessionId, userInfo)
	if err != nil {
		log.Debug("PutUserInfo->cache update username failure", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	rsp.Username = user.Name
	return nil
}
func (e *User_srv) PostUserReal(ctx context.Context, req *user_srv.RealNameRequest, rsp *user_srv.Response) error {
	log.Info(" PostUserReal->  api/v1.0/user/auth ")

	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	userInfo_redis, err := cache.GetFromCache(req.SessionId)
	if err != nil {
		log.Debug("PostUserReal->cache problem or empty data:", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	userOld := models.User{}
	err = json.Unmarshal(userInfo_redis, &userOld)
	if err != nil {
		rsp.Errno = utils.RECODE_SERVERERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	user := models.User{Uid: userOld.Uid}

	db := models.GetGorm()

	err = db.Debug().Model(&user).Updates(models.User{Real_name: req.RealName, Id_card: req.IdCard}).Error
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	userOld.Real_name = req.RealName
	userOld.Id_card = req.IdCard
	userInfo, _ := json.Marshal(userOld)
	err = cache.SaveToCache(req.SessionId, userInfo)
	if err != nil {
		log.Debug("PostUserReal->cache update failure", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	return nil
}
func (e *User_srv) PostAvatar(ctx context.Context, req *user_srv.AvaRequest, rsp *user_srv.AvaResponse) error {
	log.Info("PostAvatar->  /api/v1.0/avatar")

	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	log.Debug("Avatar:", len(req.Avatar), req.Filesize)
	fileExt := path.Ext(req.Filename)
	filename, err := utils.UploadByBuffer(req.Avatar, fileExt[1:])
	if err != nil {
		log.Debug("Errors when uploading to server")
		rsp.Errno = utils.RECODE_IOERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	userInfo_redis, err := cache.GetFromCache(req.SessionId)
	if err != nil && err != redis.Nil {
		log.Debug("PostAvatar->cache problem or no data in cache:", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	userOld := models.User{}

	err = json.Unmarshal(userInfo_redis, &userOld)
	if err != nil {
		rsp.Errno = utils.RECODE_SERVERERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	user := models.User{Uid: userOld.Uid, Avatar_url: filename}
	db := models.GetGorm()
	err = db.Debug().Model(&user).Update("avatar_url", filename).Error
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	userOld.Avatar_url = filename
	userInfo, _ := json.Marshal(userOld)

	err = cache.SaveToCache(req.SessionId, userInfo)
	if err != nil {
		log.Debug("PostAvatar->cache update failure", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	rsp.AvatarUrl = filename
	return nil
}

func (e *User_srv) GetUserInfo(ctx context.Context, req *user_srv.UserInfoRequest, rsp *user_srv.UserInfoResponse) error {
	log.Info("GetUserInfo-> url：api/v1.0/user/Or/infoauth")

	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	userInfo_redis, err := cache.GetFromCache(req.SessionId)
	if err != nil && err != redis.Nil {
		log.Debug("GetArea->cache problem or no data session expired:", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	user := models.User{}
	err = json.Unmarshal(userInfo_redis, &user)
	if err != nil {
		rsp.Errno = utils.RECODE_SERVERERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	rsp.UserId = user.Uid
	rsp.Email = user.Email
	rsp.Name = user.Name
	rsp.RealName = user.Real_name
	rsp.IdCard = user.Id_card
	rsp.AvatarUrl = user.Avatar_url
	return nil
}
func (e *User_srv) GetArea(ctx context.Context, req *user_srv.AreaRequest, rsp *user_srv.AreaResponse) error {
	log.Info("GetArea-> url:/api/v1.0/lightning/areas")
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	ca, err := cache.GetFromCache("areas_info")
	if err != nil && err != redis.Nil {
		log.Debug("GetArea->cache problem:", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	if ca != nil {
		areas_info := []map[string]interface{}{}
		err = json.Unmarshal(ca, &areas_info)
		for _, value := range areas_info {
			area := user_srv.AreaResponse_Address{Aid: int32(value["area_id"].(float64)), Aname: value["area_name"].(string)}
			rsp.Data = append(rsp.Data, &area)
		}
		log.Debug("GetArea->Areas:", rsp.Data)
		return nil
	}

	var areas []models.Area
	db := models.GetGorm()
	err = db.Debug().Find(&areas).Error
	if err != nil {
		log.Info("GetArea can't find area data in DB, Err:", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	if len(areas) == 0 {
		rsp.Errno = utils.RECODE_NODATA
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	area_json, _ := json.Marshal(areas)
	err = cache.SaveToCache("areas_info", area_json)
	if err != nil {
		log.Debug("areas data can't save to cache", err)
		//rsp.Errno = utils.RECODE_DBERR
		//rsp.Errmsg = utils.RecodeText(rsp.Errno)
		//return nil
		err = nil //TODO:this is not critical error, we can try again next time
	}
	for _, value := range areas {
		area := user_srv.AreaResponse_Address{Aid: int32(value.ID), Aname: value.Name}
		rsp.Data = append(rsp.Data, &area)
	}
	return nil
}
func (e *User_srv) GetImageCd(ctx context.Context, req *user_srv.ImageRequest, rsp *user_srv.ImageResponse) error {
	log.Info("GetImageCd-> url:/api/v1.0/imagecode/:uuid=", req.Uuid)
	cap := captcha.New()
	if err := cap.SetFont("comic.ttf"); err != nil {
		log.Info("GetImageCd->No font file")
		panic(err.Error())
	}
	//设置图片大小
	cap.SetSize(91, 41)
	//设置干扰强度
	cap.SetDisturbance(captcha.MEDIUM)
	// 设置前景色 可以多个 随机替换文字颜色 默认黑色
	//SetFrontColor(colors ...color.Color)  这两个颜色设置的函数属于不定参函数
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	// 设置背景色 可以多个 随机替换背景色 默认白色
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255},
		color.RGBA{0, 153, 0, 255})
	//生成图片 返回图片和 字符串(图片内容的文本形式)
	img, str := cap.Create(4, captcha.NUM)
	b := *img      //解引用
	c := *(b.RGBA) //解引用
	//默认返回成功
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)
	//图片信息
	rsp.Pix = []byte(c.Pix)
	rsp.Stride = int64(c.Stride)
	rsp.Max = &user_srv.ImageResponse_Point{X: int64(c.Rect.Max.X), Y: int64(c.Rect.Max.Y)}
	rsp.Min = &user_srv.ImageResponse_Point{X: int64(c.Rect.Min.X), Y: int64(c.Rect.Min.Y)}

	//将uuid与验证码存入redis
	//bm.Put(req.Uuid, str, time.Second*300)
	err := cache.SaveToCache(req.Uuid, []byte(str))
	if err != nil {
		log.Debug("GetImageCd->redis save req.uuid failure", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	return nil
}

func (e *User_srv) GetEmailCd(ctx context.Context, req *user_srv.MailRequest, rsp *user_srv.Response) error {
	log.Info("GetEmailCd-> url:api/v1.0/emailcode/:email")
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	user := models.User{}
	db := models.GetGorm()
	err := db.Debug().Where(&models.User{Email: req.Email}).First(&user).Error
	if err == nil {
		log.Debug("GetEmailCd->user already exist, Err:", err)
		rsp.Errno = utils.RECODE_DATAEXIST
		rsp.Errmsg = utils.RecodeText(rsp.Errmsg)
		return nil
	}

	value, err := cache.GetFromCache(req.Uuid)
	if err != nil || value == nil {
		log.Debug("GetEmailCd->Cache query failure")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	if req.Text != string(value) {
		log.Debug("GetEmailCd->code mismatch")
		rsp.Errno = utils.RECODE_SMSERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code_number := r.Intn(9999) + 1001
	log.Info("GetEmailCd->code_number=", code_number)
	code := strconv.Itoa(code_number)
	//发送邮箱验证码
	err = utils.SendEmail(req.Email, code)
	if err != nil {
		log.Debug("GetEmailCd->fail to send mail")
		rsp.Errno = utils.RECODE_SERVERERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	err = cache.SaveToCache(req.Email, []byte(code))
	if err != nil {
		log.Debug("GetEmailCd->cache save reqEmailCode failure", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	return nil
}
func (e *User_srv) PostReg(ctx context.Context, req *user_srv.Request, rsp *user_srv.Response) error {
	log.Info("PostReg->  /api/v1.0/users")

	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	code_redis, err := cache.GetFromCache(req.Email)
	if err != nil || code_redis == nil {
		log.Debug("PostReg->empty email code data in cache")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		//err = nil
		return nil
	}
	if req.EmailCode != string(code_redis) {
		log.Debug("PostReg->wrong email code")
		rsp.Errno = utils.RECODE_SMSERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		//TODO: can skip the check now, need open when lauch the whole program
		return nil
	}
	user := models.User{}
	user.Uid = uuid.New().String()
	user.Name = req.Email
	pwd_hash := utils.Sha256Encode(req.Password)
	user.Password_hash = pwd_hash
	user.Email = req.Email
	log.Info("PostReg->generate the register user uid", user.Uid)
	db := models.GetGorm()
	err = db.Debug().Create(&user).Error
	if err != nil {
		log.Debug("PostReg->fail to insert a user to db", user)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	sessionId := utils.Sha256Encode(user.Password_hash)
	rsp.SessionId = sessionId
	user.Password_hash = ""
	userInfo, _ := json.Marshal(user)

	rp, err := smClient.SaveSession(context.TODO(), &sm.Session{
		SessionId:   sessionId,
		SessionData: userInfo,
	})

	if err != nil {
		log.Debug("PostReg->smgr->redis save sessionid failure", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	rsp.Errno = rp.Errno
	rsp.Errmsg = rp.Errmsg

	return nil
}

func (e *User_srv) PostLogin(ctx context.Context, req *user_srv.Request, rsp *user_srv.Response) error {
	log.Info("PostLogin-> /api/v1.0/sessions")

	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)
	//database query
	var user models.User
	db := models.GetGorm()
	if req.Email == "" { //can't be empty or the gorm won't check the db
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	err := db.Debug().Where(&models.User{Email: req.Email}).First(&user).Error
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

	sessionId := utils.Sha256Encode(pwd_hash)
	rsp.SessionId = sessionId
	user.Password_hash = ""
	userInfo, _ := json.Marshal(user)

	rp, err := smClient.SaveSession(context.TODO(), &sm.Session{
		SessionId:   sessionId,
		SessionData: userInfo,
	})

	if err != nil {
		log.Debug("PostLogin->smgr->redis save sessionid failure", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	rsp.Errno = rp.Errno
	rsp.Errmsg = rp.Errmsg

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
