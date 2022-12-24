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
		usecase.logger.Error(err.Error())
		return nil, err
	}

	date := time.Unix(intDate, 0)
	_, thisWeek := date.ISOWeek()
	weekDay := getWeekDay(date)

	wt := ""

	if thisWeek%2 == 0 {
		wt = entity.EvenWeek
	} else {
		wt = entity.OddWeek
	}

	day, err := dayRepo.GetLessons(wt, weekDay, filters.GroupId)
	day.Number = (int)(weekDay)

	if err != nil {
		usecase.logger.Error(err.Error())
		return nil, err
	}

	return &day, nil
}
