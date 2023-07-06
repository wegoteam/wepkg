package log

import (
	"context"
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.Logger
	ctx    context.Context
}

func NewZapLog(ctx context.Context) *ZapLogger {
	logger, _ := zap.NewProduction()
	return &ZapLogger{logger: logger, ctx: ctx}
}

func (l *ZapLogger) Debugf(msg string, field ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *ZapLogger) Infof(msg string, field ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *ZapLogger) Warnf(msg string, field ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *ZapLogger) Errorf(msg string, field ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *ZapLogger) Fatalf(msg string, field ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *ZapLogger) Debug(msg string, fields map[string]interface{}) {
	l.logger.Debug(msg, zap.Any("args", fields))
}

func (l *ZapLogger) Info(msg string, fields map[string]interface{}) {
	l.logger.Info(msg, zap.Any("args", fields))
}

func (l *ZapLogger) Warn(msg string, fields map[string]interface{}) {
	l.logger.Warn(msg, zap.Any("args", fields))
}
func (l *ZapLogger) Error(msg string, fields map[string]interface{}) {
	l.logger.Error(msg, zap.Any("args", fields))
}
func (l *ZapLogger) Fatal(msg string, fields map[string]interface{}) {
	l.logger.Fatal(msg, zap.Any("args", fields))
}
