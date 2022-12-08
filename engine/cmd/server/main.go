package main

import (
	"engine/infrastructure/config"
	"engine/infrastructure/pg"
	"engine/infrastructure/zap"
	"engine/internal/adapter/cryptographer"
	"engine/internal/adapter/logger"
	"engine/internal/core/usecase/user"

	"engine/internal/delivery/http"
	"engine/internal/repository"
)

func main() {
	// infrastructure
	config := config.New(".")
	database := pg.New(config.DBUrl, config.IsDebug)
	zap := zap.New(config.IsDebug)

	// adapter
	logger := logger.New(zap)
	cryptographer := cryptographer.New()

	// repository
	repositoryManager := repository.NewRepositoryManager(database, nil)

	// core
	userUseCase := user.New(
		repositoryManager,
		cryptographer,
		logger,
	)

	// delivery
	httpServer := http.New(
		!config.IsDebug,
		config.Port,
		userUseCase,
	)
	httpServer.Run()
}
