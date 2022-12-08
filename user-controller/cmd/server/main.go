package main

import (
	"user-controller/infrastructure/config"
	"user-controller/infrastructure/pg"
	"user-controller/infrastructure/zap"
	"user-controller/internal/adapter/cryptographer"
	"user-controller/internal/adapter/logger"
	"user-controller/internal/adapter/repository"
	"user-controller/internal/core/usecase/user"

	"user-controller/internal/delivery/http"
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
