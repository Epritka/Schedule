package main

import (
	"user-manager/infrastructure/config"
	"user-manager/infrastructure/pg"
	"user-manager/infrastructure/zap"
	"user-manager/internal/adapter/cryptographer"
	"user-manager/internal/adapter/logger"
	"user-manager/internal/adapter/repository"
	"user-manager/internal/core/usecase/user"

	"user-manager/internal/delivery/http"
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
