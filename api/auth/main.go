package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	"github.com/fidelfly/fxms/api/auth/client"
	"github.com/fidelfly/fxms/api/auth/handler"
	"github.com/fidelfly/fxms/api/auth/proto/auth"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.fxms.api.auth"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the Auth srv client
		micro.WrapHandler(client.AuthWrapper(service)),
	)

	// Register Handler
	auth.RegisterAuthHandler(service.Server(), new(handler.Auth))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
