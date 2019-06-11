package db

import (
	"github.com/fidelfly/fxgo/logx"
	"github.com/go-xorm/xorm"
	"github.com/micro/go-micro"
)

var Engine *xorm.Engine

//export
func Initializer(instance *Instance) micro.Option {
	return func(options *micro.Options) {
		engine, err := InitEngine(instance)
		if err != nil {
			logx.Panic(err)
			panic(err)
		} else {
			Engine = engine
		}
	}
}
