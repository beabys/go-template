package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	GetLogger() any
	Debug(string, ...zap.Field)
	Info(string, ...zap.Field)
	Warn(string, ...zap.Field)
	Error(string, error, ...zap.Field)
	Fatal(string, ...zap.Field)
}

type DefaultLogger struct {
	log *zap.Logger
}

type DefaultLoggerConfig struct {
	Out   []string
	Error []string
	Level zapcore.Level
}
