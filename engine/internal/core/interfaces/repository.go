package interfaces

import (
	"engine/internal/core/entity"
)

type RepositoryManager interface {
	Transaction(func(RepositoryManager) error) error

	GetGroupRepository() GroupRepository
	GetDayRepository() DayRepository
	GetScheduleRepository() ScheduleRepository
	GetUserRepository() UserRepository
}

type ScheduleRepository interface {
	Create([]entity.Schedule) error
}

type GroupRepository interface {
	Get(int) (*entity.Group, error)
	GetByName(string) (*entity.Group, error)
}

type DayRepository interface {
	Get(weekType entity.WeekType, weekDay entity.Weekday, groupId int) (entity.Day, error)
}

type UserRepository interface {
	Get(int) (*entity.User, error)
	GetByTgId(int) (*entity.User, error)
	GetByName(string) (*entity.User, error)
	Update(entity.User) error
}
