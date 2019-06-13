package user

import (
	"github.com/fidelfly/fxms/mskit/proto"
	"github.com/fidelfly/fxms/srv/user/res"
)

func NewResource(data *UserData) *res.User {
	if data == nil {
		return nil
	}
	user := &res.User{
		Id:         data.GetId(),
		Code:       data.GetCode(),
		Name:       data.GetName(),
		Password:   data.GetPassword(),
		Avatar:     data.GetAvatar(),
		UpdateTime: proto.ToTime(data.UpdateTime),
		CreateTime: proto.ToTime(data.CreateTime),
	}
	return user
}

func NewData(user *res.User) *UserData {
	if user == nil {
		return nil
	}
	data := &UserData{
		Id:         user.Id,
		Code:       user.Code,
		Name:       user.Name,
		Email:      user.Email,
		Password:   user.Password,
		Avatar:     user.Avatar,
		CreateTime: proto.ToProtoTime(user.CreateTime),
		UpdateTime: proto.ToProtoTime(user.UpdateTime),
	}

	return data
}
