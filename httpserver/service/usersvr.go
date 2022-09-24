package service

import (
	"context"
	"learn/user-manager-system/httpsvr/dao"
	"learn/user-manager-system/httpsvr/model"
)

type UserServie struct {
	ctx        context.Context
	UserClient *dao.UserClient
}

func NewUserService(ctx context.Context) *UserServie {
	return &UserServie{
		ctx:        ctx,
		UserClient: dao.NewUserClient(),
	}
}

// convert to model user
func (u *UserServie) Login(reqParam *model.LoginHttpReq) (*model.LoginHttpRsp, string, error) {
	rsp, err := u.UserClient.Login(reqParam.Username, reqParam.Password)
	if err != nil {
		return nil, "", nil
	}
	loginHttpRsp := &model.LoginHttpRsp{
		Username: rsp.UserInfo.Username,
		Nickname: rsp.UserInfo.Nickname,
		Profile:  rsp.UserInfo.Profile,
	}

	return loginHttpRsp, rsp.Cookie, nil
}
