package service

import (
	"context"
	"learn/user-manager-system/httpsvr/dao"
	"learn/user-manager-system/httpsvr/model"
	"learn/user-manager-system/httpsvr/pkg/errcode"
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
	if rsp.ErrCode != 0 {
		return nil, "", errcode.ErrorLoginFail
	}
	loginHttpRsp := &model.LoginHttpRsp{
		UserInfo: &model.UserInfo{
			Username: rsp.UserInfo.Username,
			Nickname: rsp.UserInfo.Nickname,
			Profile:  rsp.UserInfo.Profile,
		},
	}

	return loginHttpRsp, rsp.Cookie, nil
}

func (u *UserServie) UpdateNickname(reqParam *model.UpdateNicknameReq, cookie string) (*model.UpdateNicknameRsp, error) {
	rsp, err := u.UserClient.UpdateNickname(reqParam.Nickname, cookie)
	if err != nil {
		return nil, nil
	}
	if rsp.ErrCode != 0 {
		return nil, errcode.ErrorUpdateUserFail
	}

	return &model.UpdateNicknameRsp{}, nil
}

func (u *UserServie) UpdateProfile(reqParam *model.UpdateProfileReq, cookie string) (*model.UpdateProfileRsp, error) {
	rsp, err := u.UserClient.UpdateProfile(reqParam.Profile, cookie)
	if err != nil {
		return nil, nil
	}
	if rsp.ErrCode != 0 {
		return nil, errcode.ErrorUpdateUserFail
	}

	return &model.UpdateProfileRsp{}, nil
}
