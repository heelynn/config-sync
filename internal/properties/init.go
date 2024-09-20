package properties

import (
	"config-sync/pkg/startup"
	"config-sync/pkg/zlog"
)

var Prop *Properties

func init() {

	parser := newYamlParser()
	filePath := parser.GetFilePath()
	if startup.RootConfigPath != "" {
		filePath = startup.RootConfigPath
	}
	Prop = parser.Parse(filePath)
	SetLogDefaultValues(Prop)
	log := Prop.Log
	zlog.InitLog(log.Output, log.Level, log.Path, log.MaxSize, log.MaxAge, log.MaxBackups)
	zlog.Logger.Info("Initializing properties")

}
