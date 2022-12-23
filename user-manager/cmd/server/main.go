package main

import (
	"user-manager/infrastructure/config"
	"user-manager/infrastructure/pg"
	"user-manager/infrastructure/zap"
	"user-manager/internal/adapter/logger"
	"user-manager/internal/adapter/repository"
	"user-manager/internal/core/usecase/user"

	"user-manager/internal/delivery/http"
)

func main() {
	config := config.New(".")
	database := pg.New(config.DBUrl, config.IsDebug)
	zap := zap.New(config.IsDebug)

	logger := logger.New(zap)

	repositoryManager := repository.NewRepositoryManager(database, nil)

	userUseCase := user.New(
		repositoryManager,
		logger,
	)

	httpServer := http.New(
		!config.IsDebug,
		config.Port,
		userUseCase,
	)

	httpServer.Run()
}
