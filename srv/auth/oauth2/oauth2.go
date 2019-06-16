package oauth2

import (
	"github.com/fidelfly/fxgo/authx"
	"github.com/micro/go-micro"
	"gopkg.in/oauth2.v3"

	"github.com/fidelfly/fxms/mskit/db"
)

var tokenStore oauth2.TokenStore

func GetTokenStore() oauth2.TokenStore {
	return tokenStore
}

func Initializer(instance *db.Config) micro.Option {
	return func(options *micro.Options) {
		store, err := NewDbStore(instance)
		if err != nil {
			return
		}
		tokenStore = authx.NewMultiLevelTokenStore(authx.NewMemoryTokenStore(), store)
	}
}
