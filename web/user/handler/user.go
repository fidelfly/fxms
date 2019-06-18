package handler

import (
	"net/http"
	"strconv"

	"github.com/fidelfly/fxgo/httprxr"
	"github.com/fidelfly/fxgo/routex"

	"github.com/fidelfly/fxms/mskit"
	"github.com/fidelfly/fxms/mspkg/proto/api"
	"github.com/fidelfly/fxms/mspkg/rpcc"
	"github.com/fidelfly/fxms/srv/user/proto/user"
)

var UserHandler *routex.Router

var userServer user.UserService

func init() {
	userServer = user.NewUserService("com.fxms.srv.user", rpcc.DefaultClient)

	rr := mskit.NewWebRouter("/user")
	rr.Methods(http.MethodGet).HandlerFunc(GetUser)

	rr.Methods(http.MethodGet).Path("/{id}").HandlerFunc(GetUser)

	UserHandler = rr
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := httprxr.GetRequestVars(r, "id")
	userID, _ := strconv.ParseInt(params["id"], 10, 64)
	if userID == 0 {
		user := mskit.GetUserInfo(r)
		if user != nil {
			userID = user.Id
		}
	}
	if userID > 0 {
		rsp, err := userServer.Read(r.Context(), &api.IdResponse{Id: userID})
		if err != nil {
			httprxr.ResponseJSON(w, http.StatusInternalServerError, httprxr.ExceptionMessage(err))
			return
		}
		if rsp.Id == 0 {
			httprxr.ResponseJSON(w, http.StatusNotFound, nil)
			return
		}
		httprxr.ResponseJSON(w, http.StatusOK, user.NewResource(rsp))
		return
	}
	httprxr.ResponseJSON(w, http.StatusBadRequest, httprxr.InvalidParamError("userID"))
	return
}
