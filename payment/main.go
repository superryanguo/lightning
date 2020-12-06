package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/superryanguo/lightning/payment/handler"
	"github.com/superryanguo/lightning/payment/subscriber"

	payment "github.com/superryanguo/lightning/payment/proto/payment"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("micro.super.lightning.service.payment"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	payment.RegisterPaymentHandler(service.Server(), new(handler.Payment))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("micro.super.lightning.service.payment", service.Server(), new(subscriber.Payment))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
