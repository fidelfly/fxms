package mskit

import (
	"net/http"

	"github.com/fidelfly/fxgo/routex"

	"github.com/fidelfly/fxms/mskit/intercall"
	"github.com/fidelfly/fxms/mspkg/wbr"
)

func NewWebRouter(prefix string, middlewares ...func(handler http.Handler) http.Handler) *routex.Router {
	return wbr.NewAuthRouter(prefix, intercall.AuthServerFilter, middlewares...)
}
