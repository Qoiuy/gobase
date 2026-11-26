package zlog

import "go.uber.org/zap"

func init() {
	NewLog("./logs/go-base.log")
}

func Infof(format string, args ...zap.Field) {
	globalLog.Log(zap.InfoLevel, format, args...)
}
