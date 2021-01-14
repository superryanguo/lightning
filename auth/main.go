package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/superryanguo/lightning/auth/handler"
	"github.com/superryanguo/lightning/auth/subscriber"
	"github.com/superryanguo/lightning/basic"
	"github.com/superryanguo/lightning/models"

	auth "github.com/superryanguo/lightning/auth/proto/auth"
)

func main() {
	basic.Init()
	// New Service
	service := micro.NewService(
		micro.Name("micro.super.lightning.service.auth"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		micro.Action(func(c *cli.Context) error {
			models.Init()
			handler.Init()
			return nil
		}),
	)

	// Register Handler
	auth.RegisterAuthHandler(service.Server(), new(handler.Auth))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("micro.super.lightning.service.auth", service.Server(), new(subscriber.Auth))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
