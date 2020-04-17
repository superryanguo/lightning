package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	"github.com/superryanguo/lightning/utils"
	//website "path/to/service/proto/website"
)

func GetIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Info("获取首页轮播 url：api/v1.0/lightning/index")

	//创建返回数据map
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
	log.Info("获取Session url：api/v1.0/session")

	//创建服务并初始化
	server := grpc.NewService()
	server.Init()

	// call the backend service
	exampleClient := GETSESSION.NewExampleService("go.micro.srv.GetSession", server.Client())

	//获取cookie
	userlogin, err := r.Cookie("userlogin")
	//未登录或登录超时
	if err != nil || "" == userlogin.Value {
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
	rsp, err := exampleClient.GetSession(context.TODO(), &GETSESSION.Request{
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

/*func WebsiteCall(w http.ResponseWriter, r *http.Request) {*/
//// decode the incoming request as json
//var request map[string]interface{}
//if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
//http.Error(w, err.Error(), 500)
//return
//}

//// call the backend service
//websiteClient := website.NewWebsiteService("go.micro.srv.website", client.DefaultClient)
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
/*}*/