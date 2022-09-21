package main

import (
	"learn/user-manager-system/global"
	"learn/user-manager-system/pkg/setting"
	"log"
)

type httpsvr struct {
}

func (h *httpsvr) Init() {
	if err := initSetting(); err != nil {
		log.Fatal("init setting error: ", err)
	}
	if err := initDB(); err != nil {
		log.Fatal("init setting error: ", err)
	}
}

func (h *httpsvr) Run() {

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

	return nil
}

func initDB() error {
	return nil
}

func initRedis() error {
	return nil
}

func initLog() error {
	return nil
}
