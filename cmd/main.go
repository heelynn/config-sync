package main

import (
	_ "config-sync/internal/backends/config/nacos"
	_ "config-sync/internal/backends/discovery/nacos"
	_ "config-sync/internal/properties"
	_ "config-sync/pkg/startup"
	"config-sync/pkg/zlog"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	zlog.Logger.Info("--------- start ---------")
	defer zlog.Sync()

	wait()
}

func wait() {
	var sig os.Signal
	for {
		c := make(chan os.Signal)
		signal.Notify(c)

		sig = <-c
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			zlog.Logger.Infof("--------- stop signal %v ---------", sig)
			return
		default:
			//zlog.Logger.Debugf("--------- receive signal %v ---------", sig)
		}

	}
}
