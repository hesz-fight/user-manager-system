package global

import (
	"learn/user-manager-system/httpsvr/pkg/logger"
	"learn/user-manager-system/httpsvr/pkg/setting"
	"learn/user-manager-system/httpsvr/pkg/simrpc"
)

// 三个区段配置文件的全局变量
var (
	ServerSetting *setting.ServerSetting
	AppSetting    *setting.AppSetting
)

var LogLogger *logger.Logger

var ClientPool *simrpc.ClientPool
