package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	_ "homework/docs"
	"homework/global"
	routers "homework/internal/apiserver/api"
	"homework/pkg/logger"
	"homework/pkg/setting"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

// @title Swagger Example API
// @version 1.0

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:80
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		//Addr:           ":" + global.ServerSetting.HttpPort,
		Addr:           ":80",
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout * time.Second,
		WriteTimeout:   global.ServerSetting.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() error {
	setting1, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting1.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting1.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
