package group

import (
	"engine/internal/core/entity"
)

func (usecase *useCase) Get(id int) (*entity.Group, error) {
	groupRepo := usecase.repository.GetGroupRepository()

	group, err := groupRepo.Get(id)

	if err != nil {
		usecase.logger.Error(err.Error())
		return nil, err
	}

	return group, nil
}

func (usecase *useCase) GetByName(name string) (*entity.Group, error) {
	groupRepo := usecase.repository.GetGroupRepository()

	group, err := groupRepo.GetByName(name)

	if err != nil {
		usecase.logger.Error(err.Error())
		return nil, err
	}

	return group, nil
}
