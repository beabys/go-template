package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapLogger(out []string, errorOut []string, level zapcore.Level) (*ZapLogger, error) {
	if len(out) == 0 {
		out = []string{"stdout"}
	}
	if len(errorOut) == 0 {
		errorOut = []string{"stderr"}
	}

	logConfig := zap.Config{
		OutputPaths:      out,
		ErrorOutputPaths: errorOut,
		Level:            zap.NewAtomicLevelAt(level),
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
		return nil, fmt.Errorf("error setting up zap logger - %w", err)
	}

	return &ZapLogger{log: l}, nil
}

func (l *ZapLogger) GetLogger() any {
	return l.log
}

func (l *ZapLogger) Debug(s string, lf ...LogField) {
	f := logFieldsToZapFields(lf)
	l.log.Debug(s, f...)
}

func (l *ZapLogger) Info(s string, lf ...LogField) {
	f := logFieldsToZapFields(lf)
	l.log.Info(s, f...)
}

func (l *ZapLogger) Warn(s string, lf ...LogField) {
	f := logFieldsToZapFields(lf)
	l.log.Warn(s, f...)
}

func (l *ZapLogger) Error(s string, err error, lf ...LogField) {
	t := logFieldsToZapFields(lf)
	t = append(t, zap.Error(err))
	l.log.Error(s, t...)
}

func (l *ZapLogger) Fatal(s string, lf ...LogField) {
	f := logFieldsToZapFields(lf)
	// Fatal will call os.Exit(1) after logging the message
	// so we don't need to return an error here.
	l.log.Fatal(s, f...)
}

func (l *ZapLogger) With(lf ...LogField) *zap.Logger {
	f := logFieldsToZapFields(lf)
	if len(f) == 0 {
		return l.log
	}
	return l.log.With(f...)
}

func logFieldsToZapFields(lfs []LogField) []zap.Field {
	fields := make([]zap.Field, len(lfs))
	for i, lf := range lfs {
		switch value := lf.Value.(type) {
		case int:
			fields[i] = zap.Int(lf.Key, value)
		case string:
			fields[i] = zap.String(lf.Key, value)
		case bool:
			fields[i] = zap.Bool(lf.Key, value)
		case float64:
			fields[i] = zap.Float64(lf.Key, value)
		case []byte:
			fields[i] = zap.ByteString(lf.Key, value)
		case zapcore.ObjectMarshaler:
			fields[i] = zap.Object(lf.Key, value)
		case error:
			fields[i] = zap.Error(value)
		case time.Duration:
			fields[i] = zap.Duration(lf.Key, value)
		case time.Time:
			fields[i] = zap.Time(lf.Key, value)
		default:
			// For any other type, we use zap.Any which will handle it appropriately.
			fields[i] = zap.Any(lf.Key, value)
		}
	}
	return fields
}
