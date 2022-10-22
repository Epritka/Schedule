package schedule

import (
	"engine/internal/core/entity"
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

func (usecase *ScheduleUseCase) GetDay(date time.Time, groupId int) (entity.Day, error) {
	_, thisWeek := date.ISOWeek()

	scheduleRepo := usecase.repository.GetScheduleRepository()

	day, err := scheduleRepo.GetDay(entity.WeekType(thisWeek%2), getWeekDay(date), groupId)

	if err != nil {
		return day, err
	}

	return day, nil
}
