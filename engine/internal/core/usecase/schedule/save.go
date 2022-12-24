package schedule

import (
	"engine/internal/core/entity"
)

func (usecase *useCase) Save(schedules []entity.Schedule) error {
	educationalRepo := usecase.repository.GetEducationalRepository()
	groupRepository := usecase.repository.GetGroupRepository()
	lessonRepository := usecase.repository.GetLessonRepository()
	dayRepository := usecase.repository.GetDayRepository()

	facultyIds := map[string]int{}
	educationalIds := map[string]int{}
	yearIds := map[string]int{}
	groupIds := map[string]int{}

	for _, schedule := range schedules {
		if _, find := educationalIds[schedule.EducationalInstitution.Name]; !find {
			ei, err := educationalRepo.SaveEducationalInstitution(schedule.EducationalInstitution.Name)
			if err != nil {
				usecase.logger.Error(err.Error())
				return err
			}

			educationalIds[schedule.EducationalInstitution.Name] = ei.Id
		}

		if _, find := facultyIds[schedule.Faculty.Name]; !find {
			f, err := educationalRepo.SaveFaculty(schedule.Faculty.Name)
			if err != nil {
				usecase.logger.Error(err.Error())
				return err
			}

			facultyIds[schedule.Faculty.Name] = f.Id
		}

		if _, find := yearIds[schedule.Year.Name]; !find {
			y, err := educationalRepo.SaveYear(schedule.Year.Name)
			if err != nil {
				usecase.logger.Error(err.Error())
				return err
			}

			yearIds[schedule.Year.Name] = y.Id
		}

		if _, find := groupIds[schedule.Group.Name]; !find {
			group := entity.Group{
				Name:                     schedule.Group.Name,
				EducationalInstitutionId: educationalIds[schedule.EducationalInstitution.Name],
				FacultyId:                facultyIds[schedule.Faculty.Name],
				YearId:                   yearIds[schedule.Year.Name],
			}

			g, err := groupRepository.Save(&group)
			if err != nil {
				usecase.logger.Error(err.Error())
				return err
			}

			groupIds[schedule.Group.Name] = g.Id
		}

		groupId := groupIds[schedule.Group.Name]

		for _, day := range schedule.EvenWeek {
			d, err := dayRepository.Save(entity.EvenWeek, day.Number, groupId)
			if err != nil {
				usecase.logger.Error(err.Error())
				return err
			}

			for _, lesson := range day.Lessons {
				newLesson := entity.Lesson{
					DayId:      d.Id,
					StartTime:  lesson.StartTime,
					EndTime:    lesson.EndTime,
					Name:       lesson.Name,
					Type:       lesson.Type,
					Teacher:    lesson.Teacher,
					Auditorium: lesson.Auditorium,
					SubGroup:   lesson.SubGroup,
				}

				_, err := lessonRepository.Save(&newLesson)
				if err != nil {
					usecase.logger.Error(err.Error())
					return err
				}
			}
		}

		for _, day := range schedule.OddWeek {
			d, err := dayRepository.Save(entity.OddWeek, day.Number, groupId)
			if err != nil {
				usecase.logger.Error(err.Error())
				return err
			}

			for _, lesson := range day.Lessons {
				newLesson := entity.Lesson{
					DayId:      d.Id,
					StartTime:  lesson.StartTime,
					EndTime:    lesson.EndTime,
					Name:       lesson.Name,
					Type:       lesson.Type,
					Teacher:    lesson.Teacher,
					Auditorium: lesson.Auditorium,
					SubGroup:   lesson.SubGroup,
				}

				_, err := lessonRepository.Save(&newLesson)
				if err != nil {
					usecase.logger.Error(err.Error())
					return err
				}
			}
		}
	}

	return nil
}
