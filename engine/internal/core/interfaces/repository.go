package interfaces

import (
	"engine/internal/core/entity"
)

type RepositoryManager interface {
	GetScheduleRepository() ScheduleRepository
}

type ScheduleRepository interface {
	GetDay(weekValue entity.WeekValue, weekDay entity.Weekday, groupId int) (entity.Day, error)
}
