package model

import (
	"engine/infrastructure/convert"
	"engine/internal/core/entity"
)

type LessonList []Lesson
type Lesson struct {
	Id         int
	DayId      int
	StartTime  string
	EndTime    string
	Name       string
	Type       string
	Teacher    string
	Auditorium string
	SubGroup   string
}

func NewLesson(value entity.Lesson) Lesson {
	return convert.Convert[entity.Lesson, Lesson](value)
}

func (value *Lesson) Entity() entity.Lesson {
	result := convert.DeConvert[Lesson, entity.Lesson](*value)
	return result
}

func (list *LessonList) Entity() []entity.Lesson {
	result := []entity.Lesson{}
	for _, value := range *list {
		result = append(result, value.Entity())
	}
	return result
}
