package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var globalLog *zap.Logger

func NewLog(filePath string) {
	// lumberjack 配置
	rotateLogs := &lumberjack.Logger{
		Filename:   filePath, // 日志文件的位置
		MaxSize:    15,       // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 1024,     // 保留旧文件的最大个数
		MaxAge:     365 * 3,  // 保留旧文件的最大天数
		Compress:   true,     // 是否压缩/归档旧文件
	}

	// zap 配置
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(rotateLogs),
		atomicLevel,
	)

	globalLog = zap.New(core)
}
