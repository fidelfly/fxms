package config

import (
	"github.com/fidelfly/fxms/mspkg/conf"
	"github.com/fidelfly/fxms/web/auth/oauth2"
)

type cfg struct {
	conf.MsConfig
	Auth oauth2.Config
}

var myCfg = &cfg{}

func Current() *cfg {
	return myCfg
}

var AuthCfg = &myCfg.Auth
