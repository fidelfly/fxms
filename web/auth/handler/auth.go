package handler

import (
	"github.com/fidelfly/fxms/mspkg/wbr"
	"github.com/fidelfly/fxms/web/auth/oauth2"

	"github.com/micro/go-micro/web"
)

func AuthHandler(service web.Service) {
	service.Handle("/auth/token", wbr.MsHandlerFunc(oauth2.TokenIssuer.HandleTokenRequest))
}
