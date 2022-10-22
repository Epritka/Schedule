package usecase

import "engine/internal/core/interfaces"

type ScheduleUseCase struct {
	repository interfaces.RepositoryManager
}

func NewScheduleUseCase() interfaces.ScheduleUseCase {
	return &ScheduleUseCase{}
}
