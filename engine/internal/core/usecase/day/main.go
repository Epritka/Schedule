package day

import "engine/internal/core/interfaces"

type useCase struct {
	repository interfaces.RepositoryManager
}

func NewScheduleUseCase(repository interfaces.RepositoryManager) interfaces.DayUseCase {
	return &useCase{repository: repository}
}
