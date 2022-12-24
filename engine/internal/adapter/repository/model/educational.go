package model

import (
	"engine/infrastructure/convert"
	"engine/internal/core/entity"
)

type EducationalInstitution struct {
	Id   int
	Name string
}

type Faculty struct {
	Id   int
	Name string
}

type Year struct {
	Id   int
	Name string
}

func NewEducationalInstitution(value entity.EducationalInstitution) EducationalInstitution {
	return convert.Convert[entity.EducationalInstitution, EducationalInstitution](value)
}

func (value *EducationalInstitution) Entity() entity.EducationalInstitution {
	result := convert.DeConvert[EducationalInstitution, entity.EducationalInstitution](*value)
	return result
}

func NewFaculty(value entity.Faculty) Faculty {
	return convert.Convert[entity.Faculty, Faculty](value)
}

func (value *Faculty) Entity() entity.Faculty {
	result := convert.DeConvert[Faculty, entity.Faculty](*value)
	return result
}

func NewUYear(value entity.Year) Year {
	return convert.Convert[entity.Year, Year](value)
}

func (value *Year) Entity() entity.Year {
	result := convert.DeConvert[Year, entity.Year](*value)
	return result
}
