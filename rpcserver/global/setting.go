package global

import (
	"learn/user-manager-system/rpcsvr/pkg/logger"
	"learn/user-manager-system/rpcsvr/pkg/setting"
)

// 三个区段配置文件的全局变量
var (
	ServerSetting   *setting.ServerSetting
	AppSetting      *setting.AppSetting
	DatabaseSetting *setting.DatabaseSetting
	RedisSetting    *setting.RedisSetting
)

var LogLogger *logger.Logger
