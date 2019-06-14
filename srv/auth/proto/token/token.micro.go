// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: token.proto

package token

import (
	fmt "fmt"
	_ "github.com/fidelfly/fxms/mskit/proto/api"
	base "github.com/fidelfly/fxms/mskit/proto/base"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Token service

type TokenService interface {
	Create(ctx context.Context, in *TokenData, opts ...client.CallOption) (*empty.Empty, error)
	RemoveByCode(ctx context.Context, in *base.StringValue, opts ...client.CallOption) (*empty.Empty, error)
	RemoveByAccess(ctx context.Context, in *base.StringValue, opts ...client.CallOption) (*empty.Empty, error)
	RemoveByRefresh(ctx context.Context, in *base.StringValue, opts ...client.CallOption) (*empty.Empty, error)
	GetByCode(ctx context.Context, in *base.StringValue, opts ...client.CallOption) (*TokenData, error)
	GetByAccess(ctx context.Context, in *base.StringValue, opts ...client.CallOption) (*TokenData, error)
	GetByRefresh(ctx context.Context, in *base.StringValue, opts ...client.CallOption) (*TokenData, error)
}

type tokenService struct {
	c    client.Client
	name string
}

func NewTokenService(name string, c client.Client) TokenService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "com.fxms.srv.auth"
	}
	return &tokenService{
		c:    c,
		name: name,
	}
}

func (c *tokenService) Create(ctx context.Context, in *TokenData, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "Token.Create", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenService) RemoveByCode(ctx context.Context, in *base.StringValue, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "Token.RemoveByCode", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenService) RemoveByAccess(ctx context.Context, in *base.StringValue, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "Token.RemoveByAccess", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenService) RemoveByRefresh(ctx context.Context, in *base.StringValue, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "Token.RemoveByRefresh", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenService) GetByCode(ctx context.Context, in *base.StringValue, opts ...client.CallOption) (*TokenData, error) {
	req := c.c.NewRequest(c.name, "Token.GetByCode", in)
	out := new(TokenData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenService) GetByAccess(ctx context.Context, in *base.StringValue, opts ...client.CallOption) (*TokenData, error) {
	req := c.c.NewRequest(c.name, "Token.GetByAccess", in)
	out := new(TokenData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenService) GetByRefresh(ctx context.Context, in *base.StringValue, opts ...client.CallOption) (*TokenData, error) {
	req := c.c.NewRequest(c.name, "Token.GetByRefresh", in)
	out := new(TokenData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Token service

type TokenHandler interface {
	Create(context.Context, *TokenData, *empty.Empty) error
	RemoveByCode(context.Context, *base.StringValue, *empty.Empty) error
	RemoveByAccess(context.Context, *base.StringValue, *empty.Empty) error
	RemoveByRefresh(context.Context, *base.StringValue, *empty.Empty) error
	GetByCode(context.Context, *base.StringValue, *TokenData) error
	GetByAccess(context.Context, *base.StringValue, *TokenData) error
	GetByRefresh(context.Context, *base.StringValue, *TokenData) error
}

func RegisterTokenHandler(s server.Server, hdlr TokenHandler, opts ...server.HandlerOption) error {
	type token interface {
		Create(ctx context.Context, in *TokenData, out *empty.Empty) error
		RemoveByCode(ctx context.Context, in *base.StringValue, out *empty.Empty) error
		RemoveByAccess(ctx context.Context, in *base.StringValue, out *empty.Empty) error
		RemoveByRefresh(ctx context.Context, in *base.StringValue, out *empty.Empty) error
		GetByCode(ctx context.Context, in *base.StringValue, out *TokenData) error
		GetByAccess(ctx context.Context, in *base.StringValue, out *TokenData) error
		GetByRefresh(ctx context.Context, in *base.StringValue, out *TokenData) error
	}
	type Token struct {
		token
	}
	h := &tokenHandler{hdlr}
	return s.Handle(s.NewHandler(&Token{h}, opts...))
}

type tokenHandler struct {
	TokenHandler
}

func (h *tokenHandler) Create(ctx context.Context, in *TokenData, out *empty.Empty) error {
	return h.TokenHandler.Create(ctx, in, out)
}

func (h *tokenHandler) RemoveByCode(ctx context.Context, in *base.StringValue, out *empty.Empty) error {
	return h.TokenHandler.RemoveByCode(ctx, in, out)
}

func (h *tokenHandler) RemoveByAccess(ctx context.Context, in *base.StringValue, out *empty.Empty) error {
	return h.TokenHandler.RemoveByAccess(ctx, in, out)
}

func (h *tokenHandler) RemoveByRefresh(ctx context.Context, in *base.StringValue, out *empty.Empty) error {
	return h.TokenHandler.RemoveByRefresh(ctx, in, out)
}

func (h *tokenHandler) GetByCode(ctx context.Context, in *base.StringValue, out *TokenData) error {
	return h.TokenHandler.GetByCode(ctx, in, out)
}

func (h *tokenHandler) GetByAccess(ctx context.Context, in *base.StringValue, out *TokenData) error {
	return h.TokenHandler.GetByAccess(ctx, in, out)
}

func (h *tokenHandler) GetByRefresh(ctx context.Context, in *base.StringValue, out *TokenData) error {
	return h.TokenHandler.GetByRefresh(ctx, in, out)
}
