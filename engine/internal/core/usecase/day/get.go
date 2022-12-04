package day

import (
	"engine/internal/core/entity"
	"errors"
)

func (usecase *useCase) Get(id int) (*entity.User, error) {
	userRepo := usecase.repository.GetUserRepository()

	user, err := userRepo.Get(id)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("not found")
	}

	return user, nil
}

func (usecase *useCase) GetByTgId(tgId int) (*entity.User, error) {
	userRepo := usecase.repository.GetUserRepository()

	user, err := userRepo.GetByTgId(tgId)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("not found")
	}

	return user, nil
}
