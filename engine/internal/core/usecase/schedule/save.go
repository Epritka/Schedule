package schedule

import (
	"engine/internal/core/entity"
)

func (usecase *useCase) Save(schedules []entity.Schedule) error {
	educationalRepo := usecase.repository.GetEducationalRepository()
	groupRepository := usecase.repository.GetGroupRepository()
	lessonRepository := usecase.repository.GetLessonRepository()
	dayRepository := usecase.repository.GetDayRepository()

	for _, schedule := range schedules {
		ei, err := educationalRepo.SaveEducationalInstitution(schedule.EducationalInstitution.Name)
		if err != nil {
			usecase.logger.Error(err.Error())
			return err
		}

		f, err := educationalRepo.SaveFaculty(schedule.Faculty.Name)
		if err != nil {
			usecase.logger.Error(err.Error())
			return err
		}

		y, err := educationalRepo.SaveYear(schedule.Year.Name)
		if err != nil {
			usecase.logger.Error(err.Error())
			return err
		}

		group := entity.Group{
			Name:                     schedule.Group.Name,
			EducationalInstitutionId: ei.Id,
			FacultyId:                f.Id,
			YearId:                   y.Id,
		}

		g, err := groupRepository.Save(&group)
		if err != nil {
			usecase.logger.Error(err.Error())
			return err
		}

		for _, day := range schedule.EvenWeek {
			d, err := dayRepository.Save(entity.EvenWeek, day.Number, g.Id)
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
			d, err := dayRepository.Save(entity.OddWeek, day.Number, g.Id)
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
