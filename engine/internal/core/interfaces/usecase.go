package interfaces

import (
	"engine/internal/core/entity"
)

type ScheduleUseCase interface {
	Create([]entity.Schedule) error
}

type GroupUseCase interface {
	Get(int) (entity.Group, error)
	GetByName(string) (entity.Group, error)
}

type DayUseCase interface {
	Get(entity.DayFilter) (*entity.Day, error)
}
