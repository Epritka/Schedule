package main

import (
	"engine/infrastructure/config"
	"engine/infrastructure/pg"
	"engine/infrastructure/zap"
	"engine/internal/adapter/logger"
	"engine/internal/core/usecase/day"

	"engine/internal/adapter/repository"
	"engine/internal/delivery/http"
)

func main() {
	// infrastructure
	config := config.New(".")
	database := pg.New(config.DBUrl, config.IsDebug)
	zap := zap.New(config.IsDebug)

	// adapter
	logger := logger.New(zap)

	// repository
	repositoryManager := repository.NewRepositoryManager(database, nil)

	// core
	dayUseCase := day.New(
		repositoryManager,
		logger,
	)

	// delivery
	httpServer := http.New(
		config.IsDebug,
		config.Port,
		dayUseCase,
	)

	httpServer.Run()
}
