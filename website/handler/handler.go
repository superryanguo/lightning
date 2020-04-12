package handler

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-micro/util/log"
	//website "path/to/service/proto/website"
)

func GetIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Info("获取首页轮播 url：api/v1.0/lightning/index")

	//创建返回数据map
	response := map[string]interface{}{
		"errno":  "0",
		"errmsg": "Successful",
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
