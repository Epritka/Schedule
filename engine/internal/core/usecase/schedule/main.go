package schedule

import "engine/internal/core/interfaces"

type useCase struct {
	repository interfaces.RepositoryManager
}

func NewScheduleUseCase(repository interfaces.RepositoryManager) interfaces.ScheduleUseCase {
	return &useCase{repository: repository}
}
