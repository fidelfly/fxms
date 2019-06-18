package main

import (
	"github.com/micro/go-micro/util/log"

	"github.com/fidelfly/fxms/mspkg"
	"github.com/fidelfly/fxms/web/user/config"

	"github.com/micro/go-micro/web"

	"github.com/fidelfly/fxms/web/user/handler"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("com.fxms.web.user"),
		web.Version("latest"),
	)

	// initialise service
	if err := service.Init(
		mspkg.WebServiceInitiallizer(),               // Initialize web service
		mspkg.WebConfigInitializer(config.Current()), // Read configuration
		mspkg.WebLogInitializer(),                    // Initialize log api
	); err != nil {
		log.Fatal(err)
	}

	// register call handler
	service.HandleFunc("/user/call", handler.UserCall)

	service.Handle("/user", handler.UserHandler)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
