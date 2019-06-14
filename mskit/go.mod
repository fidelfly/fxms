module github.com/fidelfly/fxms/mskit

go 1.12

require (
	github.com/fidelfly/fxgo v0.0.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/xorm v0.7.1
	github.com/micro/go-micro v1.5.0
	github.com/nats-io/nats-server/v2 v2.0.0 // indirect
	github.com/satori/go.uuid v1.2.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/oauth2.v3 v3.10.0
)

replace github.com/fidelfly/fxgo => ../../fxgo
