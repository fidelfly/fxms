package oauth2

import (
	"net/http"

	"github.com/fidelfly/fxgo/authx"
	"github.com/fidelfly/fxgo/logx"
	"github.com/micro/go-micro/web"

	"github.com/fidelfly/fxms/mskit/oauth2"
)

var server *authx.Server

func setupAuthServer(authCfg *oauth2.Config) (server *authx.Server, err error) {
	server = authx.NewOAuthServer()
	server.SetTokenStorage(NewTokenStore())
	clients := make([]authx.ClientInfo, len(authCfg.Clients))
	for index, client := range authCfg.Clients {
		clients[index] = &client
	}
	server.SetClients(clients...)
	server.SetPasswordAuthorizationHandler(passwordHandler)
	server.SetExtensionFieldsHandler(authx.NewTokenExtension(authx.UserExtension))
	return
}

func Initializer(authCfg *oauth2.Config) web.Option {
	return func(options *web.Options) {
		as, err := setupAuthServer(authCfg)
		if err != nil {
			logx.Panic(err)
			panic(err)
		} else {
			server = as
		}
	}
}

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	server.HandleTokenRequest(w, r)
}
