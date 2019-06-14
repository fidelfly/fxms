package msconst

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/web"
)

type msConst struct {
	name    string
	version string
}

var info = &msConst{}

func WebConstOption(options *web.Options) {
	info.name = options.Name
	info.version = options.Version
}

func Version() string {
	return info.version
}

func Name() string {
	return info.name
}

func ConstOption(options *micro.Options) {
	info.name = options.Server.Options().Name
	info.version = options.Server.Options().Version
}
