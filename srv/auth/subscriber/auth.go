package subscriber

import (
	"context"

	"github.com/micro/go-micro/util/log"

	"github.com/fidelfly/fxms/srv/auth/proto/auth"
)

type Auth struct{}

func (e *Auth) Handle(ctx context.Context, msg *auth.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *auth.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
