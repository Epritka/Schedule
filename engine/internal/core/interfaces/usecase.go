package interfaces

import (
	"engine/internal/core/entity"
)

type UserUseCase interface {
	Get(int) (*entity.User, error)
	GetByTgId(int) (*entity.User, error)
	Update(entity.User) error
}

type ScheduleUseCase interface {
	Create([]entity.Schedule) error
}

type GroupUseCase interface {
	Get(int) (entity.Group, error)
	GetByName(string) (entity.Group, error)
}

type DayUseCase interface {
	Get(weekType entity.WeekType, weekDay entity.Weekday, groupId int) (entity.Day, error)
}
