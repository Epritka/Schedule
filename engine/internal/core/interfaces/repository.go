package interfaces

import (
	"engine/internal/core/entity"
)

type RepositoryManager interface {
	GetScheduleRepository() ScheduleRepository
	GetUserRepository() UserRepository
}

type ScheduleRepository interface {
	Create([]entity.Schedule) error
	GetDay(weekType entity.WeekType, weekDay entity.Weekday, groupId int) (entity.Day, error)
	GetGroupId(groupName string) *int
}

type UserRepository interface {
	Update(string, int) error
}
