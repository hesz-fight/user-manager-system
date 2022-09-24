package router

import (
	"fmt"
	"learn/user-manager-system/rpcsvr/proto"
)

type UserHandler struct {
}

func (u *UserHandler) Login(req *proto.LoginRequestParam, rsp *proto.LoginResponseParam) error {
	fmt.Println("login ...")
	return nil
}
