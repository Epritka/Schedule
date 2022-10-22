package log

import (
	"engine/infrastructure/config"

	"go.uber.org/zap"
)

type Logger struct {
	Logger *zap.Logger
}

func NewLogger(config config.Config) (*Logger, error) {
	logger, err := zap.NewProduction()

	if err != nil {
		return nil, err
	}

	return &Logger{
		logger,
	}, nil
}
