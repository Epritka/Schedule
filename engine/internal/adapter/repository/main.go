package repository

import (
	"engine/internal/adapter/repository/internal/schedule"
	"engine/internal/core/interfaces"
)

type RepositoryManager struct {
}

func NewRepositoryManager() interfaces.RepositoryManager {
	return &RepositoryManager{}
}

func (rm *RepositoryManager) GetScheduleRepository() interfaces.ScheduleRepository {
	return schedule.NewScheduleRepository()
}
