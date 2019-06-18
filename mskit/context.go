package mskit

import (
	"net/http"

	"github.com/fidelfly/fxms/mskit/mscache"
	"github.com/fidelfly/fxms/mspkg/rdcache"
	"github.com/fidelfly/fxms/mspkg/wbr"
)

func GetUserInfo(r *http.Request) *mscache.UserInfo {
	userKey := wbr.GetUserKey(r)
	if len(userKey) > 0 {
		if cacheObj, err := rdcache.GetCache(mscache.UserInfoCache, userKey); err != nil {
			return nil
		} else if userInfo, ok := cacheObj.(*mscache.UserInfo); ok {
			return userInfo
		}
	}
	return nil
}
