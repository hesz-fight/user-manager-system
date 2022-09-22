package main

import (
	"log"
	"net/http"
	"time"

	"learn/user-manager-system/global"
	"learn/user-manager-system/pkg/setting"
	"learn/user-manager-system/router"

	"github.com/gin-gonic/gin"
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
	gin.SetMode(global.ServerSetting.RunMode)
	router := router.NewRouter()
	server := &http.Server{
		Addr:           "127.0.0.1" + ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
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
