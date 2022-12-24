package main

import (
	"engine/infrastructure/config"
	"engine/infrastructure/pg"
	"engine/infrastructure/zap"
	"engine/internal/adapter/logger"
	"engine/internal/core/usecase/day"
	"engine/internal/core/usecase/group"
	"engine/internal/core/usecase/schedule"
	"engine/internal/core/usecase/student"

	"engine/internal/adapter/repository"
	"engine/internal/delivery/http"
)

func main() {
	config := config.New(".")
	database := pg.New(config.DBUrl, config.IsDebug)
	zap := zap.New(config.IsDebug)

	logger := logger.New(zap)

	repositoryManager := repository.NewRepositoryManager(database, nil)

	dayUseCase := day.New(
		repositoryManager,
		logger,
	)

	scheduleUseCase := schedule.New(
		repositoryManager,
		logger,
	)

	groupUseCase := group.New(
		repositoryManager,
		logger,
	)

	studentUseCase := student.New(
		repositoryManager,
		logger,
	)

	httpServer := http.New(
		config.IsDebug,
		config.Port,
		dayUseCase,
		scheduleUseCase,
		studentUseCase,
		groupUseCase,
	)

	httpServer.Run()
}
