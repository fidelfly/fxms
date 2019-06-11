package db

import (
	"github.com/fidelfly/fxgo/logx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//export
func InitEngine(instance *Instance) (engine *xorm.Engine, err error) {
	engine, err = xorm.NewEngine("mysql", instance.getUrl())
	if err == nil {
		terr := engine.Ping()
		if terr == nil {
			logx.Infof("Database(%s) is connected!", instance.getTarget())
		} else {
			logx.Errorf("Can't connect to database(%s)", instance.getTarget())
		}
	}
	return
}
