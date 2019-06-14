package handler

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/fidelfly/fxms/mskit/proto/base"
	"github.com/fidelfly/fxms/srv/auth/oauth2"
	"github.com/fidelfly/fxms/srv/auth/proto/token"
)

type Token struct{}

func (t *Token) Create(ctx context.Context, req *token.TokenData, resp *empty.Empty) error {
	if req != nil {
		err := oauth2.GetTokenStore().Create(token.NewTokenInfo(req))
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *Token) RemoveByCode(ctx context.Context, req *base.StringValue, res *empty.Empty) error {
	code := req.Value
	if len(code) > 0 {
		err := oauth2.GetTokenStore().RemoveByCode(code)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("code is empty")
}

func (t *Token) RemoveByAccess(ctx context.Context, req *base.StringValue, res *empty.Empty) error {
	access := req.Value
	if len(access) > 0 {
		err := oauth2.GetTokenStore().RemoveByAccess(access)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("access is empty")
}
func (t *Token) RemoveByRefresh(ctx context.Context, req *base.StringValue, res *empty.Empty) error {
	refresh := req.Value
	if len(refresh) > 0 {
		err := oauth2.GetTokenStore().RemoveByRefresh(refresh)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("refresh is empty")
}
func (t *Token) GetByCode(ctx context.Context, req *base.StringValue, res *token.TokenData) error {
	code := req.Value
	if len(code) > 0 {
		info, err := oauth2.GetTokenStore().GetByCode(code)
		if err != nil {
			return err
		}
		res = token.NewTokenData(info)
		return nil
	}
	return errors.New("code is empty")
}
func (t *Token) GetByAccess(ctx context.Context, req *base.StringValue, res *token.TokenData) error {
	access := req.Value
	if len(access) > 0 {
		info, err := oauth2.GetTokenStore().GetByAccess(access)
		if err != nil {
			return err
		}
		res = token.NewTokenData(info)
		return nil
	}
	return errors.New("access is empty")
}
func (t *Token) GetByRefresh(ctx context.Context, req *base.StringValue, res *token.TokenData) error {
	refresh := req.Value
	if len(refresh) > 0 {
		info, err := oauth2.GetTokenStore().GetByRefresh(refresh)
		if err != nil {
			return err
		}
		res = token.NewTokenData(info)
		return nil
	}
	return errors.New("refresh is empty")
}
