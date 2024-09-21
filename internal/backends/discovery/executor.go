package discovery

import (
	"bytes"
	"config-sync/pkg/startup"
	"config-sync/pkg/utils/file_util"
	"config-sync/pkg/zlog"
	"path/filepath"
	"strings"
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
		if !strings.HasPrefix(templateFile, "/") {
			templateFile = startup.RootConfigPath + string(filepath.Separator) + templateFile
			zlog.Logger.Info("--- templateFile ", templateFile)
		}
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
