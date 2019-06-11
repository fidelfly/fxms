package confx

import (
	"fmt"
	"os"
	"strings"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/encoder"
	"github.com/micro/go-micro/config/encoder/toml"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/consul"
	"github.com/micro/go-micro/config/source/env"
	"github.com/micro/go-micro/config/source/file"
)

func ReadConfig(cfg interface{}, files ...string) {
	cfgs := make([]source.Source, 0)
	enc := toml.NewEncoder()
	if len(files) > 0 {
		for _, file := range files {
			if _, err := os.Stat(file); err == nil {
				cfgs = append(cfgs, FileConfig(file, enc))
			}
		}
	}
	if len(cfgs) == 0 {
		return
	}

	config.Load(cfgs...)
	config.Scan(cfg)
}

func FileConfig(path string, enc encoder.Encoder) source.Source {
	return file.NewSource(file.WithPath(path), source.WithEncoder(enc))
}

//export
func EnvConfig(prefix ...string) source.Source {
	return env.NewSource(env.WithPrefix(prefix...))
}

//export
func ConsulConfig(address string, prefix string) source.Source {
	return consul.NewSource(consul.WithAddress(address), consul.WithPrefix(prefix))
}

func Initializer(cfg interface{}, files ...string) micro.Option {
	return func(options *micro.Options) {
		name := options.Server.Options().Name
		cfgFiles := files
		if len(cfgFiles) == 0 {
			cfgFiles = append(cfgFiles, "/config.toml", fmt.Sprintf("/%s.toml", strings.ReplaceAll(name, ".", "_")))
		}

		ReadConfig(cfg, cfgFiles...)
	}
}
