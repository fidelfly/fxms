package oauth2

import (
	"context"
	"errors"

	"github.com/micro/go-micro/client"
	"gopkg.in/oauth2.v3"

	"github.com/fidelfly/fxms/srv/auth/proto/token"
)

type tokenStore struct {
	client token.TokenService
}

func NewTokenStore() *tokenStore {
	client := token.NewTokenService("com.fxms.srv.auth", client.DefaultClient)
	return &tokenStore{client: client}
}

// create and store the new token information
func (s *tokenStore) Create(info oauth2.TokenInfo) error {
	resp, err := s.client.Create(context.Background(), &token.TokenRequest{Token: token.NewTokenData(info)})
	if err != nil {
		return err
	}
	if len(resp.ErrMessage) > 0 {
		return errors.New(resp.ErrMessage)
	}
	return nil
}

// delete the authorization code
func (s *tokenStore) RemoveByCode(code string) error {
	res, err := s.client.RemoveByCode(context.Background(), &token.StringValue{Value: code})
	if err != nil {
		return err
	}
	if len(res.ErrMessage) > 0 {
		return errors.New(res.ErrMessage)
	}
	return nil
}

// use the access token to delete the token information
func (s *tokenStore) RemoveByAccess(access string) error {
	res, err := s.client.RemoveByAccess(context.Background(), &token.StringValue{Value: access})
	if err != nil {
		return err
	}
	if len(res.ErrMessage) > 0 {
		return errors.New(res.ErrMessage)
	}
	return nil
}

// use the refresh token to delete the token information
func (s *tokenStore) RemoveByRefresh(refresh string) error {
	res, err := s.client.RemoveByRefresh(context.Background(), &token.StringValue{Value: refresh})
	if err != nil {
		return err
	}
	if len(res.ErrMessage) > 0 {
		return errors.New(res.ErrMessage)
	}
	return nil
}

// use the authorization code for token information data
func (s *tokenStore) GetByCode(code string) (oauth2.TokenInfo, error) {
	res, err := s.client.GetByCode(context.Background(), &token.StringValue{Value: code})
	if err != nil {
		return nil, err
	}
	if len(res.ErrMessage) > 0 {
		return nil, errors.New(res.ErrMessage)
	}
	return token.NewTokenInfo(res.Token), nil
}

// use the access token for token information data
func (s *tokenStore) GetByAccess(access string) (oauth2.TokenInfo, error) {
	res, err := s.client.GetByAccess(context.Background(), &token.StringValue{Value: access})
	if err != nil {
		return nil, err
	}
	if len(res.ErrMessage) > 0 {
		return nil, errors.New(res.ErrMessage)
	}
	return token.NewTokenInfo(res.Token), nil
}

// use the refresh token for token information data
func (s *tokenStore) GetByRefresh(refresh string) (oauth2.TokenInfo, error) {
	res, err := s.client.GetByRefresh(context.Background(), &token.StringValue{Value: refresh})
	if err != nil {
		return nil, err
	}
	if len(res.ErrMessage) > 0 {
		return nil, errors.New(res.ErrMessage)
	}
	return token.NewTokenInfo(res.Token), nil
}
