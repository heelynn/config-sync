package main

import (
	"config-sync/internal/backends/nacos"
	"config-sync/internal/properties"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	properties.InitProperties()
	nacos.InitNacosConfig()
	wait()
}

func wait() {
	c := make(chan os.Signal)
	signal.Notify(c)
	go func() {
		fmt.Println("Go routine running")
		time.Sleep(3 * time.Second)
		fmt.Println("Go routine done")
	}()
	<-c
	fmt.Println("bye")
}
