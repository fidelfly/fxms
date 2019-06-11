package handler

import (
	"context"
	"encoding/json"

	"github.com/micro/go-micro/util/log"

	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"

	"github.com/fidelfly/fxms/api/auth/client"
	"github.com/fidelfly/fxms/srv/auth/proto/auth"
)

type Auth struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// Auth.Call is called by the API as /auth/call with post body {"name": "foo"}
func (e *Auth) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Auth.Call request")

	// extract the client from the context
	authClient, ok := client.AuthFromContext(ctx)
	if !ok {
		return errors.InternalServerError("com.fxms.api.auth.call", "auth client not found")
	}

	// make request
	response, err := authClient.Call(ctx, &auth.Request{
		Name: extractValue(req.Post["name"]),
	})
	if err != nil {
		return errors.InternalServerError("com.fxms.api.auth.call", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}
