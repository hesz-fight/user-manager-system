package v1

import (
	"learn/user-manager-system/httpsvr/global"
	"learn/user-manager-system/httpsvr/model"
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
	req := &model.LoginHttpReq{}
	response := app.NewResponse(c)
	if err := c.ShouldBindJSON(req); err != nil {
		global.LogLogger.Errorf("app.BindAndValid login request error: ", err)
		response.ToErrorResponse(errcode.InvalidParams.WithDetils(err.Error()))
	}
	svr := service.NewUserService(c.Request.Context())
	rsp, cookie, err := svr.Login(req)
	if err != nil {
		global.LogLogger.Errorf("svr.Login error: ", err)
		response.ToErrorResponse(errcode.ErrorLoginFail.WithDetils(err.Error()))
	}
	c.SetCookie("token", cookie, 0, "/", "localhost", false, true)

	response.ToResponse(rsp)
}

func (u *UserHandler) UpdateNickname(c *gin.Context) {
	req := &model.UpdateNicknameReq{}
	response := app.NewResponse(c)
	if err := c.ShouldBindJSON(req); err != nil {
		global.LogLogger.Errorf("app.BindAndValid update nickname request error: ", err)
		response.ToErrorResponse(errcode.InvalidParams.WithDetils(err.Error()))
	}
	cookie, err := c.Cookie("token")
	if err != nil {
		return
	}

	svr := service.NewUserService(c.Request.Context())
	rsp, err := svr.UpdateNickname(req, cookie)
	if err != nil {
		global.LogLogger.Errorf("svr.UpdateNickname error: ", err)
		response.ToErrorResponse(errcode.ErrorLoginFail.WithDetils(err.Error()))
	}
	response.ToResponse(rsp)
}

func (u *UserHandler) UpdateProfile(c *gin.Context) {
	req := &model.UpdateProfileReq{}
	response := app.NewResponse(c)
	if err := c.ShouldBindJSON(req); err != nil {
		global.LogLogger.Errorf("app.BindAndValid update profile request error: ", err)
		response.ToErrorResponse(errcode.InvalidParams.WithDetils(err.Error()))
	}
	cookie, err := c.Cookie("token")
	if err != nil {
		return
	}

	svr := service.NewUserService(c.Request.Context())
	rsp, err := svr.UpdateProfile(req, cookie)
	if err != nil {
		global.LogLogger.Errorf("svr.UpdateProfile error: ", err)
		response.ToErrorResponse(errcode.ErrorLoginFail.WithDetils(err.Error()))
	}
	response.ToResponse(rsp)
}
