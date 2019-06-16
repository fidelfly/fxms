module github.com/fidelfly/fxms/mspkg

go 1.12

require (
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/fidelfly/fxgo v0.0.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/xorm v0.7.1
	github.com/golang/protobuf v1.3.1
	github.com/micro/go-micro v1.5.0
	github.com/nats-io/nats-server/v2 v2.0.0 // indirect
	github.com/satori/go.uuid v1.2.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace github.com/fidelfly/fxgo => ../../fxgo
