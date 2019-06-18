package config

import (
	"github.com/fidelfly/fxms/mspkg/conf"
	"github.com/fidelfly/fxms/mspkg/rdcache"
)

type cfg struct {
	conf.MsConfig
	Redis rdcache.Config
}

var myCfg = &cfg{}

func Current() *cfg {
	return myCfg
}

var RedisCfg = &myCfg.Redis
