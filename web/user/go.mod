module github.com/fidelfly/fxms/web/user

go 1.12

require (
	github.com/fidelfly/fxgo v0.0.0
	github.com/fidelfly/fxms/mskit v0.0.0-00010101000000-000000000000
	github.com/fidelfly/fxms/mspkg v0.0.0
	github.com/fidelfly/fxms/srv/user v0.0.0
	github.com/micro/go-micro v1.5.0
)

replace (
	github.com/fidelfly/fxgo => ../../../fxgo
	github.com/fidelfly/fxms/mskit => ../../mskit
	github.com/fidelfly/fxms/mspkg => ../../mspkg
	github.com/fidelfly/fxms/srv/auth => ../../srv/auth
	github.com/fidelfly/fxms/srv/user => ../../srv/user
)
