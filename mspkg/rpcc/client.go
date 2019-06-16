package rpcc

import (
	"context"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"

	"github.com/fidelfly/fxms/mspkg/msconst"
	"github.com/fidelfly/fxms/mspkg/msctx"
)

var DefaultClient = NewClient()

func NewClient(opts ...client.Option) client.Client {
	newOpts := append([]client.Option{TraceOption}, opts...)
	return client.NewClient(newOpts...)
}

func TraceOption(cliOptions *client.Options) {
	cliOptions.CallOptions.CallWrappers = append([]client.CallWrapper{TraceWrapper}, cliOptions.CallOptions.CallWrappers...)
}

func TraceWrapper(call client.CallFunc) client.CallFunc {
	return func(ctx context.Context, node *registry.Node, req client.Request, rsp interface{}, opts client.CallOptions) error {
		data := metadata.Metadata{
			micro.HeaderPrefix + "From-Service": msconst.Name(),
		}

		if _, ok := data[msctx.CtxTraceID]; !ok {
			data[msctx.CtxTraceID] = msctx.GenerateTraceID()
		}

		ctx = msctx.SetMetadata(ctx, data)
		return call(ctx, node, req, rsp, opts)
	}
}
