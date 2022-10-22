package interfaces

import (
	"engine/internal/core/entity"
	"time"
)

type ScheduleUseCase interface {
	GetDay(date time.Time, groupId int) (entity.Day, error)
}
