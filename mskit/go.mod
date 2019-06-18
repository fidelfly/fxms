module github.com/fidelfly/fxms/mskit

go 1.12

require (
	github.com/fidelfly/fxgo v0.0.0
	github.com/fidelfly/fxms/mspkg v0.0.0
	github.com/fidelfly/fxms/srv/auth v0.0.0
	github.com/fidelfly/fxms/srv/user v0.0.0
)

replace (
	github.com/fidelfly/fxgo => ../../fxgo
	github.com/fidelfly/fxms/mspkg => ../mspkg
	github.com/fidelfly/fxms/srv/auth => ../srv/auth
	github.com/fidelfly/fxms/srv/user => ../srv/user
)
