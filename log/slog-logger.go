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

func (l *SlogLogger) Debug(msg string, fields map[string]interface{}) {
	slog.Debug()
	l.logger.WithFields(fields).Debug(msg)
}

func (l *SlogLogger) Info(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Info(msg)
}

func (l *SlogLogger) Warn(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Warn(msg)
}
func (l *SlogLogger) Error(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Error(msg)
}
func (l *SlogLogger) Fatal(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Fatal(msg)
}
