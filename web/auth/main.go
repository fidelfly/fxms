package main

import (
	"github.com/fidelfly/fxgo/logx"
	"github.com/micro/go-micro/web"

	"github.com/fidelfly/fxms/mskit"
	"github.com/fidelfly/fxms/mskit/wbr"
	"github.com/fidelfly/fxms/web/auth/config"
	"github.com/fidelfly/fxms/web/auth/handler"
	"github.com/fidelfly/fxms/web/auth/oauth2"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("com.fxms.web.auth"),
		web.Version("latest"),
	)

	// initialise service
	if err := service.Init(
		mskit.WebServiceInitiallizer(),
		mskit.WebConfigInitializer(config.Current()),
		mskit.WebLogInitializer(),
		oauth2.Initializer(config.AuthCfg),
	); err != nil {
		logx.Panic(err)
		return
	}

	// register call handler
	service.HandleFunc("/auth/call", wbr.MsHandlerFunc(handler.AuthCall))

	service.HandleFunc("/auth/token", wbr.MsHandlerFunc(oauth2.TokenHandler))

	// run service
	if err := service.Run(); err != nil {
		logx.Panic(err)
		return
	}
}
