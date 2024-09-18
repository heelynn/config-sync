package discovery

import (
	"config-sync/pkg/utils/file_util"
	"config-sync/pkg/utils/os_util"
	"config-sync/pkg/zlog"
	"os"
	"text/template"
)

// Executor is the interface
type Executor interface {
	Execute() error
	TickerExecute() error
}

const new_FILE_SUFFIX = ".new"

func WriteTemplate(templateFile string, filePath string, instances DiscoveryResult) error {
	if templateFile != "" {
		file, err := file_util.ReadFile(templateFile)
		if err != nil {
			return err
		}
		// 创建模板
		tmpl, err := template.New("instances").Parse(string(file))
		if err != nil {
			return err
		}
		// 填充模板
		// 执行模板，传入ages map
		// 打开文件准备写入
		zlog.Logger.Info("write file to ", filePath)
		writeFile, err := os.Create(filePath + new_FILE_SUFFIX)
		if err != nil {
			panic(err)
		}
		defer writeFile.Close() // 确保文件在函数结束时关闭
		err = tmpl.Execute(writeFile, instances)
		if err != nil {
			panic(err)
		}
	}
	return nil

}

// CheckFileChangedAndExecuteCommand 检查文件是否有变化，并执行命令
func CheckFileChangedAndExecuteCommand(filePath string, command string) (bool, error) {
	file, err := file_util.ReadFile(filePath)
	if err != nil {
		return false, err
	}
	newFile, err := file_util.ReadFile(filePath + new_FILE_SUFFIX)
	if err != nil {
		return false, err
	}
	// 如果文件内容没有变化，则不执行命令
	if string(file) == string(newFile) {
		zlog.Logger.Infof("[%s] file content not changed, skip command ", filePath)
		err := file_util.RemoveFile(filePath + new_FILE_SUFFIX)
		if err != nil {
			return false, err
		}
		return false, nil
	} else {
		//将新文件内容写入原文件
		file_util.WriteToFile(filePath, string(newFile))
		//删除临时文件
		file_util.RemoveFile(filePath + new_FILE_SUFFIX)
		executeCommand, err := os_util.ExecuteCommand(command)
		if err != nil {
			return false, err
		}
		zlog.Logger.Info("execute command result: ", executeCommand)
		return true, nil
	}

}
