package wbr

import (
	"net/http"

	"github.com/fidelfly/fxgo"
	"github.com/fidelfly/fxgo/httprxr"
	"github.com/micro/go-micro/metadata"
	"gopkg.in/oauth2.v3"

	"github.com/fidelfly/fxms/mspkg/msctx"
)

func MsHandler(handler http.Handler) http.Handler {
	return fxgo.AttachMiddleware(handler, TraceMiddleware)
}

func MsHandlerFunc(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = msctx.SetMetadata(ctx, metadata.Metadata{
			msctx.CtxTraceID: msctx.GenerateTraceID(),
		})
		handlerFunc(w, r.WithContext(ctx))
	}
}

func TraceMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = msctx.SetMetadata(ctx, metadata.Metadata{
			msctx.CtxTraceID: msctx.GenerateTraceID(),
		})
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AuthMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if td := httprxr.ContextGet(r, fxgo.ContextTokenKey); td != nil {
			if ti, ok := td.(oauth2.TokenInfo); ok {
				ctx = msctx.SetMetadata(ctx, metadata.Metadata{
					msctx.CtxUserID: ti.GetUserID(),
				})
				r = r.WithContext(ctx)
			}
		}

		handler.ServeHTTP(w, r)
	})
}
