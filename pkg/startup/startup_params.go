package startup

import (
	"flag"
	"fmt"
	"path/filepath"
)

// file / console

var RootConfigPath string

func init() {
	//  default value is ../conf/application.yaml
	flag.StringVar(&RootConfigPath, "config",
		filepath.Join("..", string(filepath.Separator), "conf", string(filepath.Separator), "application.yaml"),
		"root config file path")
	flag.Parse()
	fmt.Println("configPath:", RootConfigPath)
}
