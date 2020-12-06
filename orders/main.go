package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/superryanguo/lightning/orders/handler"
	"github.com/superryanguo/lightning/orders/subscriber"

	orders "github.com/superryanguo/lightning/orders/proto/orders"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("micro.super.lightning.service.orders"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	orders.RegisterOrdersHandler(service.Server(), new(handler.Orders))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("micro.super.lightning.service.orders", service.Server(), new(subscriber.Orders))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
