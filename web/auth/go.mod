module github.com/fidelfly/fxms/web/auth

go 1.12

require (
	github.com/fidelfly/fxgo v0.0.0
	github.com/fidelfly/fxms/mspkg v0.0.0
	github.com/fidelfly/fxms/srv/auth v0.0.0
	github.com/micro/go-micro v1.5.0
	gopkg.in/oauth2.v3 v3.10.0
)

replace (
	github.com/fidelfly/fxgo => ../../../fxgo
	github.com/fidelfly/fxms/mspkg => ../../mspkg
	github.com/fidelfly/fxms/srv/auth => ../../srv/auth
)
