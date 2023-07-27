package log

import (
	"context"
	"github.com/gookit/slog"
)

type SlogLogger struct {
	logger *slog.SugaredLogger
	ctx    context.Context
}

func NewSlogLogger(ctx context.Context) *SlogLogger {
	slog.Configure(func(logger *slog.SugaredLogger) {
		f := logger.Formatter.(*slog.TextFormatter)
		f.EnableColor = true
	})
	return &SlogLogger{logger: slog.Std(), ctx: ctx}
}

func (l *SlogLogger) Debugf(msg string, field ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *SlogLogger) Infof(msg string, field ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *SlogLogger) Warnf(msg string, field ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *SlogLogger) Errorf(msg string, field ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *SlogLogger) Fatalf(msg string, field ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *SlogLogger) Debug(field ...interface{}) {
	slog.Debug(field)
}

func (l *SlogLogger) Info(field ...interface{}) {
	slog.Info(field)
}

func (l *SlogLogger) Warn(field ...interface{}) {
	slog.Warn(field)
}
func (l *SlogLogger) Error(field ...interface{}) {
	slog.Error(field)
}
func (l *SlogLogger) Fatal(field ...interface{}) {
	slog.Fatal(field)
}
