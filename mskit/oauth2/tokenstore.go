package oauth2

import (
	"github.com/go-xorm/xorm"
	"gopkg.in/oauth2.v3"

	"github.com/fidelfly/fxms/mskit/db"
)

type dbStore struct {
	engine *xorm.Engine
}

//export
func NewDbStore(instance *db.Instance) (store *dbStore, err error) {
	engine, err := db.InitEngine(instance)
	if err != nil {
		return
	}

	err = engine.Sync(&Token{})
	if err != nil {
		return
	}
	store = &dbStore{engine}
	return
}

// create and store the new token information
func (s *dbStore) Create(info oauth2.TokenInfo) error {
	token := Token{
		ClientID:         info.GetClientID(),
		UserID:           info.GetUserID(),
		RedirectURI:      info.GetRedirectURI(),
		Scope:            info.GetScope(),
		Code:             info.GetCode(),
		CodeCreateAt:     info.GetCodeCreateAt(),
		CodeExpiresIn:    info.GetCodeExpiresIn(),
		Access:           info.GetAccess(),
		AccessCreateAt:   info.GetAccessCreateAt(),
		AccessExpiresIn:  info.GetAccessExpiresIn(),
		Refresh:          info.GetRefresh(),
		RefreshCreateAt:  info.GetRefreshCreateAt(),
		RefreshExpiresIn: info.GetRefreshExpiresIn(),
	}

	if _, err := s.engine.Insert(&token); err != nil {
		return err
	}
	return nil
}

// delete the authorization code
func (s *dbStore) RemoveByCode(code string) error {
	if len(code) == 0 {
		return nil
	}
	token := Token{
		Code: code,
	}
	if _, err := s.engine.Delete(&token); err != nil {
		return err
	}
	return nil
}

// use the access token to delete the token information
func (s *dbStore) RemoveByAccess(access string) error {
	if len(access) == 0 {
		return nil
	}
	token := Token{
		Access: access,
	}
	if _, err := s.engine.Delete(&token); err != nil {
		return err
	}
	return nil
}

// use the refresh token to delete the token information
func (s *dbStore) RemoveByRefresh(refresh string) error {
	if len(refresh) == 0 {
		return nil
	}
	token := Token{
		Refresh: refresh,
	}
	if _, err := s.engine.Delete(&token); err != nil {
		return err
	}
	return nil
}

// use the authorization code for token information data
func (s *dbStore) GetByCode(code string) (oauth2.TokenInfo, error) {
	if len(code) == 0 {
		return nil, nil
	}
	token := Token{
		Code: code,
	}
	if find, err := s.engine.Get(&token); find {
		return &token, nil
	} else {
		return nil, err
	}
}

// use the access token for token information data
func (s *dbStore) GetByAccess(access string) (oauth2.TokenInfo, error) {
	if len(access) == 0 {
		return nil, nil
	}
	token := Token{
		Access: access,
	}
	if find, err := s.engine.Get(&token); find {
		return &token, nil
	} else {
		return nil, err
	}
}

// use the refresh token for token information data
func (s *dbStore) GetByRefresh(refresh string) (oauth2.TokenInfo, error) {
	if len(refresh) == 0 {
		return nil, nil
	}
	token := Token{
		Refresh: refresh,
	}
	if find, err := s.engine.Get(&token); find {
		return &token, nil
	} else {
		return nil, err
	}
}
