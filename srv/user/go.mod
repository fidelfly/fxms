module github.com/fidelfly/fxms/srv/user

go 1.12

require (
	github.com/fidelfly/fxgo v0.0.0
	github.com/fidelfly/fxms/mspkg v0.0.0
	github.com/golang/protobuf v1.3.1
	github.com/micro/go-micro v1.5.0
)

replace (
	github.com/fidelfly/fxgo => ../../../fxgo
	github.com/fidelfly/fxms/mspkg => ../../mspkg
)
