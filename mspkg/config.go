package mspkg

import (
	"github.com/fidelfly/fxms/mspkg/conf"
	"github.com/fidelfly/fxms/mspkg/db"
)

var msCfg interface{}

func GetLogCfg() *conf.LogConfig {
	if msCfg == nil {
		return nil
	}
	if ms, ok := msCfg.(conf.MsConfigurable); ok {
		return ms.GetLog()
	}
	return nil
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

func IsDebug() bool {
	if msCfg == nil {
		return false
	}
	if cfg, ok := msCfg.(conf.MsConfigurable); ok {
		return cfg.GetStage() == conf.StageDevelopment
	}
	return false
}
