package oauth2

import (
	"context"

	"github.com/fidelfly/fxgo/authx"
	"gopkg.in/oauth2.v3"

	"github.com/fidelfly/fxms/mskit/proto/base"
	"github.com/fidelfly/fxms/mskit/rpcc"
	"github.com/fidelfly/fxms/srv/auth/proto/token"
)

type TokenSrvStore struct {
	client token.TokenService
}

func NewTokenStore() oauth2.TokenStore {
	c := token.NewTokenService("com.fxms.srv.auth", rpcc.DefaultClient)
	return authx.NewMultiLevelTokenStore(authx.NewMemoryTokenStore(), &TokenSrvStore{client: c})
}

// create and store the new token information
func (s *TokenSrvStore) Create(info oauth2.TokenInfo) error {
	_, err := s.client.Create(context.Background(), token.NewTokenData(info))
	return err
}

// delete the authorization code
func (s *TokenSrvStore) RemoveByCode(code string) error {
	_, err := s.client.RemoveByCode(context.Background(), &base.StringValue{Value: code})
	return err
}

// use the access token to delete the token information
func (s *TokenSrvStore) RemoveByAccess(access string) error {
	_, err := s.client.RemoveByAccess(context.Background(), &base.StringValue{Value: access})
	return err
}

// use the refresh token to delete the token information
func (s *TokenSrvStore) RemoveByRefresh(refresh string) error {
	_, err := s.client.RemoveByRefresh(context.Background(), &base.StringValue{Value: refresh})
	return err
}

// use the authorization code for token information data
func (s *TokenSrvStore) GetByCode(code string) (oauth2.TokenInfo, error) {
	res, err := s.client.GetByCode(context.Background(), &base.StringValue{Value: code})
	if err != nil {
		return nil, err
	}
	ti := token.NewTokenInfo(res)
	return ti, nil
}

// use the access token for token information data
func (s *TokenSrvStore) GetByAccess(access string) (oauth2.TokenInfo, error) {
	res, err := s.client.GetByAccess(context.Background(), &base.StringValue{Value: access})
	if err != nil {
		return nil, err
	}
	ti := token.NewTokenInfo(res)
	return ti, nil
}

// use the refresh token for token information data
func (s *TokenSrvStore) GetByRefresh(refresh string) (oauth2.TokenInfo, error) {
	res, err := s.client.GetByRefresh(context.Background(), &base.StringValue{Value: refresh})
	if err != nil {
		return nil, err
	}
	ti := token.NewTokenInfo(res)
	return ti, nil
}
