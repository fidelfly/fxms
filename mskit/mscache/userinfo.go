package mscache

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/fidelfly/fxms/mspkg/proto/api"
	"github.com/fidelfly/fxms/mspkg/rpcc"
	"github.com/fidelfly/fxms/srv/user/proto/user"
)

var UserInfoCache *userInfoCache

type UserInfo struct {
	Id   int64
	Code string
	Name string
}

func init() {
	UserInfoCache = &userInfoCache{
		userSrv: user.NewUserService("com.fxms.srv.user", rpcc.DefaultClient),
	}
}

type userInfoCache struct {
	userSrv user.UserService
}

func (uic *userInfoCache) New() interface{} {
	return &UserInfo{}
}

func (uic *userInfoCache) GetRedisKey(k interface{}) string {
	return fmt.Sprintf("ms.cache.userinfo.%v", k)
}

func (uic *userInfoCache) GetDefaultExpiration() time.Duration {
	return 0
}

func (uic *userInfoCache) Resolve(k interface{}, v interface{}) (err error) {
	if v == nil {
		return errors.New("resolve: v is nil")
	}
	ui, ok := v.(*UserInfo)
	if !ok {
		return errors.New("resolve: v is not *UserInfo")
	}

	uid := int64(0)
	switch rv := k.(type) {
	case int64:
		uid = rv
	case string:
		uid, err = strconv.ParseInt(rv, 10, 64)
	}
	if err != nil {
		return
	}

	rsp, err := uic.userSrv.Read(context.Background(), &api.IdResponse{Id: uid})

	if err != nil {
		return
	}

	ui.Id = rsp.Id
	ui.Name = rsp.Name
	ui.Code = rsp.Code

	return
}
