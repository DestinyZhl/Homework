package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	_ "homework/docs"
	"homework/global"
	routers "homework/internal/apiserver/api"
	"homework/pkg/logger"
	"homework/pkg/setting"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	sigtermChan := make(chan os.Signal, 1)
	signal.Notify(sigtermChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		s.ListenAndServe()
	}()

	global.Logger.Info("Http Server Started.")
	<-sigtermChan
	global.Logger.Info("Http Server Stopped.")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := s.Shutdown(ctx); err != nil {
		global.Logger.Fatal("Http Server Shutdown Failed:%+v", err)
	}
	global.Logger.Info("Http Server Exited Properly")

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
