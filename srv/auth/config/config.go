package config

import (
	"github.com/fidelfly/fxms/mspkg/conf"
	"github.com/fidelfly/fxms/mspkg/db"
	"github.com/fidelfly/fxms/mspkg/rdcache"
)

type cfg struct {
	conf.MsConfig
	Database db.Config
	Redis    rdcache.Config
}

func (c cfg) GetDbConfig() *db.Config {
	return &c.Database
}

var myCfg = &cfg{}

func Current() *cfg {
	return myCfg
}

var Database = &myCfg.Database

var RedisCfg = &myCfg.Redis
