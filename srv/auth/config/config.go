package config

import (
	"github.com/fidelfly/fxms/mskit/conf"
	"github.com/fidelfly/fxms/mskit/db"
	"github.com/fidelfly/fxms/mskit/oauth2"
)

type cfg struct {
	conf.MsConfig
	Database *db.Config
	Auth     *oauth2.Config
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
