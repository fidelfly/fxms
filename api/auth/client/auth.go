package client

import (
	"context"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"

	"github.com/fidelfly/fxms/srv/auth/proto/auth"
)

type authKey struct{}

// FromContext retrieves the client from the Context
func AuthFromContext(ctx context.Context) (auth.AuthService, bool) {
	c, ok := ctx.Value(authKey{}).(auth.AuthService)
	return c, ok
}

// Client returns a wrapper for the AuthClient
func AuthWrapper(service micro.Service) server.HandlerWrapper {
	client := auth.NewAuthService("com.fxms.srv.auth", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, authKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
