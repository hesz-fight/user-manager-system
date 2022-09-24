package dao

import (
	"learn/user-manager-system/httpsvr/global"
	"learn/user-manager-system/httpsvr/proto"
)

type UserClient struct {
}

func NewUserClient() *UserClient {
	return &UserClient{}
}

func (u *UserClient) Login(username string, password string) (*proto.LoginResponseParam, error) {
	req := &proto.LoginRequestParam{}
	rsp := &proto.LoginResponseParam{}
	err := global.ClientPool.Call("UserService.Login", req, rsp)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func (u *UserClient) UpdateNickname(username string, password string) (*proto.UpdateNicknameResponseParam, error) {
	req := &proto.UpdateNicknameRequestParam{}
	rsp := &proto.UpdateNicknameResponseParam{}
	if err := global.ClientPool.Call("UserService.UpdateNickname", req, rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}

func (u *UserClient) UpdateProfile(username string, password string) (*proto.LoginResponseParam, error) {
	req := &proto.LoginRequestParam{}
	rsp := &proto.LoginResponseParam{}
	if err := global.ClientPool.Call("UserService.UpdateProfile", req, rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}
