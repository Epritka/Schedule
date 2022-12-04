package schedule

import (
	"engine/internal/core/entity"
	"strings"
)

func (r *ScheduleRepository) Create(schedules []entity.Schedule) error {
	for _, s := range schedules {
		schedule := Schedule{}

		r.ids = append(r.ids, s.EducationalInstitution.Name)
		appendedIndex := len(r.ids) - 1
		r.educationalInstitutions[s.EducationalInstitution.Name] = appendedIndex
		schedule.EducationalInstitutionId = appendedIndex

		r.ids = append(r.ids, s.Faculty.Name)
		appendedIndex = len(r.ids) - 1
		r.faculties[s.Faculty.Name] = appendedIndex
		schedule.FacultyId = appendedIndex

		r.ids = append(r.ids, s.Year.Name)
		appendedIndex = len(r.ids) - 1
		r.years[s.Year.Name] = appendedIndex
		schedule.YearId = appendedIndex

		r.ids = append(r.ids, s.Group.Name)
		appendedIndex = len(r.ids) - 1
		r.groups[strings.ToLower(s.Group.Name)] = appendedIndex
		schedule.GroupId = appendedIndex

		for _, w := range s.EvenWeek {
			day := Day{
				WeekType: false,
				Number:   *w.Number,
				Lessons:  make([]Lesson, 0),
			}

			for _, l := range w.Lessons {
				day.Lessons = append(
					day.Lessons,
					Lesson{
						Time:           Time{Start: l.Time.Start, End: l.Time.End},
						Name:           l.Name,
						Type:           l.Type,
						Teacher:        l.Teacher,
						Auditorium:     l.Auditorium,
						SubGroupNumber: l.SubGroupNumber,
					})
			}

			schedule.Week = append(schedule.Week, day)
		}

		for _, w := range s.OddWeek {
			day := Day{
				WeekType: true,
				Number:   *w.Number,
				Lessons:  make([]Lesson, 0),
			}

			for _, l := range w.Lessons {
				day.Lessons = append(
					day.Lessons,
					Lesson{
						Time:           Time{Start: l.Time.Start, End: l.Time.End},
						Name:           l.Name,
						Type:           l.Type,
						Teacher:        l.Teacher,
						Auditorium:     l.Auditorium,
						SubGroupNumber: l.SubGroupNumber,
					})
			}

			schedule.Week = append(schedule.Week, day)
		}

		r.schedules = append(r.schedules, schedule)
	}

	return nil
}
