package wbr

import (
	"net/http"

	"github.com/fidelfly/fxgo"
	"github.com/fidelfly/fxgo/logx"
	"github.com/fidelfly/fxgo/routex"
	"github.com/micro/go-micro/web"
)

func NewRouter(prefix string, middlewares ...func(handler http.Handler) http.Handler) *routex.Router {
	rr := fxgo.NewRouter(prefix)
	rr.EnableAudit(logx.StandardLogger())
	rr.Use(middlewares...)
	return rr.PathPrefix(prefix).Subrouter()
}

//export
func NewAuthRouter(prefix string, authFilter func(w http.ResponseWriter, req *http.Request, next http.Handler),
	middlewares ...func(handler http.Handler) http.Handler) *routex.Router {
	rr := fxgo.NewRouter(prefix)
	rr.Use(TraceMiddleware)
	rr.EnableAudit(logx.StandardLogger())
	rr.EnableAuthFilter(authFilter)
	rr.Use(AuthMiddleware)
	rr.Use(middlewares...)
	return rr.PathPrefix(prefix).Subrouter()
}

type WebRouteRegister func(server web.Service)

//export
func RegisterRoute(service web.Service, registers ...WebRouteRegister) {
	if len(registers) > 0 {
		for _, reg := range registers {
			reg(service)
		}
	}
}
