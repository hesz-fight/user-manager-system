package dao

import (
	"learn/user-manager-system/rpcsvr/model"

	"github.com/jinzhu/gorm"
)

type UserDao struct {
	engine *gorm.DB
}

func NewUserDao(engine *gorm.DB) *UserDao {
	return &UserDao{engine: engine}
}

func (u *UserDao) GetByUsername(username string, password string) (*model.User, error) {
	return nil, nil
}

func (u *UserDao) GetById() (*model.User, error) {
	return nil, nil
}

func (u *UserDao) Update(user *model.User) error {
	return nil
}
