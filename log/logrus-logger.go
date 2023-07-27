package log

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
)

type LogrusLogger struct {
	logger *logrus.Logger
	ctx    context.Context
}

func NewLogrusLogger(ctx context.Context) *LogrusLogger {
	logger := logrus.New()
	logger.Out = os.Stdout
	return &LogrusLogger{logger: logger, ctx: ctx}
}

func (l *LogrusLogger) Debugf(msg string, field ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *LogrusLogger) Infof(msg string, field ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *LogrusLogger) Warnf(msg string, field ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *LogrusLogger) Errorf(msg string, field ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *LogrusLogger) Fatalf(msg string, field ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *LogrusLogger) Debug(field ...interface{}) {

	l.logger.Debug(field)
}

func (l *LogrusLogger) Info(field ...interface{}) {
	l.logger.Info(field)
}

func (l *LogrusLogger) Warn(field ...interface{}) {
	l.logger.Warn(field)
}
func (l *LogrusLogger) Error(field ...interface{}) {
	l.logger.Error(field)
}
func (l *LogrusLogger) Fatal(field ...interface{}) {
	l.logger.Fatal(field)
}
