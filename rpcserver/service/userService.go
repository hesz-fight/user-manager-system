package service

import (
	"fmt"
	"learn/user-manager-system/rpcsvr/proto"
)

type UserService struct {
}

func (u *UserService) Login(req *proto.LoginRequestParam, rsp *proto.LoginResponseParam) error {
	fmt.Println("login ...")
	return nil
}
