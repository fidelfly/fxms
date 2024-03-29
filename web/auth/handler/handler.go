package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/fidelfly/fxgo/logx"

	"github.com/fidelfly/fxms/mspkg/rpcc"
	"github.com/fidelfly/fxms/mspkg/wbr"
	"github.com/fidelfly/fxms/srv/auth/proto/auth"
)

func AuthCall(w http.ResponseWriter, r *http.Request) {
	logx.Debug("Calling Web Auth call....")
	logx.Infof("TraceID = %s", wbr.GetTraceID(r))
	logx.Infof("Call From Service = %s", wbr.GetFromService(r))
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	authClient := auth.NewAuthService("com.fxms.srv.auth", rpcc.DefaultClient)
	rsp, err := authClient.Call(context.WithValue(r.Context(), "fidel", "fidelValue"), &auth.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
