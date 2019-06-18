package msctx

import (
	"context"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	uuid "github.com/satori/go.uuid"
)

const CtxTraceID = "Trace-Id"

const CtxUserID = "User-Id"

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

//export
func GetMetaValue(ctx context.Context, key string) (string, bool) {
	mda, _ := metadata.FromContext(ctx)
	if name, ok := mda[key]; ok {
		return name, ok
	}
	return "", false
}

func GetFromService(ctx context.Context) string {
	name, _ := GetMetaValue(ctx, micro.HeaderPrefix+"From-Service")
	return name
}

func GetTraceID(ctx context.Context) string {
	id, _ := GetMetaValue(ctx, CtxTraceID)
	return id
}

func GenerateTraceID() string {
	return uuid.NewV4().String()
}

func GetUserID(ctx context.Context) string {
	uid, _ := GetMetaValue(ctx, CtxUserID)
	return uid
}
