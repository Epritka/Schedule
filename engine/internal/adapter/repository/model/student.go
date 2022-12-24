package model

import (
	"engine/infrastructure/convert"
	"engine/internal/core/entity"
)

type StudentList []Student

type Student struct {
	Id      int `json:"id"`
	UserId  int `json:"userId"`
	GroupId int `json:"groupId"`
}

func NewStudent(value entity.Student) Student {
	return convert.Convert[entity.Student, Student](value)
}

func (value *Student) Entity() entity.Student {
	result := convert.DeConvert[Student, entity.Student](*value)
	return result
}

func (list *StudentList) Entity() []entity.Student {
	result := []entity.Student{}
	for _, value := range *list {
		result = append(result, value.Entity())
	}
	return result
}
