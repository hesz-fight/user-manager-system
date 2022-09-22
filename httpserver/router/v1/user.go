package v1

import "github.com/gin-gonic/gin"

type UserHandler struct{}

func NewUserHnadler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) Login(c *gin.Context) {

}

func (u *UserHandler) UpdateNickname(c *gin.Context) {

}

func (u *UserHandler) UpdateProfile(c *gin.Context) {

}
