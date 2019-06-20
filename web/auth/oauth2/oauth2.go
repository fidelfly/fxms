package oauth2

import (
	"github.com/fidelfly/fxgo"
	"github.com/fidelfly/fxgo/authx"
	"github.com/fidelfly/fxgo/logx"
	"github.com/micro/go-micro/web"
)

//var server *authx.Server

var TokenIssuer *fxgo.TokenIssuer

func setupAuthServer(authCfg *Config) (server *authx.Server, err error) {
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

func Initializer(authCfg *Config) web.Option {
	return func(options *web.Options) {
		as, err := setupAuthServer(authCfg)
		if err != nil {
			logx.Panic(err)
			panic(err)
		} else {
			TokenIssuer = fxgo.NewTokenIssuer(as, "")
		}
	}
}
