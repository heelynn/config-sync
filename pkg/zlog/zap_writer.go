package zlog

import (
	"config-sync/pkg/startup"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
)

func getLogWriter() zapcore.WriteSyncer {
	/*
		Filename: 日志文件的位置
		MaxSize：在进行切割之前，日志文件的最大大小（以 MB 为单位）
		MaxBackups：保留旧文件的最大个数
		MaxAges：保留旧文件的最大天数
		Compress：是否压缩 / 归档旧文件
	*/
	logPath := startup.RootLogPath

	_, err := os.Stat(logPath)
	if os.IsNotExist(err) {
		panic("log file path [" + logPath + "] not exists")
	}
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath.Join(logPath, string(filepath.Separator), "info.log"),
		MaxSize:    logConfig.MaxSize,
		MaxBackups: logConfig.MaxBackups,
		MaxAge:     logConfig.MaxAge,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
