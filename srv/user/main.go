package main

import (
	"github.com/fidelfly/fxgo/logx"
	"github.com/micro/go-micro"

	"github.com/fidelfly/fxms/mspkg"
	"github.com/fidelfly/fxms/srv/user/config"
	"github.com/fidelfly/fxms/srv/user/handler"
	"github.com/fidelfly/fxms/srv/user/proto/user"
	"github.com/fidelfly/fxms/srv/user/res"
	"github.com/fidelfly/fxms/srv/user/subscriber"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.fxms.srv.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		mspkg.ConfigInitializer(config.Current()), // Read configuration
		mspkg.LogInitializer(),                    // Initial log api
		mspkg.DbInitializer(),                     // Initial database api
		mspkg.DbSynchronize(new(res.User)),        // Synchronize tables
		Initializer(),                             // Other Initial tasks
	)

	logx.CaptureError(
		// Register Handler
		user.RegisterUserHandler(service.Server(), new(handler.User)),

		// Register Struct as Subscriber
		micro.RegisterSubscriber("com.fxms.srv.user", service.Server(), new(subscriber.User)),

		// Register Function as Subscriber
		micro.RegisterSubscriber("com.fxms.srv.user", service.Server(), subscriber.Handler),
	)

	// Run service
	if err := service.Run(); err != nil {
		logx.Panic(err)
	}
}
