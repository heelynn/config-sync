package config

import (
	"config-sync/pkg/utils/file_util"
	"config-sync/pkg/utils/os_util"
	"config-sync/pkg/zlog"
)

type ConfigExecutor interface {
	RegisterChangedListener() error
}

// SyncConfigToFile writes the nacos config to file
func SyncConfigToFile(filepath, filename string, content string) {
	// Get nacos properties
	// Write to file
	fileName := file_util.GetFileName(filepath, filename)
	err := file_util.WriteToFile(fileName, content)
	if err != nil {
		zlog.Logger.Error(err)
	}
	zlog.Logger.Infof("nacos config generated config written to file,fileName: %s", fileName)
	zlog.Logger.Debugf("[%s] write content  \n %s ", fileName, content)

}

// ExecuteCommand executes the command
func ExecuteCommand(command string) string {
	// Execute command
	if command != "" {
		result, err := os_util.ExecuteCommand(command)
		if err != nil {
			zlog.Logger.Error(err)
		}
		zlog.Logger.Debugf("\n -- Command--\n %s", command)
		zlog.Logger.Debugf("\n -- Command Result--\n %s", result)
		return result
	}
	return ""

}
