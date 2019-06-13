package config

import (
	"github.com/fidelfly/fxms/mskit/conf"
	oauth22 "github.com/fidelfly/fxms/web/auth/oauth2"
)

type cfg struct {
	conf.MsConfig
	Auth *oauth22.Config
}

var myCfg = &cfg{}

func Current() *cfg {
	return myCfg
}

var AuthCfg = myCfg.Auth
