package oauth2

import (
	"context"
	"strconv"

	"github.com/fidelfly/fxms/mspkg/rpcc"
	"github.com/fidelfly/fxms/srv/user/proto/user"
)

var userSrv user.UserService

func init() {
	userSrv = user.NewUserService("com.fxms.srv.user", rpcc.DefaultClient)
}

func passwordHandler(username, password string) (string, error) {
	rsp, err := userSrv.Validate(context.TODO(), &user.ValidateRequest{
		Code:     username,
		Password: password,
	})
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(rsp.Id, 10), nil
}
