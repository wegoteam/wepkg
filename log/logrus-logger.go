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

func (l *LogrusLogger) Debug(msg string, fields map[string]interface{}) {

	l.logger.WithFields(fields).Debug(msg)
}

func (l *LogrusLogger) Info(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Info(msg)
}

func (l *LogrusLogger) Warn(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Warn(msg)
}
func (l *LogrusLogger) Error(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Error(msg)
}
func (l *LogrusLogger) Fatal(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Fatal(msg)
}
