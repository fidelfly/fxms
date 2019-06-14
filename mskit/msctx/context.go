package msctx

import (
	"context"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	uuid "github.com/satori/go.uuid"
)

const CtxTraceID = "Trace-Id"

func SetMetadata(ctx context.Context, data metadata.Metadata) context.Context {
	// copy metadata
	mda, _ := metadata.FromContext(ctx)
	md := metadata.Copy(mda)
	// set headers
	for k, v := range data {
		if _, ok := md[k]; !ok {
			md[k] = v
		}
	}
	return metadata.NewContext(ctx, md)
}

func GetFromService(ctx context.Context) string {
	mda, _ := metadata.FromContext(ctx)
	if name, ok := mda[micro.HeaderPrefix+"From-Service"]; ok {
		return name
	}
	return ""
}

func GetTraceID(ctx context.Context) string {
	mda, _ := metadata.FromContext(ctx)
	if id, ok := mda[CtxTraceID]; ok {
		return id
	}
	return ""
}

func GenerateTraceID() string {
	return uuid.NewV4().String()
}
