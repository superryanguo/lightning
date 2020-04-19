package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"github.com/superryanguo/lightning/basic"
	"github.com/superryanguo/lightning/session_mgr/handler"
	"github.com/superryanguo/lightning/session_mgr/subscriber"

	session_mgr "github.com/superryanguo/lightning/session_mgr/proto/session_mgr"
)

func main() {
	basic.Init()
	// New Service
	service := micro.NewService(
		micro.Name("micro.super.lightning.srv.session_mgr"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	session_mgr.RegisterSession_mgrHandler(service.Server(), new(handler.Session_mgr))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("micro.super.lightning.srv.session_mgr", service.Server(), new(subscriber.Session_mgr))

	// Register Function as Subscriber
	micro.RegisterSubscriber("micro.super.lightning.srv.session_mgr", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
