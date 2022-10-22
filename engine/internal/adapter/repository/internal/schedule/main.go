package schedule

import (
	"engine/internal/core/entity"
	"engine/internal/core/interfaces"
	"fmt"
)

type ScheduleRepository struct {
}

func NewScheduleRepository() interfaces.ScheduleRepository {
	return &ScheduleRepository{}
}

func (r *ScheduleRepository) GetDay(
	weekValue entity.WeekValue,
	weekDay entity.Weekday,
	groupId int,
) (entity.Day, error) {
	return entity.Day{}, fmt.Errorf("not implemented")
}
