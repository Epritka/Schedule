package repository

import (
	"engine/internal/core/entity"
	"engine/internal/core/interfaces"
)

type dayRepository struct {
}

func NewDayRepository() interfaces.DayRepository {
	return &dayRepository{}
}

func (r *dayRepository) Get(
	weekType entity.WeekType,
	weekDay entity.Weekday,
	groupId int,
) (entity.Day, error) {
	return entity.Day{}, nil
}
