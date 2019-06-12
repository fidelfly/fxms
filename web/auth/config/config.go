package config

import (
	"github.com/fidelfly/fxms/mskit/conf"
	"github.com/fidelfly/fxms/mskit/oauth2"
)

type cfg struct {
	conf.MsConfig
	Auth *oauth2.Config
}

var myCfg = &cfg{}

func Current() *cfg {
	return myCfg
}

var AuthCfg = myCfg.Auth
