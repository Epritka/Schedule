package interfaces

import "engine/internal/core/entity"

type RepositoryManager interface {
	BeginTran() (RepositoryManager, error)
	CommitTran() error
	RollbackTran() error
	Transaction(callback func(RepositoryManager) error) error

	GetGroupRepository() GroupRepository
	GetDayRepository() DayRepository
	GetScheduleRepository() ScheduleRepository
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
