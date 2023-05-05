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

func (l *ZeroLogger) Debug(msg string, fields map[string]interface{}) {
	l.logger.Debug().Fields(fields).Msg(msg)
}

func (l *ZeroLogger) Info(msg string, fields map[string]interface{}) {
	l.logger.Info().Fields(fields).Msg(msg)
}

func (l *ZeroLogger) Warn(msg string, fields map[string]interface{}) {
	l.logger.Warn().Fields(fields).Msg(msg)
}
func (l *ZeroLogger) Error(msg string, fields map[string]interface{}) {
	l.logger.Error().Fields(fields).Msg(msg)
}
func (l *ZeroLogger) Fatal(msg string, fields map[string]interface{}) {
	l.logger.Fatal().Fields(fields).Msg(msg)
}
