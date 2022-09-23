package router

import (
	"learn/user-manager-system/httpsvr/global"
	v1 "learn/user-manager-system/httpsvr/router/v1"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	engine := gin.New()

	if global.ServerSetting.RunMode == "degun" {
		engine.Use(gin.Logger())
		engine.Use(gin.Recovery())
	}

	userHandler := v1.NewUserHnadler()
	engine.Group("api/v1")
	{
		engine.POST("/user/login", userHandler.Login)
		engine.POST("/user/update/nickname", userHandler.UpdateNickname)
		engine.POST("/user/update/profile", userHandler.UpdateProfile)
	}

	return engine
}
