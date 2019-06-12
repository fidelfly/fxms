package mskit

import (
	"fmt"
	"strings"

	"github.com/micro/go-micro"

	"github.com/fidelfly/fxms/mskit/confx"
	"github.com/fidelfly/fxms/mskit/db"
)

//export
func LogInitializer(level string) micro.Option {
	return func(options *micro.Options) {
		SetupLog(level, options.Server.Options().Name)
	}
}

//export
func DbInitializer(instance *db.Instance) micro.Option {
	return func(options *micro.Options) {
		db.InitEngine(instance)
	}
}

//export
func ConfigInitializer(cfg interface{}, files ...string) micro.Option {
	return func(options *micro.Options) {
		name := options.Server.Options().Name
		cfgFiles := files
		if len(cfgFiles) == 0 {
			cfgFiles = append(cfgFiles, "/config.toml", fmt.Sprintf("/%s.toml", strings.ReplaceAll(name, ".", "_")))
		}

		confx.ReadConfig(cfg, cfgFiles...)
	}
}
