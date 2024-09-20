package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.SugaredLogger
var baseLogger *zap.Logger

// var baseLogger *zap.Logger
var logConfig *LogConfig

func initLogConfig(Output string, Level string, Path string, MaxSize int, MaxAge int, MaxBackups int) {
	logConfig = &LogConfig{
		Output:     Output,
		Level:      Level,
		Path:       Path,
		MaxSize:    MaxSize,
		MaxAge:     MaxAge,
		MaxBackups: MaxBackups,
	}
}

func InitLog(Output string, Level string, Path string, MaxSize int, MaxAge int, MaxBackups int) {
	initLogConfig(Output, Level, Path, MaxSize, MaxAge, MaxBackups)
	var writeSyncer zapcore.WriteSyncer
	if logConfig.Output == "console" {
		// 打印到控制台
		writeSyncer = zapcore.AddSync(os.Stdout)
	} else if logConfig.Output == "file" {
		// 打印到文件
		writeSyncer = getLogWriter()
	}
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, getLogLevel())
	baseLogger = zap.New(core)

	Logger = baseLogger.Sugar()
}

func getLogLevel() zapcore.Level {
	switch logConfig.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func Sync() {
	baseLogger.Sync()
}

// logConfig 定义了日志配置的结构体
type LogConfig struct {
	Output     string // 输出方式，例如 "console" 或 "file"
	Level      string // 日志级别，例如 "info"
	Path       string // 日志文件路径
	MaxSize    int    // 日志文件最大大小（MB）
	MaxAge     int    // 日志文件保留的最大天数
	MaxBackups int    // 日志文件的最大备份数量
}
