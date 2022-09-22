package router

import (
	"learn/user-manager-system/global"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	engine := gin.New()

	if global.ServerSetting.RunMode == "degun" {
		engine.Use(gin.Logger())
		engine.Use(gin.Recovery())
	}

	engine.Group("api/v1")
	{
		engine.POST("/user/login")
		engine.POST("/user/update/nickname")
		engine.POST("/user/update/profile")
	}

	return engine
}
