package student

import (
	"engine/internal/core/entity"
)

func (usecase *useCase) Get(id int) (*entity.Student, error) {
	studentRepo := usecase.repository.GetStudentRepository()

	student, err := studentRepo.Get(id)

	if err != nil {
		usecase.logger.Error(err.Error())
		return nil, err
	}

	return student, nil
}

func (usecase *useCase) GetList() ([]entity.Student, error) {
	studentRepo := usecase.repository.GetStudentRepository()

	student, err := studentRepo.GetList()

	if err != nil {
		usecase.logger.Error(err.Error())
		return nil, err
	}

	return student, nil
}
