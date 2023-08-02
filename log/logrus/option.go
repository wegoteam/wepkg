package logrus

import (
	"github.com/sirupsen/logrus"
)

// Option logger options
type Option interface {
	apply(cfg *config)
}

type option func(cfg *config)

func (fn option) apply(cfg *config) {
	fn(cfg)
}

type config struct {
	logger *logrus.Logger
	hooks  []logrus.Hook
}

func defaultConfig() *config {
	// std logger
	stdLogger := logrus.StandardLogger()
	// default json format
	stdLogger.SetFormatter(new(logrus.JSONFormatter))

	return &config{
		logger: logrus.StandardLogger(),
		hooks:  []logrus.Hook{},
	}
}

// WithLogger configures logger
func WithLogger(logger *logrus.Logger) Option {
	return option(func(cfg *config) {
		cfg.logger = logger
	})
}

// WithHook configures logrus hook
func WithHook(hook logrus.Hook) Option {
	return option(func(cfg *config) {
		cfg.hooks = append(cfg.hooks, hook)
	})
}
