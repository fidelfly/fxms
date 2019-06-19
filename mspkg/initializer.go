package mspkg

import (
	"fmt"
	"strings"
	"time"

	"github.com/fidelfly/fxgo/logx"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/web"

	"github.com/fidelfly/fxms/mspkg/conf"
	"github.com/fidelfly/fxms/mspkg/db"
	"github.com/fidelfly/fxms/mspkg/msconst"
	"github.com/fidelfly/fxms/mspkg/rdcache"
	"github.com/fidelfly/fxms/mspkg/rpcc"
)

//export
func LogInitializer(levels ...string) micro.Option {
	return func(options *micro.Options) {
		level := ""
		if len(levels) > 0 {
			level = levels[0]
		}
		if len(level) == 0 {
			if logCfg := GetLogCfg(); logCfg != nil {
				level = logCfg.GetLevel()
			}
		}
		if len(level) == 0 {
			level = "warning"
		}
		SetupLog(level, options.Server.Options().Name)
	}
}

//export
func WebLogInitializer(levels ...string) web.Option {
	return func(options *web.Options) {
		level := ""
		if len(levels) > 0 {
			level = levels[0]
		}
		if len(level) == 0 {
			if logCfg := GetLogCfg(); logCfg != nil {
				level = logCfg.GetLevel()
			}
		}
		if len(level) == 0 {
			level = "warning"
		}
		SetupLog(level, options.Name)
	}
}

//export
func DbInitializer(cfgs ...*db.Config) micro.Option {
	return func(options *micro.Options) {
		var cfg *db.Config
		if len(cfgs) > 0 {
			cfg = cfgs[0]
		} else {
			cfg = GetDbConfig()
		}
		if cfg != nil {
			logx.CaptureError(db.InitEngine(cfg))
		} else {
			logx.Error("database config is not found")
		}

	}
}

//export
func DbSynchronize(targets ...interface{}) micro.Option {
	return func(options *micro.Options) {
		if db.Engine == nil {
			panic("database engine is not initialized")
		}

		if len(targets) > 0 {
			logx.CaptureError(db.Engine.Sync(targets...))
		}
	}
}

//export
func WebDbInitializer(cfgs ...*db.Config) web.Option {
	return func(options *web.Options) {
		var cfg *db.Config
		if len(cfgs) > 0 {
			cfg = cfgs[0]
		} else {
			cfg = GetDbConfig()
		}
		if cfg != nil {
			logx.CaptureError(db.InitEngine(cfg))
		} else {
			logx.Error("database config is not found")
		}

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

		conf.ReadConfig(cfg, cfgFiles...)

		msCfg = cfg
	}
}

//export
func WebConfigInitializer(cfg interface{}, files ...string) web.Option {
	return func(options *web.Options) {
		name := options.Name
		cfgFiles := files
		if len(cfgFiles) == 0 {
			cfgFiles = append(cfgFiles, "/config.toml", fmt.Sprintf("/%s.toml", strings.ReplaceAll(name, ".", "_")))
		}

		conf.ReadConfig(cfg, cfgFiles...)

		msCfg = cfg
	}
}

//export
func ServiceInitiallizer() micro.Option {
	return func(options *micro.Options) {
		msconst.ConstOption(options)
		DebugInit(options)
	}
}

//export
func WebServiceInitiallizer() web.Option {
	return func(options *web.Options) {
		msconst.WebConstOption(options)
		WebDebugInit(options)
	}
}

//export
func CacheInitializer(cfg *rdcache.Config) micro.Option {
	return func(options *micro.Options) {
		rdcache.InitRedis(cfg)
		rdcache.InitCache(rdcache.Redisdb)
	}
}

//export
func WebCacheInitializer(cfg *rdcache.Config) web.Option {
	return func(options *web.Options) {
		rdcache.InitRedis(cfg)
		rdcache.InitCache(rdcache.Redisdb)
	}
}

func DebugInit(options *micro.Options) {
	debug()
}

func WebDebugInit(options *web.Options) {
	debug()
}

func debug() {
	if IsDebug() {
		rpcc.DefaultClient.Init(func(options *client.Options) {
			options.CallOptions.RequestTimeout = 10 * time.Minute
		})
	}
}

/*//export
func DebugInitializer() micro.Option {
	return func(options *micro.Options) {
		if IsDebug() {
			rpcc.DefaultClient.Init(func(options *client.Options) {
				options.CallOptions.RequestTimeout = 10 * time.Minute
			})
		}
	}
}*/
