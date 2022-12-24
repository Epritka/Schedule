package student

import (
	"engine/internal/core/entity"
)

func (usecase *useCase) Save(student *entity.Student) (*entity.Student, error) {
	studentRepo := usecase.repository.GetStudentRepository()

	_, err := studentRepo.Save(student)
	if err != nil {
		usecase.logger.Error(err.Error())
		return nil, err
	}

	return student, nil
}
