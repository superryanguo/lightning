package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/superryanguo/lightning/basic"
	"github.com/superryanguo/lightning/models"
	"github.com/superryanguo/lightning/user_srv/handler"
	"github.com/superryanguo/lightning/user_srv/subscriber"

	user_srv "github.com/superryanguo/lightning/user_srv/proto/user_srv"
)

func main() {
	basic.Init()
	models.Init()
	handler.Init() // I want to control the init
	// New Service
	service := micro.NewService(
		micro.Name("micro.super.lightning.service.user_srv"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user_srv.RegisterUserSrvHandler(service.Server(), new(handler.User_srv))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("micro.super.lightning.service.user_srv", service.Server(), new(subscriber.User_srv))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
