package oauth2

import "github.com/fidelfly/fxgo/authx"

type Config struct {
	Clients []authx.AuthClient
}

type Configurable interface {
	GetAuthConfig() *Config
}
