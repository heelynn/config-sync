package main

import (
	"config-sync/internal/backends/nacos-config"
	"config-sync/internal/properties"
	"config-sync/pkg/zlog"
	"os"
	"os/signal"
)

func main() {
	// init logger
	zlog.NewZapLogger()
	defer zlog.Sync()

	// start config-sync
	zlog.Logger.Info("config-sync start")

	properties.InitProperties()
	nacos_config.InitNacosConfig()

	wait()
}

func wait() {
	c := make(chan os.Signal)
	signal.Notify(c)
	<-c
	zlog.Logger.Info("config-sync stop")
}
