package logger

import (
	"fmt"

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

func NewDefaultLogger(config *DefaultLoggerConfig) (*DefaultLogger, error) {
	if config.Out == nil {
		config.Out = []string{"stdout"}
	}
	if config.Error == nil {
		config.Error = []string{"stderr"}
	}

	logConfig := zap.Config{
		OutputPaths:      config.Out,
		ErrorOutputPaths: config.Error,
		Level:            zap.NewAtomicLevelAt(config.Level),
		Encoding:         "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	l, err := logConfig.Build()
	if err != nil {
		return nil, fmt.Errorf("error setting up default logger - %w", err)
	}

	return &DefaultLogger{l}, nil
}

func (l *DefaultLogger) GetLogger() any {
	return l.log
}

func (l *DefaultLogger) Debug(s string, f ...zap.Field) {
	l.log.Debug(s, f...)
}

func (l *DefaultLogger) Info(s string, f ...zap.Field) {
	l.log.Info(s, f...)
}

func (l *DefaultLogger) Warn(s string, f ...zap.Field) {
	l.log.Warn(s, f...)
}

func (l *DefaultLogger) Error(s string, err error, t ...zap.Field) {
	t = append(t, zap.Error(err))
	l.log.Error(s, t...)
}

func (l *DefaultLogger) Fatal(s string, f ...zap.Field) {
	l.log.Fatal(s, f...)
}

func (l *DefaultLogger) With(f ...zap.Field) *zap.Logger {
	return l.log.With(f...)
}
