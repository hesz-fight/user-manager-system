package setting

import "time"

type ServerSetting struct {
	RunMode      string
	HTTPHost     string
	HTTPPort     string
	RPCHost      string
	RPCPort      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSetting struct {
	DefaultPageSize      int
	MaxPageSize          int
	LogSavePath          string
	LogFileName          string
	LogFileExt           string
	UploadSavePath       string
	UploadServerUrl      string
	UploadImageMaxSize   int
	UploadImageAllowExts []string
}
