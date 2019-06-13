package config

import (
	oauth22 "github.com/fidelfly/fxms/web/auth/oauth2"

	"github.com/fidelfly/fxms/mskit/conf"
	"github.com/fidelfly/fxms/mskit/db"
)

type cfg struct {
	conf.MsConfig
	Database *db.Config
	Auth     *oauth22.Config
}

func (c cfg) GetDbConfig() *db.Config {
	return c.Database
}

var myCfg = &cfg{}

func Current() *cfg {
	return myCfg
}

var Database = myCfg.Database

var AuthCfg = myCfg.Auth
