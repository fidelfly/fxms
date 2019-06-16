package mspkg

import (
	"github.com/fidelfly/fxms/mspkg/conf"
	"github.com/fidelfly/fxms/mspkg/db"
)

var msCfg interface{}

func GetLogLevel() string {
	if msCfg == nil {
		return ""
	}
	if ms, ok := msCfg.(conf.MsConfigurable); ok {
		return ms.GetLogLevel()
	}
	return ""
}

func GetDbConfig() *db.Config {
	if msCfg == nil {
		return nil
	}
	if cfg, ok := msCfg.(db.Configurable); ok {
		return cfg.GetDbConfig()
	}
	return nil
}
