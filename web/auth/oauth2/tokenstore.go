package oauth2

import (
	"context"

	"github.com/fidelfly/fxgo/authx"
	"github.com/fidelfly/fxgo/logx"
	"gopkg.in/oauth2.v3"

	"github.com/fidelfly/fxms/mskit/proto/base"
	"github.com/fidelfly/fxms/mskit/rpcc"
	"github.com/fidelfly/fxms/srv/auth/proto/token"
)

type tokenStore struct {
	client token.TokenService
	cache  oauth2.TokenStore
}

func NewTokenStore() *tokenStore {
	c := token.NewTokenService("com.fxms.srv.auth", rpcc.DefaultClient)
	return &tokenStore{client: c, cache: authx.NewMemoryTokenStore()}
}

// create and store the new token information
func (s *tokenStore) Create(info oauth2.TokenInfo) error {
	_, err := s.client.Create(context.Background(), token.NewTokenData(info))
	if err != nil {
		return err
	}
	logx.CaptureError(s.cache.Create(info))
	return nil
}

// delete the authorization code
func (s *tokenStore) RemoveByCode(code string) error {
	err := s.cache.RemoveByCode(code)
	if err != nil {
		return err
	}
	_, err = s.client.RemoveByCode(context.Background(), &base.StringValue{Value: code})
	return err
}

// use the access token to delete the token information
func (s *tokenStore) RemoveByAccess(access string) error {
	err := s.cache.RemoveByAccess(access)
	if err != nil {
		return err
	}
	_, err = s.client.RemoveByAccess(context.Background(), &base.StringValue{Value: access})
	return err
}

// use the refresh token to delete the token information
func (s *tokenStore) RemoveByRefresh(refresh string) error {
	err := s.cache.RemoveByRefresh(refresh)
	if err != nil {
		return err
	}
	_, err = s.client.RemoveByRefresh(context.Background(), &base.StringValue{Value: refresh})
	return err
}

// use the authorization code for token information data
func (s *tokenStore) GetByCode(code string) (oauth2.TokenInfo, error) {
	ti, err := s.cache.GetByCode(code)
	if err == nil && ti != nil {
		return ti, nil
	}
	res, err := s.client.GetByCode(context.Background(), &base.StringValue{Value: code})
	if err != nil {
		return nil, err
	}
	return token.NewTokenInfo(res), nil
}

// use the access token for token information data
func (s *tokenStore) GetByAccess(access string) (oauth2.TokenInfo, error) {
	ti, err := s.cache.GetByAccess(access)
	if err == nil && ti != nil {
		return ti, nil
	}
	res, err := s.client.GetByAccess(context.Background(), &base.StringValue{Value: access})
	if err != nil {
		return nil, err
	}
	return token.NewTokenInfo(res), nil
}

// use the refresh token for token information data
func (s *tokenStore) GetByRefresh(refresh string) (oauth2.TokenInfo, error) {
	ti, err := s.cache.GetByRefresh(refresh)
	if err == nil && ti != nil {
		return ti, nil
	}
	res, err := s.client.GetByRefresh(context.Background(), &base.StringValue{Value: refresh})
	if err != nil {
		return nil, err
	}
	return token.NewTokenInfo(res), nil
}
