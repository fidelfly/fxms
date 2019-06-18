package wbr

import (
	"net/http"
	"strconv"

	"github.com/fidelfly/fxgo"
	"github.com/fidelfly/fxgo/httprxr"

	"github.com/fidelfly/fxms/mspkg/msctx"
)

func GetFromService(r *http.Request) string {
	return msctx.GetFromService(r.Context())
}

func GetTraceID(r *http.Request) string {
	return msctx.GetTraceID(r.Context())
}

func GetUserID(r *http.Request) int64 {
	ctxObj := httprxr.ContextGet(r, fxgo.ContextUserKey)
	if ctxObj == nil {
		return 0
	}
	if userKey, ok := ctxObj.(string); ok {
		userId, _ := strconv.ParseInt(userKey, 10, 64)
		return userId
	}
	return 0
}

func GetUserKey(r *http.Request) string {
	ctxObj := httprxr.ContextGet(r, fxgo.ContextUserKey)
	if ctxObj == nil {
		return ""
	}
	if userKey, ok := ctxObj.(string); ok {
		return userKey
	}
	return ""
}
