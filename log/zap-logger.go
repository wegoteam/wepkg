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

func (l *ZapLogger) Debug(field ...interface{}) {
	l.logger.Debug("", zap.Any("args", field))
}

func (l *ZapLogger) Info(field ...interface{}) {
	l.logger.Info("", zap.Any("args", field))
}

func (l *ZapLogger) Warn(field ...interface{}) {
	l.logger.Warn("", zap.Any("args", field))
}
func (l *ZapLogger) Error(field ...interface{}) {
	l.logger.Error("", zap.Any("args", field))
}
func (l *ZapLogger) Fatal(field ...interface{}) {
	l.logger.Fatal("", zap.Any("args", field))
}
