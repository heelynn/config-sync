package file_util

import (
	"fmt"
	"os"
	"path/filepath"
)

func WriteToFile(fileName string, content string) error {

	// 使用os.Create创建或打开文件
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating file: %s\n", err)
		return err
	}
	defer file.Close() // 确保文件在函数结束时关闭

	// 写入内容到文件
	contents := []byte(content)
	_, err = file.Write(contents)
	if err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
		return err
	}

	fmt.Printf("File '%s' has been created and written to.\n", fileName)
	return nil
}

func GetFileName(path string, fileName string) string {
	if IsLastCharPathSeparator(path) {
		return filepath.Join(path, fileName)
	} else {
		return filepath.Join(path, string(filepath.Separator), fileName)
	}
}

// IsLastCharPathSeparator 检查给定的路径字符串的最后一个字符是否为文件路径分隔符
func IsLastCharPathSeparator(path string) bool {
	// 获取当前操作系统的路径分隔符
	separator := filepath.Separator

	// 检查字符串长度是否大于0，并且最后一个字符是否为分隔符
	if len(path) > 0 && (int32)(path[len(path)-1]) == separator {
		return true
	}

	return false
}
