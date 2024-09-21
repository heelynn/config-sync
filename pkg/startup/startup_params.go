package startup

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// file / console

var RootConfig string
var RootConfigPath string

func init() {
	initRootConfigPath()
	flag.StringVar(&RootConfig, "config", filepath.Join(RootConfigPath, string(filepath.Separator), "application.yaml"), "root application.yaml config file path")
	flag.StringVar(&RootConfigPath, "config-path", RootConfigPath, "root config directory")
	flag.Parse()
	fmt.Println("configPath:", RootConfig)
	fmt.Println("RootConfigPath:", RootConfigPath)
}

// initRootConfigPath init root config path (../conf)
func initRootConfigPath() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	supperDir := filepath.Dir(wd)
	RootConfigPath = filepath.Join(supperDir, string(filepath.Separator), "conf")

}
