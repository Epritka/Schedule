package repository

import (
	"engine/internal/adapter/repository/internal/schedule"
	"engine/internal/adapter/repository/internal/user"
	"engine/internal/core/interfaces"
)

type RepositoryManager struct {
	userRepository     interfaces.UserRepository
	scheduleRepository interfaces.ScheduleRepository
}

func NewRepositoryManager() interfaces.RepositoryManager {
	return &RepositoryManager{}
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
