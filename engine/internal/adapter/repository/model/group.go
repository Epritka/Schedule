package model

import (
	"engine/infrastructure/convert"
	"engine/internal/core/entity"
)

type GroupList []Group

type Group struct {
	tableName struct{} `sql:"student_group"`

	Id                       int
	Name                     string
	FacultyId                int
	YearId                   int
	EducationalInstitutionId int
}

func NewGroup(value entity.Group) Group {
	return convert.Convert[entity.Group, Group](value)
}

func (value *Group) Entity() entity.Group {
	result := convert.DeConvert[Group, entity.Group](*value)
	return result
}

func (list *GroupList) Entity() []entity.Group {
	result := []entity.Group{}
	for _, value := range *list {
		result = append(result, value.Entity())
	}
	return result
}
