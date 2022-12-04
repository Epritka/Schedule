package main

import (
	"engine/infrastructure/config"
	"engine/internal/core/usecase/schedule"
	"engine/internal/core/usecase/user"
	"engine/internal/delivery/http"
	"engine/internal/mock/repository"
	"log"
)

func main() {
	config, err := config.New("./config.yaml")

	if err != nil {
		log.Fatal("Error initialize config:", err)
	}

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
