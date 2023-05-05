package log

import (
	"context"
)

//https://github.com/uber-go/zap
//https://github.com/Sirupsen/logrus
//https://github.com/gookit/slog
//https://github.com/rs/zerolog

type LoggerWrapper struct {
	logger Logger
}

type Logger interface {
	Debug(msg string, fields map[string]interface{})
	Info(msg string, fields map[string]interface{})
	Warn(msg string, fields map[string]interface{})
	Error(msg string, fields map[string]interface{})
	Fatal(msg string, fields map[string]interface{})

	Debugf(msg string, field ...interface{})
	Infof(msg string, field ...interface{})
	Warnf(msg string, field ...interface{})
	Errorf(msg string, field ...interface{})
	Fatalf(msg string, field ...interface{})
}

func NewLoggerWrapper(loggerType string, ctx context.Context) *LoggerWrapper {
	var logger Logger

	switch loggerType {
	case "logrus":
		logger = NewLogrusLogger(ctx)
	case "zap":
		logger = NewZapLog(ctx)
	case "zerolog":
		logger = NewZeroLogger(ctx)
	case "slog":
		logger = NewSlogLogger(ctx)
	default:
		logger = NewLogrusLogger(ctx)
	}
	return &LoggerWrapper{logger: logger}
}
