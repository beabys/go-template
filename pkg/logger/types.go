package logger

import (
	"log/slog"

	"go.uber.org/zap"
)

type Logger interface {
	GetLogger() any
	Debug(string, ...LogField)
	Info(string, ...LogField)
	Warn(string, ...LogField)
	Error(string, error, ...LogField)
	Fatal(string, ...LogField)
}

type ZapLogger struct {
	log *zap.Logger
}

type SlogLogger struct {
	log *slog.Logger
}

type LogField struct {
	Key   string
	Value any
}
