package setting

import "time"

type ServerSetting struct {
	Host         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSetting struct {
	LogSavePath          string
	LogFileName          string
	LogFileExt           string
	UploadSavePath       string
	UploadServerUrl      string
	UploadImageMaxSize   int
	UploadImageAllowExts []string
}

type DatabaseSetting struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type RedisSetting struct {
	Host        string
	Port        int
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}
