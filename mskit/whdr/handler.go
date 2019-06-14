package whdr

import (
	"net/http"

	"github.com/micro/go-micro/metadata"

	"github.com/fidelfly/fxms/mskit/msctx"
)

func MsHandler(handler http.Handler) http.Handler {
	return TraceMiddleware(handler)
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
