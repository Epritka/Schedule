package logger

import (
	"engine/internal/core/interfaces"

	"go.uber.org/zap"
)

type logger struct {
	lw *zap.SugaredLogger
}

func New(lw *zap.SugaredLogger) interfaces.Logger {
	logger := &logger{
		lw: lw,
	}
	return logger
}

func (logger *logger) Error(msg string, args ...interface{}) {
	logger.lw.Errorw(msg, args...)
}

func (logger *logger) Fatal(msg string, args ...interface{}) {
	logger.lw.Fatalw(msg, args...)
}

func (logger *logger) Info(msg string, args ...interface{}) {
	logger.lw.Infow(msg, args...)
}

func (logger *logger) Warn(msg string, args ...interface{}) {
	logger.lw.Warnw(msg, args...)
}

func (logger *logger) Debug(msg string, args ...interface{}) {
	logger.lw.Debugw(msg, args...)
}
