package zlog

import (
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger
var baseLogger *zap.Logger

func init() {
	baseLogger, _ = zap.NewDevelopment()
	Logger = baseLogger.Sugar()
}

func Sync() {
	baseLogger.Sync()
}
