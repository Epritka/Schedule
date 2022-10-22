package interfaces

import (
	"engine/internal/core/entity"
	"time"
)

type ScheduleUseCase interface {
	Create([]entity.Schedule) error
	GetDay(date time.Time, groupId int) (entity.Day, error)
}

type UserUseCase interface {
	Update(login, groupName string) *int
}
