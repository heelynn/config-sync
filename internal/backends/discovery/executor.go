package discovery

import (
	"bytes"
	"config-sync/pkg/utils/file_util"
	"config-sync/pkg/utils/os_util"
	"config-sync/pkg/zlog"
	"text/template"
)

// Executor is the interface
type Executor interface {
	Execute() error
	TickerExecute() error
}

//func WriteTemplate(templateFile string, filePath string, instances DiscoveryResult) error {
//	if templateFile != "" {
//		file, err := file_util.ReadFile(templateFile)
//		if err != nil {
//			return err
//		}
//		// 创建模板
//		tmpl, err := template.New("instances").Parse(string(file))
//		if err != nil {
//			return err
//		}
//		// 填充模板
//		// 执行模板，传入ages map
//		// 打开文件准备写入
//		zlog.Logger.Info("write file to ", filePath)
//		writeFile, err := os.Create(filePath + new_FILE_SUFFIX)
//		if err != nil {
//			panic(err)
//		}
//		defer writeFile.Close() // 确保文件在函数结束时关闭
//		err = tmpl.Execute(writeFile, instances)
//		if err != nil {
//			panic(err)
//		}
//	}
//	return nil
//
//}

func GenerateTemplate(templateFile string, instances DiscoveryResult) (string, error) {
	if templateFile != "" {
		file, err := file_util.ReadFile(templateFile)
		if err != nil {
			return "", err
		}
		// 创建模板
		tmpl, err := template.New("instances").Parse(string(file))
		if err != nil {
			return "", err
		}
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, instances)
		if err != nil {
			panic(err)
		}
		return buf.String(), nil
	}
	return "", nil

}

// CheckFileChangedAndExecuteCommand 检查文件是否有变化，并执行命令
func CheckFileChangedAndExecuteCommand(filePath string, templateContent string, command string) (bool, error) {
	fileContent, err := file_util.ReadFile(filePath)
	if err != nil {
		return false, err
	}
	// 如果文件内容没有变化，则不执行命令
	if string(fileContent) == string(templateContent) {
		zlog.Logger.Infof("[%s] file content not changed, skip command ", filePath)
		if err != nil {
			return false, err
		}
		return false, nil
	} else {
		//将新文件内容写入原文件
		file_util.WriteToFile(filePath, templateContent)
		zlog.Logger.Infof("nacos discovery generated config written to file,fileName: %s", filePath)
		zlog.Logger.Debugf("[%s] write content  \n %s ", filePath, templateContent)
		//执行命令
		if command != "" {
			executeCommand, err := os_util.ExecuteCommand(command)
			if err != nil {
				return false, err
			}
			zlog.Logger.Info("execute command result: ", executeCommand)
		}
		return true, nil
	}

}
