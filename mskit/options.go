package mskit

import (
	"github.com/micro/go-micro"
)

func LogInitializer(level string) micro.Option {
	return func(options *micro.Options) {
		SetupLog(level, options.Server.Options().Name)
	}
}
