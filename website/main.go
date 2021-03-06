package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/micro/cli/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
	"github.com/superryanguo/lightning/basic"
	"github.com/superryanguo/lightning/models"
	"github.com/superryanguo/lightning/website/handler"
)

const (
	webPort = ":8081"
)

func main() {
	basic.Init()
	models.Init()
	// create new web service
	service := web.NewService(
		web.Name("micro.super.lightning.web.website"),
		web.Version("latest"),
		web.Address(webPort),
	)

	// initialise service
	if err := service.Init(
		web.Action(
			func(c *cli.Context) {
				// 初始化handler
				handler.Init()
			}),
	); err != nil {
		log.Fatal(err)
	}

	rou := httprouter.New()
	rou.NotFound = http.FileServer(http.Dir("html"))
	rou.GET("/api/v1.0/imagecode/:uuid", handler.GetImageCd)
	rou.GET("/api/v1.0/emailcode/:email", handler.GetEmailCd)
	//获取地区数据
	rou.GET("/api/v1.0/lightning/areas", handler.GetArea)
	////注册
	rou.POST("/api/v1.0/users", handler.PostReg)
	//获取session
	rou.GET("/api/v1.0/session", handler.GetSession)
	////登录
	rou.POST("/api/v1.0/userlogin", handler.PostLogin)
	//登出
	rou.DELETE("/api/v1.0/session", handler.DeleteSession)
	//获取用户信息
	rou.GET("/api/v1.0/user", handler.GetUserInfo)
	//获取首页轮播图
	rou.GET("/api/v1.0/lightning/house/index", handler.GetIndex) //TODO:not sure the relatin with the v1.0/houses/ address
	//上传用户头像
	rou.POST("/api/v1.0/user/avatar", handler.PostAvatar)
	//修改用户名
	rou.PUT("/api/v1.0/user/name", handler.PutUserInfo)
	//查看用户是否实名认证
	rou.GET("/api/v1.0/user/infoauth", handler.GetUserInfo)
	//进行实名认证
	rou.POST("/api/v1.0/user/infoauth", handler.PostUserAuth)

	// register html handler
	//service.Handle("/", http.FileServer(http.Dir("html")))
	service.Handle("/", rou)

	// register call handler
	//service.HandleFunc("/website/call", handler.WebsiteCall)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
