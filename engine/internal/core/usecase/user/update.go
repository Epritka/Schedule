package user

import (
	"engine/internal/core/entity"
	"errors"
)

func (usecase *useCase) Update(user entity.User) error {
	groupRepo := usecase.repository.GetGroupRepository()
	userRepo := usecase.repository.GetUserRepository()

	group, err := groupRepo.Get(user.GroupId)

	if err != nil {
		return err
	}

	if group == nil {
		return errors.New("not found")
	}

	err = userRepo.Update(user)

	if err != nil {
		return err
	}

	return nil
}
