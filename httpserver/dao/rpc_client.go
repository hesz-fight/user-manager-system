package dao

import (
	"learn/user-manager-system/httpsvr/proto"
)

type UserClient struct {
}

func NewUserClient() *UserClient {
	return &UserClient{}
}

func (u *UserClient) Login(username string, password string) (*proto.LoginResponseParam, error) {

	return nil, nil
}
