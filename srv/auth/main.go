package main

import (
	"github.com/fidelfly/fxgo/logx"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	"github.com/fidelfly/fxms/mskit"
	"github.com/fidelfly/fxms/mskit/confx"
	"github.com/fidelfly/fxms/srv/auth/config"
	"github.com/fidelfly/fxms/srv/auth/handler"
	"github.com/fidelfly/fxms/srv/auth/oauth2"
	"github.com/fidelfly/fxms/srv/auth/proto/auth"
	"github.com/fidelfly/fxms/srv/auth/subscriber"
)

func main() {

	const srvName = "com.fxms.srv.auth"

	// New Service
	service := micro.NewService(
		micro.Name(srvName),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		confx.Initializer(config.Current()),
		mskit.LogInitializer("warning"),
		oauth2.Initializer(config.Database, config.AuthCfg),
	)

	logx.CaptureError(
		// Register Handler
		auth.RegisterAuthHandler(service.Server(), new(handler.Auth)),

		// Register Struct as Subscriber
		micro.RegisterSubscriber(srvName, service.Server(), new(subscriber.Auth)),

		// Register Function as Subscriber
		micro.RegisterSubscriber(srvName, service.Server(), subscriber.Handler),
	)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
