package day

import (
	"engine/internal/core/entity"
	"strconv"
	"time"
)

func getWeekDay(date time.Time) entity.Weekday {
	weekDay := date.Weekday()
	switch weekDay {
	case 0:
		return entity.Saturday
	default:
		return entity.Weekday(weekDay)
	}
}

func (usecase *useCase) Get(filters entity.DayFilter) (*entity.Day, error) {
	dayRepo := usecase.repository.GetDayRepository()

	intDate, err := strconv.ParseInt(filters.Date, 10, 64)
	if err != nil {
		return nil, err
	}

	date := time.Unix(intDate, 0)
	_, thisWeek := date.ISOWeek()
	weekDay := getWeekDay(date)

	day, err := dayRepo.Get(entity.WeekType(thisWeek%2), weekDay, filters.GroupId)
	day.Number = (*int)(&weekDay)

	if err != nil {
		return day, err
	}

	return day, nil
}
