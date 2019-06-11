module github.com/fidelfly/fxms/api/auth

go 1.12

require (
	github.com/fidelfly/fxms/srv/auth v0.0.0
	github.com/golang/protobuf v1.3.1
	github.com/micro/go-micro v1.5.0
)

replace github.com/fidelfly/fxms/srv/auth => ../../srv/auth
