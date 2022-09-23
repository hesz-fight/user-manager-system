package v1

import (
	"learn/user-manager-system/httpsvr/global"
	"learn/user-manager-system/httpsvr/pkg/app"
	"learn/user-manager-system/httpsvr/pkg/errcode"
	"learn/user-manager-system/httpsvr/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func NewUserHnadler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) Login(c *gin.Context) {
	param := &service.LoginRequest{}
	response := app.NewResponse(c)
	if err := c.ShouldBindJSON(param); err != nil {
		global.LogLogger.Errorf("app.BindAndValid error: ", err)
		response.ToErrorResponse(errcode.InvalidParams.WithDetils(err.Error()))
	}
	svr := service.NewUserService(c.Request.Context())
	flag, err := svr.Login(param)
	if err != nil {
		global.LogLogger.Errorf("svr.Login error: ", err)
		response.ToErrorResponse(errcode.ErrorLoginFail.WithDetils(err.Error()))
	}
	if !flag {
		response.ToErrorResponse(errcode.ErrorLoginFail.WithDetils(err.Error()))
	}
	response.ToResponse(gin.H{})
}

func (u *UserHandler) UpdateNickname(c *gin.Context) {

}

func (u *UserHandler) UpdateProfile(c *gin.Context) {

}
