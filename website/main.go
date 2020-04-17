package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-micro/util/log"

	"github.com/micro/go-micro/web"
	"github.com/superryanguo/lightning/website/handler"
)

const (
	webPort = ":8081"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.lightning.website"),
		web.Version("latest"),
		web.Address(webPort),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// register call handler
	//service.HandleFunc("/website/call", handler.WebsiteCall)

	rou := httprouter.New()
	//映射静态页面
	rou.NotFound = http.FileServer(http.Dir("html"))
	//rou.GET("/api/v1.0/imagecode/:uuid", handler.GetImageCd)
	////获取邮箱验证码
	//rou.GET("/api/v1.0/emailcode/:email", handler.GetEmailCd)
	////注册
	//rou.POST("/api/v1.0/users", handler.PostReg)
	//获取session
	rou.GET("/api/v1.0/session", handler.GetSession)
	////登录
	//rou.POST("/api/v1.0/sessions", handler.PostLogin)
	////登出
	//rou.DELETE("/api/v1.0/session", handler.DeleteSession)
	////获取用户信息
	//rou.GET("/api/v1.0/user", handler.GetUserInfo)
	//获取首页轮播图
	rou.GET("/api/v1.0/lightning/index", handler.GetIndex)
	//上传用户头像
	//rou.POST("/api/v1.0/user/avatar", handler.PostAvatar)
	////修改用户名
	//rou.PUT("/api/v1.0/user/name", handler.PutUserInfo)
	////查看用户是否实名认证
	//rou.GET("/api/v1.0/user/auth", handler.GetUserInfo)
	////进行实名认证
	//rou.POST("/api/v1.0/user/auth", handler.PostUserAuth)

	// register html handler
	//service.Handle("/", http.FileServer(http.Dir("html")))
	service.Handle("/", rou)
	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}