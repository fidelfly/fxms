package oauth2

import (
	"github.com/fidelfly/fxgo/authx"
	"github.com/fidelfly/fxgo/logx"
	"github.com/micro/go-micro"

	"github.com/fidelfly/fxms/mskit/db"
	"github.com/fidelfly/fxms/mskit/oauth2"
)

var Server *authx.Server

func SetupAuthServer(db *db.DbInstance, authCfg *oauth2.Config) (server *authx.Server, err error) {
	store, err := oauth2.NewDbStore(db)
	if err != nil {
		return
	}
	server = authx.NewOAuthServer()
	server.SetTokenStorage(store)
	clients := make([]authx.ClientInfo, len(authCfg.Clients))
	for index, client := range authCfg.Clients {
		clients[index] = &client
	}
	server.SetClients(clients...)
	//server.SetPasswordAuthorizationHandler(pwdHandler)
	server.SetExtensionFieldsHandler(authx.NewTokenExtension(authx.UserExtension))
	return
}

func Initializer(instance *db.DbInstance, authCfg *oauth2.Config) micro.Option {
	return func(options *micro.Options) {
		as, err := SetupAuthServer(instance, authCfg)
		if err != nil {
			logx.Panic(err)
			panic(err)
		} else {
			Server = as
		}
	}
}
