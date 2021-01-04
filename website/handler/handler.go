package handler

import (
	"context"
	"encoding/json"
	"image"
	"image/png"
	"net/http"
	"reflect"

	"github.com/afocus/captcha"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	auth "github.com/superryanguo/lightning/auth/proto/auth"
	"github.com/superryanguo/lightning/models"
	sm "github.com/superryanguo/lightning/session_mgr/proto/session_mgr"
	user "github.com/superryanguo/lightning/user_srv/proto/user_srv"
	"github.com/superryanguo/lightning/utils"
)

var (
	smClient   sm.SessionMgrService
	userClient user.UserSrvService
	authClient auth.Service
)

func Init() {
	smClient = sm.NewSessionMgrService("micro.super.lightning.service.session_mgr", client.DefaultClient)
	userClient = user.NewUserSrvService("micro.super.lightning.service.user_srv", client.DefaultClient)
	authClient = auth.NewService("micro.super.lightning.service.auth", client.DefaultClient)
}

func GetImageCd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Info("GetImageCd-> url:api/v1.0/imagecode/:uuid")

	// call the backend service
	rsp, err := userClient.GetImageCd(context.TODO(), &user.ImageRequest{
		Uuid: ps.ByName("uuid"),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var img image.RGBA
	img.Stride = int(rsp.Stride)
	img.Rect.Min.X = int(rsp.Min.X)
	img.Rect.Min.Y = int(rsp.Min.Y)
	img.Rect.Max.X = int(rsp.Max.X)
	img.Rect.Max.Y = int(rsp.Max.Y)
	img.Pix = []uint8(rsp.Pix)

	var image captcha.Image
	image.RGBA = &img

	log.Debug("GetImageCd->send the image to webpage")
	png.Encode(w, image)

}

func PostUserAuth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Info(" PutUserRealName->  api/v1.0/user/infoauth ")

	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	userlogin, err := r.Cookie("userlogin")
	if err != nil {
		resp := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			log.Debug(err)
			return
		}
		return
	}

	rsp, err := userClient.PostUserReal(context.TODO(), &user.RealNameRequest{
		SessionId: userlogin.Value,
		RealName:  request["real_name"].(string),
		IdCard:    request["id_card"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
}

func PutUserInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Info("PutUserInfo-> url：api/v1.0/user/name")

	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	userlogin, err := r.Cookie("userlogin")
	if err != nil {
		resp := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			log.Info(err)
			return
		}
		return
	}
	rsp, err := userClient.PutUserInfo(context.TODO(), &user.PutRequest{
		SessionId: userlogin.Value,
		Username:  request["name"].(string),
	})

	if err != nil {
		http.Error(w, err.Error(), 502)
		return
	}
	data := make(map[string]interface{})
	data["name"] = rsp.Username
	resp := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), 503)
		log.Info(err)
		return
	}
	return
}

func PostAvatar(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Info("PostAvatar->  /api/v1.0/user/avatar")

	userlogin, err := r.Cookie("userlogin")
	if err != nil || "" == userlogin.Value {
		resp := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			return
		}
		return
	}
	file, handler, err := r.FormFile("avatar")
	if err != nil {
		log.Debug("Avatar file error")
		resp := map[string]interface{}{
			"errno":  utils.RECODE_IOERR,
			"errmsg": utils.RecodeText(utils.RECODE_IOERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			return
		}
		return
	}

	log.Info("FileSize:", handler.Size)
	log.Info("FileName:", handler.Filename)

	filebuffer := make([]byte, handler.Size)
	_, err = file.Read(filebuffer)
	if err != nil {
		log.Debug("AvatarFile receive error")
		resp := map[string]interface{}{
			"errno":  utils.RECODE_IOERR,
			"errmsg": utils.RecodeText(utils.RECODE_IOERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			return
		}
		return
	}

	rsp, err := userClient.PostAvatar(context.TODO(), &user.AvaRequest{
		Avatar:    filebuffer,
		SessionId: userlogin.Value,
		Filename:  handler.Filename,
		Filesize:  handler.Size,
	})

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := make(map[string]interface{})
	data["avatar_url"] = utils.AddDomain2Url(rsp.AvatarUrl)
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}
	w.Header().Set("Content-type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
func GetUserInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Info("GetUserInfo-> url：api/v1.0/user/Or/infoauth")

	userlogin, err := r.Cookie("userlogin")
	if err != nil {
		resp := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			log.Info(err)
			return
		}
		return
	}
	rsp, err := userClient.GetUserInfo(context.TODO(), &user.UserInfoRequest{
		SessionId: userlogin.Value,
	})

	if err != nil {
		http.Error(w, err.Error(), 502)
		return
	}

	data := make(map[string]interface{})

	data["user_id"] = rsp.UserId
	data["name"] = rsp.Name
	data["email"] = rsp.Email
	data["real_name"] = rsp.RealName
	data["id_card"] = rsp.IdCard
	data["avatar_url"] = utils.AddDomain2Url(rsp.AvatarUrl)
	resp := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), 503)
		log.Info(err)
		return
	}
	return
}

