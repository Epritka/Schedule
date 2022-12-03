package schedule

import (
	"engine/internal/core/entity"
	"fmt"
	"strings"
)

func (r *ScheduleRepository) GetDay(
	weekType entity.WeekType,
	weekDay entity.Weekday,
	groupId int,
) (entity.Day, error) {

	for _, s := range r.schedules {
		if s.GroupId == groupId {
			for _, w := range s.Week {
				if w.Number == int(weekDay) && w.WeekType == (weekType != 0) {
					day := entity.Day{
						Lessons: make([]entity.Lesson, 0),
					}

					for _, l := range w.Lessons {
						day.Lessons = append(
							day.Lessons,
							entity.Lesson{
								Time:           entity.Time{Start: l.Time.Start, End: l.Time.End},
								Name:           l.Name,
								Type:           l.Type,
								Teacher:        l.Teacher,
								Auditorium:     l.Auditorium,
								SubGroupNumber: l.SubGroupNumber,
							})

					}
					return day, nil
				}
			}
		}
	}

	return entity.Day{}, fmt.Errorf("not found")
}
func (r *ScheduleRepository) GetGroupId(groupName string) *int {
	if id, found := r.groups[strings.ToLower(groupName)]; found {
		fmt.Println(id)
		fmt.Println(strings.ToLower(groupName))
		return &id
	}

	return nil
}
