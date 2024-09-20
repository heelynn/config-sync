package zlog

import (
	"go.uber.org/zap/zapcore"
	"time"
)

const (
	logTmFmtWithMS = "2006-01-02 15:04:05.000"
)

func getEncoder() zapcore.Encoder {
	// 自定义时间输出格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + t.Format(logTmFmtWithMS) + "]")
	}
	//// 自定义日志级别显示
	//customLevelEncoder := func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	//	enc.AppendString("[" + level.CapitalString() + "]")
	//}

	//// 自定义文件：行号输出项
	//customCallerEncoder := func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	//	enc.AppendString("[" + caller.TrimmedPath() + "]")
	//}
	// 自定义文件：行号输出项
	customCallerEncoder := func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + caller.TrimmedPath() + "]")
	}

	encoderConf := zapcore.EncoderConfig{
		CallerKey:      "caller_line", // 打印文件名和行数
		LevelKey:       "level_name",
		MessageKey:     "msg",
		TimeKey:        "ts",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     customTimeEncoder,                // 自定义时间格式
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 小写编码器
		EncodeCaller:   customCallerEncoder,              // 全路径编码器
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	return zapcore.NewConsoleEncoder(encoderConf)
}
