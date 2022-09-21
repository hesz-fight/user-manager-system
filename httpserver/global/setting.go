package global

import "learn/user-manager-system/pkg/setting"

// 三个区段配置文件的全局变量
var (
	ServerSetting   *setting.ServerSetting
	AppSetting      *setting.AppSetting
	DatabaseSetting *setting.DatabaseSetting
)