func GetArea(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Info("GetArea-> url:api/v1.0/lightning/areas")

	rsp, err := userClient.GetArea(context.TODO(), &user.AreaRequest{})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	area_list := []models.Area{}
	for _, value := range rsp.Data {
		tmp := models.Area{ID: int(value.Aid), Name: value.Aname}
		area_list = append(area_list, tmp)
	}

	log.Debug("area_list:", area_list)
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   area_list,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func GetEmailCd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Info("GetEmailCd-> url:api/v1.0/emailcode/:email")

	email := ps.ByName("email")
	log.Info("GetEmailCd->", email)

	log.Info("GetEmailCd->", r.URL.Query())
	text := r.URL.Query()["text"][0]
	id := r.URL.Query()["id"][0]

	rsp, err := userClient.GetEmailCd(context.TODO(), &user.MailRequest{
		Email: email,
		Uuid:  id,
		Text:  text,
	})

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}

	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func GetIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Info("GetIndex-> html show api/v1.0/lightning/house/index")

	response := map[string]interface{}{
		"errno":  utils.RECODE_OK,
		"errmsg": utils.RecodeText(utils.RECODE_OK),
	}
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func GetSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Info("GetSession-> Retrieve the session url：api/v1.0/session")

	userlogin, err := r.Cookie("userlogin")
	if err != nil || "" == userlogin.Value {
		log.Debug("GetSession-> no login info found...")
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	//如果cookie有值就发送到服务端
	rsp, err := smClient.GetSession(context.TODO(), &sm.Request{
		SessionId: userlogin.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := make(map[string]string)
	u := models.User{}
	err = json.Unmarshal([]byte(rsp.Data), &u)
	if err != nil {
		log.Info("Data unmarshal json error")
		http.Error(w, err.Error(), 500)
		return
	}
	data["name"] = u.Name
	//创建返回数据map
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func DeleteSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Info("DeleteSession-> Delete the session /api/v1.0/session")

	userlogin, err := r.Cookie("userlogin")
	if err != nil {
		resp := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			log.Info(err)
			return
		}
		return
	}

	rsp, err := smClient.DeleteSession(context.TODO(), &sm.Request{
		SessionId: userlogin.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	////读取cookie
	//cookie, err := r.Cookie("userlogin")
	////如果读取失败或者cookie中的value不存在则创建cookie
	//if err != nil || "" == cookie.Value {
	//return
	//} else {
	//cookie := http.Cookie{Name: "userlogin", Path: "/", MaxAge: 600}
	//http.SetCookie(w, &cookie)
	//}

	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}
	w.Header().Set("Content-type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
func PostReg(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Info("PostReg-> /api/v1.0/users")

	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	for key, value := range request {
		if key != "password" {
			log.Debug(key, ":", value, "[", reflect.TypeOf(value), "]")
		} else {
			log.Debug(key, ":", "******", "[", reflect.TypeOf(value), "]")
		}
	}

	if request["email"] == "" || request["password"] == "" || request["email_code"] == "" {
		resp := map[string]interface{}{
			"errno":  utils.RECODE_NODATA,
			"errmsg": "empty input found",
		}
		w.Header().Set("Content-type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			log.Info(err)
			return
		}
		log.Info("PostReg-> empty email password or emailcode")
		return
	}

	// call the backend service
	rsp, err := userClient.PostReg(context.TODO(), &user.Request{
		Email:     request["email"].(string),
		Password:  request["password"].(string),
		EmailCode: request["email_code"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}

	//读取cookie
	cookie, err := r.Cookie("userlogin")
	//如果读取失败或者cookie中的value不存在则创建cookie
	if err != nil || "" == cookie.Value {
		cookie := http.Cookie{Name: "userlogin", Value: rsp.SessionId, Path: "/", MaxAge: 600}
		http.SetCookie(w, &cookie)
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
func PostLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Info("PostLogin-> to  /api/v1.0/sessions")

	//decode the user input
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	for key, value := range request {
		log.Info(key, "-", value, "-", reflect.TypeOf(value))
	}

	if request["email"] == "" || request["password"] == "" {
		resp := map[string]interface{}{
			"errno":  utils.RECODE_NODATA,
			"errmsg": "empty email or password",
		}
		w.Header().Set("Content-type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			log.Info(err)
			return
		}
		log.Info("PostLogin-> empty email or password")
		return
	}

	// call the backend service
	log.Debugf("PostLogin-> call the postlogin client=%v", userClient)
	rsp, err := userClient.PostLogin(context.TODO(), &user.Request{
		Email:    request["email"].(string),
		Password: request["password"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//读取cookie
	cookie, err := r.Cookie("userlogin")
	//如果读取失败或者cookie中的value不存在则创建cookie
	if err != nil || "" == cookie.Value {
		cookie := http.Cookie{Name: "userlogin", Value: rsp.SessionId, Path: "/", MaxAge: 600}
		http.SetCookie(w, &cookie)
	}

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}
	w.Header().Set("Content-type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
