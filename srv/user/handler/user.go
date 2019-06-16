package handler

import (
	"context"
	"errors"

	"github.com/micro/go-micro/util/log"

	"github.com/fidelfly/fxms/mspkg/db"
	"github.com/fidelfly/fxms/mspkg/proto/api"
	"github.com/fidelfly/fxms/srv/user/proto/user"
	"github.com/fidelfly/fxms/srv/user/res"
)

type User struct{}

func (e *User) Delete(ctx context.Context, req *user.DeleteRequest, rsp *api.IdResponse) error {
	if req.Id <= 0 {
		return errors.New("invalid value of id")
	}

	data := &res.User{Id: req.Id}
	if id, err := db.Delete(data, db.ID(req.Id)); err != nil {
		return err
	} else {
		rsp.Id = id
	}
	return nil
}

func (e *User) Read(ctx context.Context, req *user.ReadRequest, rsp *user.UserData) error {
	if req.Id <= 0 {
		return errors.New("invalid value of id")
	}

	data := &res.User{Id: req.Id}
	if _, err := db.Read(data); err != nil {
		return err
	}
	rsp = user.NewData(data)
	return nil
}

func (e *User) Update(ctx context.Context, req *user.UpdateRequest, rsp *api.IdResponse) error {
	if req.Data == nil {
		return errors.New("data is empty")
	}
	id, err := db.Update(user.NewResource(req.Data))
	if err != nil {
		return err
	}
	rsp.Id = id
	return nil
}

func (e *User) Create(ctx context.Context, req *user.CreateRequest, rsp *api.IdResponse) error {
	if req.Data == nil {
		return errors.New("data is empty")
	}
	id, err := db.Create(user.NewResource(req.Data))
	if err != nil {
		return err
	}
	rsp.Id = id
	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (e *User) Call(ctx context.Context, req *user.Request, rsp *user.Response) error {
	log.Log("Received User.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *User) Stream(ctx context.Context, req *user.StreamingRequest, stream user.User_StreamStream) error {
	log.Logf("Received User.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&user.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *User) PingPong(ctx context.Context, stream user.User_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&user.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
