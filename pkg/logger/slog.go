package logger

import (
	"context"
	"log/slog"
	"os"
	"strings"
)

func NewSlogLogger(level string) (*SlogLogger, error) {
	var l slog.Level
	switch strings.ToLower(level) {
	case "info":
		l = slog.LevelInfo
	case "warn":
		l = slog.LevelWarn
	case "error":
		l = slog.LevelError
	default:
		l = slog.LevelDebug
	}

	opts := &slog.HandlerOptions{
		Level: l,
		// added to ensure that the log level is always in lowercase
		// this is useful for consistency, especially when using JSON format.
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				return slog.Attr{Key: a.Key, Value: slog.StringValue(strings.ToLower(a.Value.String()))}
			}
			return a
		},
	}
	handler := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	return &SlogLogger{log: handler}, nil
}

func (l *SlogLogger) GetLogger() any {
	return l.log
}

func (l *SlogLogger) Debug(s string, lf ...LogField) {
	f := logFieldsToSlogFields(lf)
	l.log.Debug(s, f...)
}

func (l *SlogLogger) Info(s string, lf ...LogField) {
	f := logFieldsToSlogFields(lf)
	l.log.Info(s, f...)
}

func (l *SlogLogger) Warn(s string, lf ...LogField) {
	f := logFieldsToSlogFields(lf)
	l.log.Warn(s, f...)
}

func (l *SlogLogger) Error(s string, err error, lf ...LogField) {
	errField := LogField{"error", err.Error()}
	lf = append(lf, errField)
	t := logFieldsToSlogFields(lf)
	l.log.Error(s, t...)
}

func (l *SlogLogger) Fatal(s string, lf ...LogField) {
	f := logFieldsToSlogFields(lf)
	// Fatal will call os.Exit(1) after logging the message
	// so we don't need to return an error here.
	l.log.Log(context.Background(), slog.Level(12), s, f...)
	os.Exit(1)
}

func (l *SlogLogger) With(lf ...LogField) *slog.Logger {
	if len(lf) == 0 {
		return l.log
	}
	f := logFieldsToSlogFields(lf)
	return l.log.With(f...)
}

func logFieldsToSlogFields(lfs []LogField) []any {
	var f []any
	for _, v := range lfs {
		f = append(f, v.Key, v.Value)
	}
	return f
}
