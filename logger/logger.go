package logger

import (
	"github.com/dubbogo/getty"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log       getty.Logger
	zapLogger *zap.Logger

	zapLoggerConfig        = zap.NewDevelopmentConfig()
	zapLoggerEncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
)

func init() {
	zapLoggerConfig.EncoderConfig = zapLoggerEncoderConfig
	zapLogger, _ = zapLoggerConfig.Build()
	log = zapLogger.Sugar()
}

func ZapLogger() getty.Logger {
	return log
}
