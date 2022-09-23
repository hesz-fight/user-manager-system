package main

import (
	"fmt"
	"log"
	"time"

	"learn/user-manager-system/rpcsvr/global"
	"learn/user-manager-system/rpcsvr/pkg/logger"
	"learn/user-manager-system/rpcsvr/pkg/setting"

	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"gopkg.in/natefinch/lumberjack.v2"
)

type rpcsvr struct {
}

func (h *rpcsvr) Init() {
	if err := initSetting(); err != nil {
		log.Fatal("init setting error: ", err)
	}
	if err := initDB(); err != nil {
		log.Fatal("init db error: ", err)
	}
	if err := initRedis(); err != nil {
		log.Fatal("init redis error: ", err)
	}
	if err := initLog(); err != nil {
		log.Fatal("init log error: ", err)
	}
}

func (h *rpcsvr) Run() {

}

func initSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	if err := s.ReadSection("Server", &global.ServerSetting); err != nil {
		return err
	}
	if err = s.ReadSection("App", &global.AppSetting); err != nil {
		return err
	}
	if err = s.ReadSection("Database", &global.DatabaseSetting); err != nil {
		return err
	}
	if err = s.ReadSection("Redis", &global.DatabaseSetting); err != nil {
		return err
	}

	return nil
}

func initDB() error {
	str := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	var err error
	global.DB, err = gorm.Open(global.DatabaseSetting.DBType,
		fmt.Sprintf(str,
			global.DatabaseSetting.UserName,
			global.DatabaseSetting.Password,
			global.DatabaseSetting.Host,
			global.DatabaseSetting.DBName,
			global.DatabaseSetting.Charset,
			global.DatabaseSetting.ParseTime,
		))
	if err != nil {
		return err
	}

	return nil
}

func initRedis() error {
	address := fmt.Sprintf("%s:%d", global.RedisSetting.Host, global.RedisSetting.Port)
	global.RedisPool = &redis.Pool{
		MaxIdle:     global.RedisSetting.MaxIdle,                                  // 最大空闲连接实例的数量
		MaxActive:   global.RedisSetting.MaxActive,                                // 表示和数据库的最大连接数，0表示没有限制
		IdleTimeout: time.Duration(global.RedisSetting.IdleTimeout) * time.Second, // 最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}

	return nil
}

func initLog() error {
	filepath := global.AppSetting.LogSavePath + "/" +
		global.AppSetting.LogFileName + global.AppSetting.LogFileExt

	writer := &lumberjack.Logger{
		Filename:  filepath,
		MaxSize:   600,
		MaxAge:    60,
		LocalTime: true,
	}
	global.LogLogger = logger.NerLogger(writer, "", log.LstdFlags)

	return nil
}

func main() {
	h := rpcsvr{}
	h.Init()
	h.Run()
}
