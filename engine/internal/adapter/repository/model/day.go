package model

import (
	"engine/infrastructure/convert"
	"engine/internal/core/entity"
)

type DayList []Day
type Day struct {
	Id       int
	Number   int
	WeekType string
	GroupId  int

	Lessons LessonList `pg:"rel:has-many,array"`
}

func NewDay(value entity.Day) Day {
	return convert.Convert[entity.Day, Day](value)
}

func (value *Day) Entity() entity.Day {
	result := convert.DeConvert[Day, entity.Day](*value)
	result.Lessons = value.Lessons.Entity()
	return result
}

func (list *DayList) Entity() []entity.Day {
	result := []entity.Day{}
	for _, value := range *list {
		result = append(result, value.Entity())
	}
	return result
}
