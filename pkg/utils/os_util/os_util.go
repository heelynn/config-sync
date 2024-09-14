package os_util

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

func ExecuteCommand(command string) (string, error) {
	// 创建一个命令对象，这里以"ls"为例
	cmd := exec.Command(command)
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd.exe", "/c", command)
	case "darwin":
		cmd = exec.Command("/bin/sh", "-c", command)
	case "linux":
		cmd = exec.Command("/bin/sh", "-c", command)
	default:
		return "", fmt.Errorf("不支持的系统类型: %s", runtime.GOOS)
	}
	// 创建一个缓冲区来保存命令的输出
	var out bytes.Buffer
	cmd.Stdout = &out

	// 执行命令
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("命令执行失败: %s\n", err)
	}

	// 输出命令结果
	return out.String(), nil
}
