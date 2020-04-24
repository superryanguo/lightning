package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/superryanguo/lightning/user_srv/handler"
	"github.com/superryanguo/lightning/user_srv/subscriber"

	user_srv "github.com/superryanguo/lightning/user_srv/proto/user_srv"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("micro.super.lightning.srv.user_srv"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user_srv.RegisterUser_srvHandler(service.Server(), new(handler.User_srv))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("micro.super.lightning.srv.user_srv", service.Server(), new(subscriber.User_srv))

	// Register Function as Subscriber
	micro.RegisterSubscriber("micro.super.lightning.srv.user_srv", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
