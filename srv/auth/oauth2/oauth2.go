package oauth2

import (
	"github.com/micro/go-micro"

	"github.com/fidelfly/fxms/mskit/db"
)

var tokenStore *dbStore

func GetTokenStore() *dbStore {
	return tokenStore
}

func Initializer(instance *db.Config) micro.Option {
	return func(options *micro.Options) {
		store, err := NewDbStore(instance)
		if err != nil {
			return
		}
		tokenStore = store
	}
}
