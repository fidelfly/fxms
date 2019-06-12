package token

import (
	"time"

	"gopkg.in/oauth2.v3"
)

type info struct {
	*TokenData
}

func NewTokenData(info oauth2.TokenInfo) *TokenData {
	return &TokenData{
		ClientId:         info.GetClientID(),
		UserId:           info.GetUserID(),
		RedirectUri:      info.GetRedirectURI(),
		Scope:            info.GetScope(),
		Code:             info.GetCode(),
		CodeCreateAt:     info.GetCodeCreateAt().Unix(),
		CodeExpiresIn:    int64(info.GetCodeExpiresIn()),
		Access:           info.GetAccess(),
		AccessCreateAt:   info.GetAccessCreateAt().Unix(),
		AccessExpiresIn:  int64(info.GetAccessExpiresIn()),
		Refresh:          info.GetRefresh(),
		RefreshCreateAt:  info.GetRefreshCreateAt().Unix(),
		RefreshExpiresIn: int64(info.GetRefreshExpiresIn()),
	}
}

// NewTokenInfoWrapper create to TokenInfoWrapper model instance
func NewTokenInfo(data *TokenData) *info {
	return &info{data}
}

// New create to TokenInfoWrapper model instance
func (t *info) New() oauth2.TokenInfo {
	return NewTokenInfo(t.TokenData)
}

// GetClientID the client id
func (t *info) GetClientID() string {
	return t.GetClientId()
}

// SetClientID the client id
func (t *info) SetClientID(clientID string) {
	t.ClientId = clientID
}

// GetUserID the user id
func (t *info) GetUserID() string {
	return t.GetUserId()
}

// SetUserID the user id
func (t *info) SetUserID(userID string) {
	t.UserId = userID
}

// GetRedirectURI redirect URI
func (t *info) GetRedirectURI() string {
	return t.GetRedirectUri()
}

// SetRedirectURI redirect URI
func (t *info) SetRedirectURI(redirectURI string) {
	t.RedirectUri = redirectURI
}

// GetScope get scope of authorization
func (t *info) GetScope() string {
	return t.Scope
}

// SetScope get scope of authorization
func (t *info) SetScope(scope string) {
	t.Scope = scope
}

// GetCode authorization code
func (t *info) GetCode() string {
	return t.Code
}

// SetCode authorization code
func (t *info) SetCode(code string) {
	t.Code = code
}

// GetCodeCreateAt create Time
func (t *info) GetCodeCreateAt() time.Time {
	return time.Unix(t.CodeCreateAt, 0)
}

// SetCodeCreateAt create Time
func (t *info) SetCodeCreateAt(createAt time.Time) {
	t.CodeCreateAt = createAt.Unix()
}

// GetCodeExpiresIn the lifetime in seconds of the authorization code
func (t *info) GetCodeExpiresIn() time.Duration {
	return time.Duration(t.CodeExpiresIn)
}

// SetCodeExpiresIn the lifetime in seconds of the authorization code
func (t *info) SetCodeExpiresIn(exp time.Duration) {
	t.CodeExpiresIn = int64(exp)
}

// GetAccess access TokenInfoWrapper
func (t *info) GetAccess() string {
	return t.Access
}

// SetAccess access TokenInfoWrapper
func (t *info) SetAccess(access string) {
	t.Access = access
}

// GetAccessCreateAt create Time
func (t *info) GetAccessCreateAt() time.Time {
	return time.Unix(t.AccessCreateAt, 0)
}

// SetAccessCreateAt create Time
func (t *info) SetAccessCreateAt(createAt time.Time) {
	t.AccessCreateAt = createAt.Unix()
}

// GetAccessExpiresIn the lifetime in seconds of the access TokenInfoWrapper
func (t *info) GetAccessExpiresIn() time.Duration {
	return time.Duration(t.AccessExpiresIn)
}

// SetAccessExpiresIn the lifetime in seconds of the access TokenInfoWrapper
func (t *info) SetAccessExpiresIn(exp time.Duration) {
	t.AccessExpiresIn = int64(exp)
}

// GetRefresh refresh TokenInfoWrapper
func (t *info) GetRefresh() string {
	return t.Refresh
}

// SetRefresh refresh TokenInfoWrapper
func (t *info) SetRefresh(refresh string) {
	t.Refresh = refresh
}

// GetRefreshCreateAt create Time
func (t *info) GetRefreshCreateAt() time.Time {
	return time.Unix(t.RefreshCreateAt, 0)
}

// SetRefreshCreateAt create Time
func (t *info) SetRefreshCreateAt(createAt time.Time) {
	t.RefreshCreateAt = createAt.Unix()
}

// GetRefreshExpiresIn the lifetime in seconds of the refresh TokenInfoWrapper
func (t *info) GetRefreshExpiresIn() time.Duration {
	return time.Duration(t.RefreshExpiresIn)
}

// SetRefreshExpiresIn the lifetime in seconds of the refresh TokenInfoWrapper
func (t *info) SetRefreshExpiresIn(exp time.Duration) {
	t.RefreshExpiresIn = int64(exp)
}
