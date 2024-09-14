package properties

import "config-sync/pkg/zlog"

var Prop *Properties

func InitProperties() {
	zlog.Logger.Info("Initializing properties")
	parser := newYamlParser()
	Prop = parser.Parse(parser.GetFilePath())

}
