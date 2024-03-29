package main

import (
	"github.com/micro/go-micro/util/log"

	"github.com/fidelfly/fxms/mspkg"
	"github.com/fidelfly/fxms/mspkg/wbr"
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
		mspkg.WebConfigInitializer(config.Current()), // Read configuration
		mspkg.WebLogInitializer(),                    // Initialize log api
		mspkg.WebServiceInitiallizer(),               // Initialize web service
	); err != nil {
		log.Fatal(err)
	}

	// register call handler
	service.HandleFunc("/user/call", handler.UserCall)

	wbr.RegisterRoute(service, handler.UserHandler)
	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
