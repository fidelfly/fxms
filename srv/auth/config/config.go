package config

import (
	"github.com/fidelfly/fxms/mskit/db"
	"github.com/fidelfly/fxms/mskit/oauth2"
)

type cfg struct {
	Database db.DbInstance
	AuthCfg  oauth2.Config
}

var myCfg = &cfg{}

func Current() *cfg {
	return myCfg
}

var Database = &myCfg.Database

var AuthCfg = &myCfg.AuthCfg