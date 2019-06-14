package whdr

import (
	"net/http"

	"github.com/fidelfly/fxms/mskit/msctx"
)

func GetFromService(r *http.Request) string {
	return msctx.GetFromService(r.Context())
}

func GetTraceID(r *http.Request) string {
	return msctx.GetTraceID(r.Context())
}
