package repository

import (
	"engine/internal/core/interfaces"
	"engine/internal/mock/repository/internal/schedule"
	"engine/internal/mock/repository/internal/user"
)

type RepositoryManager struct {
	userRepository     interfaces.UserRepository
	scheduleRepository interfaces.ScheduleRepository
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
	if rm.userRepository == nil {
		rm.userRepository = user.NewUserRepository()
	}
	return rm.userRepository
}
