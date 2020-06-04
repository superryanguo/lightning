package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	auth "github.com/superryanguo/lightning/auth/proto/auth"
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

func GetIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Info("Get the Index html show：api/v1.0/lightning/index")

	//创建回数据map
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
	log.Info("Retrieve the session url：api/v1.0/session")

	//获取cookie
	userlogin, err := r.Cookie("userlogin")
	//未登录或登录超时
	if err != nil || "" == userlogin.Value {
		log.Debug("no login info found...")
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
	data["name"] = rsp.Data
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

func PostReg(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Info("PostReg  /api/v1.0/users")

	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	for key, value := range request {
		log.Info(key, value, reflect.TypeOf(value))
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
		log.Info("empty email password or emailcode")
		return
	}

	// call the backend service
	rsp, err := userClient.PostLogin(context.TODO(), &user.Request{
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
	log.Info("PostLoginInfo to  /api/v1.0/sessions")

	//decode the user input
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	for key, value := range request {
		log.Debug(key, value, reflect.TypeOf(value))
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
		log.Info("empty email or password")
		return
	}

	// call the backend service
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

//func WebsiteCall(w http.ResponseWriter, r *http.Request) {
//// decode the incoming request as json
//var request map[string]interface{}
//if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
//http.Error(w, err.Error(), 500)
//return
//}

//// call the backend service
//websiteClient := website.NewWebsiteService("micro.super.lightning.service.website", client.DefaultClient)
//rsp, err := websiteClient.Call(context.TODO(), &website.Request{
//Name: request["name"].(string),
//})
//if err != nil {
//http.Error(w, err.Error(), 500)
//return
//}

//// we want to augment the response
//response := map[string]interface{}{
//"msg": rsp.Msg,
//"ref": time.Now().UnixNano(),
//}

//// encode and write the response as json
//if err := json.NewEncoder(w).Encode(response); err != nil {
//http.Error(w, err.Error(), 500)
//return
//}
//}
