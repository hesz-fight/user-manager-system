package service

import (
	"context"
	"learn/user-manager-system/httpsvr/dao"
	"learn/user-manager-system/httpsvr/global"
)

type UserServie struct {
	ctx     context.Context
	userDao *dao.UserDao
}

func NewUserService(ctx context.Context) *UserServie {
	userDao := dao.NewUserDao(global.DBEngine)
	return &UserServie{
		ctx:     ctx,
		userDao: userDao,
	}
}

func (u *UserServie) Login(param *LoginRequest) (bool, error) {
	user, err := u.userDao.GetByUsername(param.Username, param.Password)
	if err != nil {
		return false, err
	}

	if param.Password != user.Password {
		return false, err
	}

	return true, nil
}
