package log

import (
	"context"
	"github.com/rs/zerolog"
	"os"
)

type ZeroLogger struct {
	logger *zerolog.Logger
	ctx    context.Context
}

func NewZeroLogger(ctx context.Context) *ZeroLogger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zlogger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	return &ZeroLogger{logger: &zlogger, ctx: ctx}
}

func (l *ZeroLogger) Debugf(msg string, field ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *ZeroLogger) Infof(msg string, field ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *ZeroLogger) Warnf(msg string, field ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *ZeroLogger) Errorf(msg string, field ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *ZeroLogger) Fatalf(msg string, field ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *ZeroLogger) Debug(field ...interface{}) {
	l.logger.Debug().Msg("")
}

func (l *ZeroLogger) Info(field ...interface{}) {
}

func (l *ZeroLogger) Warn(field ...interface{}) {
}
func (l *ZeroLogger) Error(field ...interface{}) {
}
func (l *ZeroLogger) Fatal(field ...interface{}) {
}
