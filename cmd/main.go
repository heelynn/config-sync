package main

import (
	_ "config-sync/internal/backends/config/nacos"
	_ "config-sync/internal/backends/discovery/nacos"
	_ "config-sync/internal/properties"
	_ "config-sync/pkg/startup"
	"config-sync/pkg/zlog"
	"os"
	"os/signal"
)

func main() {

	zlog.Logger.Info("--------- start ---------")
	defer zlog.Sync()

	wait()
}

func wait() {
	c := make(chan os.Signal)
	signal.Notify(c)
	<-c
	zlog.Logger.Info(("--------- stoped ---------"))
}
