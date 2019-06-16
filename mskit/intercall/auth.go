package intercall

import (
	"net/http"

	"github.com/fidelfly/fxgo"
	"github.com/fidelfly/fxgo/authx"
	"github.com/fidelfly/fxgo/httprxr"

	"github.com/fidelfly/fxms/mspkg/rpcc"

	"github.com/fidelfly/fxms/srv/auth/proto/token"
)

var tokenSrv = token.NewTokenService("com.fxms.srv.auth", rpcc.DefaultClient)

//export
func AuthServerFilter(w http.ResponseWriter, r *http.Request, next http.Handler) {
	if authCode, ok := httprxr.BearerAuth(r); ok {
		rsp, err := tokenSrv.Validate(r.Context(), &token.TokenRequest{AccessToken: authCode})
		if err != nil {
			httprxr.ResponseJSON(w, http.StatusInternalServerError, httprxr.ExceptionMessage(err))
			return
		} else if rsp.Err != nil {
			httprxr.ResponseJSON(w, http.StatusUnauthorized, httprxr.ErrorMessage(rsp.Err))
			return
		} else if rsp.TokenInfo != nil {
			r = httprxr.ContextSet(r, fxgo.ContextUserKey, rsp.TokenInfo.GetUserId(), fxgo.ContextTokenKey, token.NewTokenInfo(rsp.TokenInfo))
		}

		next.ServeHTTP(w, r)
	} else {
		httprxr.ResponseJSON(w, http.StatusInternalServerError, httprxr.NewErrorMessage(authx.UnauthorizedErrorCode, "invalid access"))
	}
}
