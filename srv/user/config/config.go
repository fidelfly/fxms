package config

import (
	"github.com/fidelfly/fxms/mspkg/conf"
	"github.com/fidelfly/fxms/mspkg/db"
)

type cfg struct {
	conf.MsConfig
	Database *db.Config
}

func (c cfg) GetDbConfig() *db.Config {
	return c.Database
}

var myCfg = &cfg{}

func Current() *cfg {
	return myCfg
}

var Database = myCfg.Database
