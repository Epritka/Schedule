package repository

import (
	"engine/internal/core/interfaces"
)

type RepositoryManager struct {
}

func NewRepositoryManager() interfaces.RepositoryManager {
	return &RepositoryManager{}
}

func (rm *RepositoryManager) Transaction(callback func(r interfaces.RepositoryManager) error) error {
	return nil
}

func (rm *RepositoryManager) GetScheduleRepository() interfaces.ScheduleRepository {
	if rm.scheduleRepository == nil {
		rm.scheduleRepository = schedule.NewScheduleRepository()
	}
	return rm.scheduleRepository
}

func (rm *RepositoryManager) GetUserRepository() interfaces.UserRepository {
	return NewUserRepository()
}
