package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	GetLogger() any
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Fatal(...interface{})
}

type DefaultLogger struct {
	log *logrus.Logger
}
type DefaultLoggerConfig struct {
	Formater logrus.Formatter
	Out      io.Writer
	Level    logrus.Level
}

func NewDefaultLogger(config *DefaultLoggerConfig) *DefaultLogger {
	if config.Formater == nil {
		config.Formater = &logrus.JSONFormatter{}
	}
	if config.Formater == nil {
		config.Out = os.Stderr
	}
	if config.Formater == nil {
		config.Level = logrus.DebugLevel
	}
	logger := &logrus.Logger{
		Formatter: config.Formater,
		Out:       config.Out,
		Level:     config.Level,
	}
	return &DefaultLogger{logger}
}

func (l *DefaultLogger) GetLogger() any {
	return l.log
}

func (l *DefaultLogger) Debug(v ...interface{}) {
	l.log.Debug(v...)
}

func (l *DefaultLogger) Info(v ...interface{}) {
	l.log.Info(v...)
}

func (l *DefaultLogger) Warn(v ...interface{}) {
	l.log.Warn(v...)
}

func (l *DefaultLogger) Error(v ...interface{}) {
	l.log.Error(v...)
}

func (l *DefaultLogger) Fatal(v ...interface{}) {
	l.log.Fatal(v...)
}

func (l *DefaultLogger) With(key string, value any) *logrus.Entry {
	return l.log.WithField(key, value)
}
