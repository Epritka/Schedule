package main

import (
	"engine/infrastructure/config"
	"engine/internal/adapter/repository"
	"engine/internal/core/usecase/schedule"
	"engine/internal/core/usecase/user"
	"engine/internal/delivery/http"
	"log"
)

func main() {
	config := config.Config{Port: 2513, IsDebug: true}

	repository := repository.NewRepositoryManager()

	scheduleUseCase := schedule.NewScheduleUseCase(repository)
	userUseCase := user.NewUserUseCase(repository)

	server, err := http.NewHttpServer(
		config,
		scheduleUseCase,
		userUseCase,
	)

	if err != nil {
		log.Fatal("Error initialize server:", err)
	}

	if err := server.Start(); err != nil {
		log.Fatal("Server error:", err)
	}
}
