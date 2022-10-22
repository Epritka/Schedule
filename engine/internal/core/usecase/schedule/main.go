package schedule

import "engine/internal/core/interfaces"

type ScheduleUseCase struct {
	repository interfaces.RepositoryManager
}

func NewScheduleUseCase(repository interfaces.RepositoryManager) interfaces.ScheduleUseCase {
	return &ScheduleUseCase{repository: repository}
}
