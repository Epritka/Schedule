package day

import (
	"engine/internal/core/entity"
	"fmt"
	"strconv"
	"time"
)

func getWeekDay(date time.Time) entity.Weekday {
	weekDay := date.Weekday()
	fmt.Println(weekDay)
	switch weekDay {
	case 0:
		return entity.Sunday
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

	// day.Number = (int)(weekDay)

	day, err := dayRepo.GetLessons(wt, weekDay, filters.GroupId)

	if err != nil {
		usecase.logger.Error(err.Error())
		return nil, err
	}

	return day, nil
}
