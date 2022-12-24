package interfaces

import (
	"engine/internal/core/entity"
)

type ScheduleUseCase interface {
	Save([]entity.Schedule) error
}

type GroupUseCase interface {
	Get(int) (*entity.Group, error)
	GetByName(string) (*entity.Group, error)
}

type StudentUseCase interface {
	Get(int) (*entity.Student, error)
	GetList() ([]entity.Student, error)
	Save(*entity.Student) (*entity.Student, error)
}

type DayUseCase interface {
	Get(entity.DayFilter) (*entity.Day, error)
}
