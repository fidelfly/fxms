package user

import (
	"github.com/fidelfly/fxms/mspkg/proto"
	"github.com/fidelfly/fxms/srv/user/res"
)

func NewResource(data *UserData, skipPwd ...bool) *res.User {
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
	if len(skipPwd) > 0 && skipPwd[0] {
		user.Password = ""
	}
	return user
}

func NewData(user *res.User, skipPwd ...bool) *UserData {
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

	if len(skipPwd) > 0 && skipPwd[0] {
		data.Password = ""
	}

	return data
}

func (ud *UserData) FillData(user *res.User, skipPwd ...bool) {
	if user == nil {
		return
	}
	ud.Id = user.Id
	ud.Code = user.Code
	ud.Name = user.Name
	ud.Email = user.Email
	ud.Avatar = user.Avatar
	ud.Password = user.Password
	ud.CreateTime = proto.ToProtoTime(user.CreateTime)
	ud.UpdateTime = proto.ToProtoTime(user.UpdateTime)

	if len(skipPwd) > 0 && skipPwd[0] {
		ud.Password = ""
	}

	return
}
