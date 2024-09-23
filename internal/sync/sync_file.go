package sync

import (
	"config-sync/pkg/utils/file_util"
	"config-sync/pkg/utils/os_util"
	"config-sync/pkg/zlog"
)

// CheckFileChangedAndExecuteCommand 检查文件是否有变化，并执行命令
func CheckFileChangedAndExecuteCommand(filePath string, content string, command string) error {
	// 是否需要执行命令
	needExecuteCommand := false
	// 判断文件是否存在,不存在则创建文件
	exists, err := file_util.FileExists(filePath)
	if err != nil {
		return err
	}
	if !exists {
		zlog.Logger.Infof("file [%s] not exists, create new file and write content", filePath)
		// 如果文件不存在，则创建文件并写入内容
		err = file_util.WriteToFile(filePath, content)
		if err != nil {
			zlog.Logger.Errorf("write file error: %s", err.Error())
			return err
		}
		// 文件不存在，需要执行命令行
		needExecuteCommand = true
	} else {
		// 已存在文件，判断文件内容是否有变化
		fileContent, err := file_util.ReadFile(filePath)
		if err != nil {
			zlog.Logger.Errorf("read file error: %s", err.Error())
			return err
		}
		if string(fileContent) == content {
			// 文件内容没有变化，直接退出，不执行命令
			zlog.Logger.Infof("[%s] file content not changed, skip execute command ", filePath)
			return nil
		} else {

			//将新文件内容写入原文件
			file_util.WriteToFile(filePath, content)
			zlog.Logger.Debugf("[%s] write content  \n %s ", filePath, content)
			// 文件内容有变化，需要执行命令行
			needExecuteCommand = true
		}
	}
	// 如果需要执行命令行，则执行命令
	if needExecuteCommand {
		//执行命令
		if command != "" {
			executeCommand, err := os_util.ExecuteCommand(command)
			if err != nil {
				return err
			}
			zlog.Logger.Info("execute command result: ", executeCommand)
		} else {
			zlog.Logger.Info("command is empty, skip execute command ")
		}
		return nil
	}
	return nil

}
