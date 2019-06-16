package wbr

import (
	"github.com/fidelfly/fxgo"
	"github.com/fidelfly/fxgo/routex"
)

func NewRouter(prefix string) *routex.Router {
	rr := fxgo.NewRouter(prefix)

	return rr.PathPrefix(prefix).Subrouter()
}
