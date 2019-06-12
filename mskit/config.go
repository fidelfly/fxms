package mskit

import (
	"github.com/fidelfly/fxms/mskit/conf"
	"github.com/fidelfly/fxms/mskit/db"
)

var msCfg interface{}

func GetVersion() string {
	if msCfg == nil {
		return ""
	}
	if ms, ok := msCfg.(conf.MsConfigurable); ok {
		return ms.GetVersion()
	}
	return ""
}

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
	if db, ok := msCfg.(db.Configurable); ok {
		return db.GetDbConfig()
	}
	return nil
}
