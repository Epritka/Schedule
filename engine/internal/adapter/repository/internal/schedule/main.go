package schedule

import (
	"engine/internal/core/interfaces"
)

type Schedule struct {
	EducationalInstitutionId int
	FacultyId                int
	YearId                   int
	GroupId                  int
	Week                     []Day
}

type Day struct {
	WeekType bool
	Number   int
	Lessons  []Lesson
}

type Lesson struct {
	Time           Time
	Name           string
	Type           string
	Teacher        string
	Auditorium     string
	SubGroupNumber *int
}

type Time struct {
	Start string
	End   string
}
type ScheduleRepository struct {
	ids                     []string
	faculties               map[string]int
	years                   map[string]int
	groups                  map[string]int
	educationalInstitutions map[string]int
	schedules               []Schedule
}

func NewScheduleRepository() interfaces.ScheduleRepository {
	return &ScheduleRepository{
		ids:                     make([]string, 0),
		faculties:               map[string]int{},
		years:                   map[string]int{},
		groups:                  map[string]int{},
		educationalInstitutions: map[string]int{},
		schedules:               make([]Schedule, 0),
	}
}
