module github.com/fidelfly/fxms/srv/auth

go 1.12

require (
	github.com/fidelfly/fxgo v0.0.0
	github.com/fidelfly/fxms/mskit v0.0.0
	github.com/go-xorm/xorm v0.7.1
	github.com/golang/protobuf v1.3.1
	github.com/micro/go-micro v1.5.0
	gopkg.in/oauth2.v3 v3.10.0
)

replace (
	github.com/fidelfly/fxgo => ../../../fxgo
	github.com/fidelfly/fxms/mskit => ../../mskit
)
