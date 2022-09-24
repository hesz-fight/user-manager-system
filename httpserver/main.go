package main

import (
	"log"
	"net/http"
	"time"

	"learn/user-manager-system/httpsvr/global"
	"learn/user-manager-system/httpsvr/pkg/logger"
	"learn/user-manager-system/httpsvr/pkg/setting"
	"learn/user-manager-system/httpsvr/pkg/simrpc"
	"learn/user-manager-system/httpsvr/router"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

type httpsvr struct {
}

func (h *httpsvr) Init() {
	if err := initSetting(); err != nil {
		log.Fatal("init setting error: ", err)
	}
	if err := initLog(); err != nil {
		log.Fatal("init log error: ", err)
	}
	if err := initClientPool(); err != nil {
		log.Fatal("init client error: ", err)
	}
}

func (h *httpsvr) Run() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := router.NewRouter()
	server := &http.Server{
		Addr:           global.ServerSetting.HTTPPort + ":" + global.ServerSetting.HTTPPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout * time.Second,
		WriteTimeout:   global.ServerSetting.WriteTimeout * time.Second,
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

func initClientPool() error {
	var err error
	address := global.ServerSetting.RPCHost + ":" + global.ServerSetting.RPCPort
	global.ClientPool, err = simrpc.NewClientPool(100, "tcp", address)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	h := httpsvr{}
	h.Init()
	h.Run()
}
